# GO CONFIGURATION WITH MUMAX
## ENVIRONMENT SETUP:
Execute this to use GOPATH explicitly, to be depreciated once scripts are fixed:
- `go env -w GO111MODULE=off`

add the following to `~/.bashrc`:
- `export GOPATH="$HOME/go"`
- `export PATH="$GOPATH/bin:$PATH"`
###
**don't forget to activate the environment: `. ~/.bashrc`**

## GO file setup:
- Move `go/` to `~/` directory in your server
