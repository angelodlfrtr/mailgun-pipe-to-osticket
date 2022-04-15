package cmd

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"gopkg.in/yaml.v3"
)

// Main entry
func Main() {
	// Check args
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "invalid config file: usage: osthttpiper [config.yaml]")
		os.Exit(1)
	}

	// Try to open config file
	yamlConfigBytes, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to read config file: %s\n", err.Error())
		os.Exit(1)
	}

	// Load default config
	cfg := NewConfig()
	if err := yaml.Unmarshal(yamlConfigBytes, cfg); err != nil {
		fmt.Fprintf(os.Stderr, "failed to parse yaml config: %s\n", err.Error())
		os.Exit(1)
	}
	ostScriptExecTimeout, err := cfg.parseOstScriptExecTimeout()
	if err != nil {
		fmt.Fprintf(
			os.Stderr,
			"failed to parse ost_script_exec_timeout to duration in config: %s\n",
			err.Error(),
		)
		os.Exit(1)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// We accept only POST requests
		if r.Method != "POST" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		// Check auth token
		if urlAuthToken := r.URL.Query().Get("auth_token"); urlAuthToken != cfg.AuthToken {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		// Check we have body
		if r.Body == nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// Get "body-mime" entry from request
		mimeEmail := r.PostFormValue("body-mime")
		if len(mimeEmail) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// Pipe body to osticket api/pipe.php script
		cmdCtx, cmdCancel := context.WithTimeout(context.Background(), ostScriptExecTimeout)
		defer cmdCancel()
		cmd := exec.CommandContext(cmdCtx, cfg.OstScriptPath)
		cmd.Stdin = strings.NewReader(mimeEmail)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			fmt.Fprintf(os.Stderr, "failed to pipe email with error: %s\n", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// Done
		w.WriteHeader(http.StatusOK)
	})

	if err := http.ListenAndServe(cfg.ListenAddr, nil); err != nil {
		fmt.Fprintf(os.Stderr, "failed to http server: %s\n", err.Error())
		os.Exit(1)
	}
}
