protoc:
	protoc pb/*.proto --go_out=plugins=grpc:.

build: 
	env GOOS=linux GOARCH=amd64 go build -o /tmp/bin server/server.go