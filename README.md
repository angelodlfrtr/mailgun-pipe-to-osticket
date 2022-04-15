# Pipe Mailgun HTTP emails to local OsTicket

Create a simple HTTP server to receive forwarded emails from mailgun and pipe thems to
local OsTicket pipe system.

## How

Make osticket `api/pipe.php` executable : `chmod +x path/to/osticket/api/pipe.php`.

Create yaml config file :

```yaml
listen_addr: localhost:6789
auth_token: CHANGE_ME
ost_script_path: path/to/osticket/api/pipe.php
ost_script_exec_timeout: 10s
```

Run : `mailgunostpiper [path to config]`

Use your webserver (apache, nginx, ...) to forward request on "somepathmime" to localhost:6789.

Then you can forward your emails to `http://your.server:port/somepathmime/?auth_token=CHANGE_ME`

**Important note**: Mailgun send the original mime message only when forward url ends with `mime`.

See : https://documentation.mailgun.com/en/latest/user_manual.html#routes

## Build

Run `make build`, binary generated to `build/mailgunostpiper`.

## License

Copyright (c) 2022 The contributors

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
