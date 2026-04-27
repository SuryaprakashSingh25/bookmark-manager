package services

import (
	"bookmark-api/internal/grpcclient"
	"bookmark-api/internal/models"
	"bookmark-api/internal/repositories"
	pb "bookmark-api/proto"
	"context"
)

type BookmarkService struct {
	repo *repositories.BookmarkRepository
}

func NewBookmarkService() *BookmarkService {
	return &BookmarkService{
		repo: &repositories.BookmarkRepository{},
	}
}

func (s *BookmarkService) CreateBookmark(
	ctx context.Context,
	url string,
) (*models.Bookmark, error) {
	previewRes, err := grpcclient.Client.GetPreview(
		ctx,
		&pb.PreviewRequest{
			Url: url,
		},
	)

	if err != nil {
		previewRes = &pb.PreviewResponse{}
	}

	return s.repo.Create(
		ctx,
		url,
		previewRes.Title,
		previewRes.Description,
	)
}
