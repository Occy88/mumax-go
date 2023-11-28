# mumax-go
# mumax-go

Put go installation into /home/$username/ location. E.g. for user with username "ndelser", put it in "/home/ndelser/go"
Change GOPATH and PATH in ~/.bashrc to
PATH="$GOPATH/bin:$PATH"
export GOPATH="$HOME/go"
go env -w GO111MODULE=off
