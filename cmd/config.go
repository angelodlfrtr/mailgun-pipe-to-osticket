package cmd

import "time"

// Config definition
type Config struct {
	ListenAddr           string `yaml:"listen_addr"`
	AuthToken            string `yaml:"auth_token"`
	OstScriptPath        string `yaml:"ost_script_path"`
	OstScriptExecTimeout string `yaml:"ost_script_exec_timeout"`
}

// NewConfig with defaults
func NewConfig() *Config {
	return &Config{
		ListenAddr:           "localhost:6789",
		AuthToken:            "CHANGE_ME",
		OstScriptPath:        "/var/www/api/pipe.php",
		OstScriptExecTimeout: "5s",
	}
}

func (config *Config) parseOstScriptExecTimeout() (time.Duration, error) {
	return time.ParseDuration(config.OstScriptExecTimeout)
}
