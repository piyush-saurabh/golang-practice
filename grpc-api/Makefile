# Install every dependencies using `make install`

# Generate .pb.go files from .proto file
START=$(pwd)

# All the .proto files are stored at location pb/v1/
protoc:
	protoc -I pb/v1/ \
		--go_out=plugins=grpc:pb \
		--gogrpcmock_out=:pb \
		pb/v1/*.proto

# adding new dependencies
install:
	go get -u \
		github.com/golang/protobuf/proto \
		github.com/golang/protobuf/protoc-gen-go \
		google.golang.org/grpc \
		github.com/gogo/protobuf/protoc-gen-gogoslick \
		github.com/gogo/protobuf/gogoproto \
		github.com/DATA-DOG/go-sqlmock \
		github.com/onsi/ginkgo/ginkgo \
		github.com/onsi/gomega/...
	go install github.com/SafetyCulture/s12-proto/protobuf/protoc-gen-gogrpcmock
	go get -tags 'mysql' -u github.com/golang-migrate/migrate/cmd/migrate

# Remove the protobuf file which we generate
	rm ./pb/**/*.pb.go

# ginko is the tool to test
test:
	ginkgo -r -failFast