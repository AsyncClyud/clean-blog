package service

import (
	"blog/internal/models"
	"blog/internal/storage"
	"net/http"
)

type PostService struct {
	repo storage.PostRepository
}

func NewPostService(repo storage.PostRepository) *PostService {
	return &PostService{repo: repo}
}

func (pr PostService) ValidateArticle(article models.Article) (status_code int) {
	if len(article.Title) < 3 {
		return http.StatusBadRequest
	}
	if len(article.Content) == 0 {
		return http.StatusBadRequest
	}
	return http.StatusOK
}
