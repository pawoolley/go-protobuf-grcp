####

Simple Go app using protobuf and gRPC.  Client sends a message.  Server echoes is back.

#### Pre reqs

See [here](https://grpc.io/docs/languages/go/quickstart/#prerequisites).

#### Generate code from protobuf

```bash
# See generate.go
go generate
```

#### Running the example

1.
   ```bash
   go build -o server/server server/server.go
   ./server/server
   ```
1.
   ```bash
   # From a separate terminal to the server
   go build -o client/client client/client.go
   ./client/client
   ```
