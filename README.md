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
- Clone this repostitory onto your server: `git clone https://github.com/occy88/mumax-go.git`
- Move `go/` to `~/` directory in your server: `mv ~/mumax-go/go ~/`
- execute: `cd ~/go/src/github.com/mumax && rm -rf 3 && git clone https://github.com/mumax/3`
- The above steps makes sure that the mumax/3 directory is properly cloned from https://github.com/mumax/3 if you get any error, fix it manually.
