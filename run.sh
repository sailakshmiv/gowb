echo Setting up GOPATH
export GOPATH="$(dirname "$(dirname "$(dirname "$(dirname "$(pwd)")")")")"
export PATH=$PATH:$GOPATH/bin
echo Building gowb
go clean
go get
gowb


