package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"golang.org/x/net/html"

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

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return &pb.PreviewResponse{}, nil
	}

	var title, description string

	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode {
			if n.Data == "title" && n.FirstChild != nil {
				title = n.FirstChild.Data
			}
			if n.Data == "meta" {
				var name, content string
				for _, attr := range n.Attr {
					if attr.Key == "name" {
						name = attr.Val
					}
					if attr.Key == "content" {
						content = attr.Val
					}
				}
				if name == "description" {
					description = content
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)

	return &pb.PreviewResponse{
		Title:       title,
		Description: description,
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
