package main

import (
	"context"
	"log"
	"net"
	"net/http"

	pb "preview-service/proto"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedPreviewServiceServer
}

func (s *server) GetPreview(ctx context.Context, req *pb.PreviewRequest) (*pb.PreviewResponse, error) {
	url := req.GetUrl()

	// Fetch webpage
	resp, err := http.Get(url)
	if err != nil {
		return &pb.PreviewResponse{}, nil
	}
	defer resp.Body.Close()

	// VERY SIMPLE parsing (we improve later)
	buf := make([]byte, 1024)
	resp.Body.Read(buf)

	return &pb.PreviewResponse{
		Title:       string(buf[:50]), // placeholder
		Description: "Preview fetched",
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	pb.RegisterPreviewServiceServer(s, &server{})

	log.Println("gRPC Preview Service running on :50051")
	s.Serve(lis)
}
