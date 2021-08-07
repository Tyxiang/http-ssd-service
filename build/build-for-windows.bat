cd ../cmd
gofmt -w ./
cd ../bin
go build ../cmd/main.go
pause