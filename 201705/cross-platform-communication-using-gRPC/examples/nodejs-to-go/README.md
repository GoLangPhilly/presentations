First [install protoc](https://github.com/google/protobuf/blob/master/README.md)

To generate the Go gRPC sources:

`protoc -I ./pb ./pb/docker/service.proto --go_out=plugins=grpc:.`


To build the code:

`cd docker`
`go build .`

To run the gRPC Server:

`./docker`
   

In another terminal:

`cd nodejs_client`
`npm install`

*Get All containers*

`node client.js 1` 

*Get Stats for Container*

`node client.js 2 <container>`
