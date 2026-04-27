package grpcclient

import (
	"log"

	pb "bookmark-api/proto"

	"google.golang.org/grpc"
)

var Client pb.PreviewServiceClient

func InitGRPC() {
	conn, err := grpc.Dial("preview-service:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal("Failed to connect to gRPC server:", err)
	}
	Client = pb.NewPreviewServiceClient(conn)
}
