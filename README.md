# grpc-test
GoLang gRPC test project

### server:
```bash
go run server.go
```
Environment variables for server:
- GRPC_TEST_SERVER_PORT - server port number (default 9099)
- GRPC_TEST_CLIENT_PORT - client port number (default 9077)


### client:
```bash
go run client.go http://google.com
```

Environment variables for client:
- GRPC_TEST_SERVER - server address (default "127.0.0.1:9099")
- GRPC_TEST_CLIENT_PORT - client port number (default 9077)

client writes log-file `client.log`
