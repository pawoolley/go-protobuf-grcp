package go_protobuf_grcp

// The purpose of this file is just to contain this comment, which tells
// Go what command to execute when someone types "go generate" in the
// root of the project.  In this case we're invoking protoc to generate
// Go files from the protobuf file.

//go:generate protoc --go_out=./ --go_opt=paths=source_relative --go-grpc_out=./ --go-grpc_opt=paths=source_relative protobuf/HelloWorld.proto
