.PHONY:proto
proto:
	protoc --go_out=plugins=grpc:. --proto_path=. ./proto/role/role.proto
	protoc --go_out=plugins=grpc:. --proto_path=. ./proto/user/user.proto
	protoc --go_out=plugins=grpc:. --proto_path=. ./proto/message/message.proto