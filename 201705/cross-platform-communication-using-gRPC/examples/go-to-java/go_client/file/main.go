package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	grpcClient "./transfer"
	"google.golang.org/grpc"
)

const (
	chunkSize = 1024 * 1024
)

type config struct {
	serverHostPost string
}

func getConfig() *config {

	envConfig := config{}

	envConfig.serverHostPost = os.Getenv("GRPC_HOST_PORT")
	if envConfig.serverHostPost == "" {
		envConfig.serverHostPost = "localhost:9080"
	}

	return &envConfig
}

func main() {

	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s <File to Upload>", os.Args[0])
		os.Exit(1)
	}
	filename := os.Args[1]

	// Set up a connection to the server.
	conn, err := grpc.Dial(getConfig().serverHostPost, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := grpcClient.NewFileTransferServiceClient(conn)
	uploadFile(filename, client)

}

func uploadFile(filename string, client grpcClient.FileTransferServiceClient) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	stream, err := client.Upload(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	for {
		chunk := make([]byte, chunkSize)
		n, err := f.Read(chunk)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if n < len(chunk) {
			chunk = chunk[:n]
		}
		fr := &grpcClient.FileRequest{
			Data:     chunk,
			Filename: filename,
		}
		stream.Send(fr)
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatal(err)
	}
	if res.IsOk {
		fmt.Printf("uploaded file: %s - %d", res.Filename, res.Size)
	} else {
		fmt.Fprintln(os.Stderr, "Filed to upload file!")
	}
}
