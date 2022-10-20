####

Simple Go app using protobuf and gRPC.  Client sends a message.  Server echoes is back.

#### Pre reqs

See [here](https://grpc.io/docs/languages/go/quickstart/#prerequisites).

#### Generate code from protobuf

```bash
make generate
```

#### Running the example

1.
   ```bash
   cd server
   go build server.go
   ./server
   ```
1.
   ```bash
   # From a separate terminal to the server
   cd client
   go build client.go
   ./client
   ```
