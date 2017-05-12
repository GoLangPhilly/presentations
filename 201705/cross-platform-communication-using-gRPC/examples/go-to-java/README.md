First [install protoc](https://github.com/google/protobuf/blob/master/README.md)
Then install [Maven](https://maven.apache.org/install.html)

To generate the Go gRPC sources:

`protoc -I ./pb ./pb/file/service.proto --go_out=plugins=grpc:.`

To generate the Java gRPC sources:

`protoc -I ./pb --java_out ./fileserver/src/main/java ./pb/file/server.proto --grpc_out ./fileserver/src/main/java --plugin=protoc-gen-grpc=./grpc_java_plugin/protoc-gen-grpc-java-1.2.0-osx-x86_64.exe`

To build the Go code:

`cd go_client`
`go build -o go_client ./file`


To build and run the Java gRPC server:

`mvn clean package exec:java -Dexec.mainClass=org.golangphilly.grpc.file.transfer.FileTransferServerMain -Dexec.args="$HOME"`

In another terminal:
To run the Go client. This will upload the binary itself to the Java gRPC server:

`./go_client go_client`
  
