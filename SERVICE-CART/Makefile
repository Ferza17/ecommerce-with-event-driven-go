create:
	protoc --proto_path=model/proto model/proto/*.proto --go_out=.
	protoc --proto_path=model/proto model/proto/*.proto --go-grpc_out=.

test:
	@go clean -testcache && go test -v -race -short -failfast -cover ./...

clean:
	rm -r model/pb/*