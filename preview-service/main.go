package main

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"golang.org/x/net/html"

	pb "preview-service/proto"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedPreviewServiceServer
}

func extractTitleFromJSON(data map[string]interface{}) string {
	// Try common JSON-LD title fields
	if title, ok := data["name"].(string); ok && title != "" {
		return title
	}
	if title, ok := data["headline"].(string); ok && title != "" {
		return title
	}
	return ""
}

func extractDescriptionFromJSON(data map[string]interface{}) string {
	// Try common JSON-LD description fields
	if desc, ok := data["description"].(string); ok && desc != "" {
		return desc
	}
	if desc, ok := data["about"].(string); ok && desc != "" {
		return desc
	}
	return ""
}

func parseJSONLD(n *html.Node) (string, string) {
	var title, description string

	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "script" {
			var scriptType string
			for _, attr := range n.Attr {
				if attr.Key == "type" && attr.Val == "application/ld+json" {
					scriptType = attr.Val
					break
				}
			}

			if scriptType == "application/ld+json" && n.FirstChild != nil {
				jsonStr := strings.TrimSpace(n.FirstChild.Data)

				var data map[string]interface{}
				if err := json.Unmarshal([]byte(jsonStr), &data); err == nil {
					if t := extractTitleFromJSON(data); t != "" && title == "" {
						title = t
					}
					if d := extractDescriptionFromJSON(data); d != "" && description == "" {
						description = d
					}
				}
			}
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(n)

	return title, description
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

	// Fallback to JSON-LD if title or description not found
	if title == "" || description == "" {
		jsonTitle, jsonDesc := parseJSONLD(doc)
		if title == "" {
			title = jsonTitle
		}
		if description == "" {
			description = jsonDesc
		}
	}

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
