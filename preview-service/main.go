package main

import (
	"context"
	"io"
	"log"
	"net"
	"net/http"
	"net/url"
	"sync"
	"time"

	"golang.org/x/net/html"

	pb "preview-service/proto"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedPreviewServiceServer
}

func (s *server) GetPreview(ctx context.Context, req *pb.PreviewRequest) (*pb.PreviewResponse, error) {
	rawURL := req.GetUrl()

	parsedURL, err := url.ParseRequestURI(rawURL)
	if err != nil {
		return &pb.PreviewResponse{}, nil
	}

	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	httpReq, err := http.NewRequest("GET", parsedURL.String(), nil)
	if err != nil {
		return &pb.PreviewResponse{}, nil
	}

	httpReq.Header.Set("User-Agent", "Mozilla/5.0")

	resp, err := client.Do(httpReq)
	if err != nil {
		return &pb.PreviewResponse{}, nil
	}
	defer resp.Body.Close()

	limitedReader := io.LimitReader(resp.Body, 1_000_000)

	doc, err := html.Parse(limitedReader)
	if err != nil {
		return &pb.PreviewResponse{}, nil
	}

	var title string
	var description string
	var mu sync.Mutex

	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode {
			if n.Data == "title" && n.FirstChild != nil && title == "" {
				title = n.FirstChild.Data
			}
			if n.Data == "meta" {
				var name, property, content string
				for _, attr := range n.Attr {
					switch attr.Key {
					case "name":
						name = attr.Val

					case "property":
						property = attr.Val

					case "content":
						content = attr.Val
					}
				}
				if name == "description" && description == "" {
					mu.Lock()
					description = content
					mu.Unlock()
				}
				if property == "og:title" {
					mu.Lock()
					title = content
					mu.Unlock()
				}
				if property == "og:description" {
					mu.Lock()
					description = content
					mu.Unlock()
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
