mockgen:
	mockgen -source=./internal/product/repository.go -destination=./internal/product/mocks/rep_mock.go
	mockgen -source=./internal/user/repository.go -destination=./internal/user/mocks/rep_mock.go
	mockgen -source=./internal/session/repository.go -destination=./internal/session/mocks/rep_mock.go

	mockgen -source=./internal/product/delivery/grpc/product_grpc.pb.go -destination=./internal/product/delivery/grpc/mocks/pb_mock.go
	mockgen -source=./internal/user/delivery/grpc/user_grpc.pb.go       -destination=./internal/user/delivery/grpc/mocks/pb_mock.go
	mockgen -source=./internal/session/delivery/grpc/session_grpc.pb.go -destination=./internal/session/delivery/grpc/mocks/pb_mock.go

proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative  internal/product/delivery/grpc/product.proto
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative  internal/user/delivery/grpc/user.proto
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative  internal/session/delivery/grpc/session.proto

test:
	go test ./...