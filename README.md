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
   go build server.go
   ```
1.
   ```bash
   go build client.go
   ```
