package grpcclient

import (
	"log"

	"bookmark-api/internal/config"
	pb "bookmark-api/proto"

	"google.golang.org/grpc"
)

var Client pb.PreviewServiceClient

func InitGRPC() {
	conn, err := grpc.Dial(config.AppConfig.GRPCPreviewAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatal("Failed to connect to gRPC server:", err)
	}
	Client = pb.NewPreviewServiceClient(conn)
}
