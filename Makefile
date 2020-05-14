.PHONY: proto clean

proto:
	protoc -I proto/ proto/*.proto --go_out=plugins=grpc:proto

clean:
	rm proto/*.pb.go

server:
	go run server.go
