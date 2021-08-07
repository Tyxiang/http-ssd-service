cd ../cmd
gofmt -w ./
set GOARCH=amd64
set GOOS=linux
cd ../bin
go build ../cmd/main.go
pause