if type go >/dev/null; then
	:
else
    export PATH=$PATH:/usr/local/go/bin
fi
export GOPATH=$PWD
export PATH=$PATH:$PWD/bin
