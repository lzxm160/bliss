export GOROOT=/usr/local/go
export GOPATH=/root/gopath:`pwd`
export GOBIN=/root/gopath/bin

rm -fr bliss
go build -o bliss ./src
./bliss