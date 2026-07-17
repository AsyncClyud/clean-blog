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
	if len(article.Title) <= 3 {
		return http.StatusBadRequest
	}
	if len(article.Content) <= 3 {
		return http.StatusUnprocessableEntity
	}
	return http.StatusOK
}

func (pr PostService) GetArticles() (articles string) {
	return pr.repo.GetAllArticles()
}

func (pr PostService) GetArticleById(id int) (article string) {
	return pr.repo.GetArticleById(id)
}

func (pr PostService) InsertArticle(article models.Article, Author_Id int) (status_code int) {
	status := pr.ValidateArticle(article)
	if status == 400 {
		return http.StatusBadRequest
	}
	pr.repo.InsertArticle(article, Author_Id)
	return http.StatusOK
}

func (pr PostService) UpdateArticle(article models.Article) (status_code int) {
	status := pr.ValidateArticle(article)
	if status == 400 {
		return http.StatusBadRequest
	}
	pr.repo.UpdateArticle(article)
	return http.StatusOK
}

func (pr PostService) DeleteArticle(article models.Article) {
	pr.repo.DeleteArticle(article)
}

func (pr PostService) ValidateComment(comment models.Comment) (status_code int) {
	if len(comment.Comment_content) == 0 {
		return http.StatusBadRequest
	}
	return http.StatusOK
}

func (pr PostService) GetArticleCommentsById(id int) (comments string) {
	return pr.repo.GetArticleCommentsById(id)
}

func (pr PostService) InsertComment(comment models.Comment, author_id int) (status_code int) {
	status := pr.ValidateComment(comment)
	if status != http.StatusOK {
		return status
	}
	pr.repo.InsertComment(comment, author_id)
	return http.StatusOK
}
