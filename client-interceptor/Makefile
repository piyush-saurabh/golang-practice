protoc:
	protoc pb/*.proto --go_out=plugins=grpc:.

build: 
	env GOOS=linux GOARCH=amd64 go build -o /mnt/f/Temp/Shared/grpc server/server.go
	env GOOS=linux GOARCH=amd64 go build -o /mnt/f/Temp/Shared/grpc client/*.go