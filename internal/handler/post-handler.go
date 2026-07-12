package handler

import (
	"blog/internal/contextutil"
	"blog/internal/models"
	"blog/internal/service"
	"blog/internal/storage"
	"encoding/json"
	"html/template"
	"net/http"
	"strconv"
)

type PostHandler struct {
	repo        storage.PostRepository
	postservice service.PostService
}

func NewPostHandler(repo storage.PostRepository, postservice service.PostService) *PostHandler {
	return &PostHandler{repo: repo, postservice: postservice}
}

func (psh PostHandler) ArticlePageHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("web/article/article.html")
	if err != nil {
		http.Error(w, "Invalid HTML file", http.StatusBadGateway)
		return
	}
	tmpl.Execute(w, nil)
}

func (psh PostHandler) GetArticlesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	articles := psh.repo.GetAllArticles()

	err := json.NewEncoder(w).Encode(articles)
	if err != nil {
		http.Error(w, "error", http.StatusBadGateway)
		return
	}
}

func (psh PostHandler) GetArticleByIdHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	Id, err := strconv.Atoi(r.PathValue("Id"))
	if err != nil {
		http.Error(w, "Invalid Id", http.StatusBadRequest)
		return
	}
	article := psh.repo.GetArticleById(Id)

	error := json.NewEncoder(w).Encode(article)
	if error != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}
}

func (psh PostHandler) CreateArticlePageHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("web/article/create_article.html")
	if err != nil {
		http.Error(w, "Invalid HTML file", http.StatusBadGateway)
		return
	}
	tmpl.Execute(w, nil)
}

func (psh PostHandler) InsertArticleHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Accept", "application/json")

	var data models.Article
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}

	status_code := psh.postservice.ValidateArticle(data)
	if status_code == 200 {
		userID, ok := r.Context().Value(contextutil.UserIDKey).(int)
		if !ok {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		psh.repo.InsertArticle(data, userID)
		ResponseCreateArticle(status_code, w, r)
	} else {
		ResponseCreateArticle(status_code, w, r)
	}

}

func (psh PostHandler) UpdateArticleHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Accept", "application/json")

	var data models.Article
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}

	psh.repo.UpdateArticle(data)

}

func (psh PostHandler) DeleteArticleHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Accept", "application/json")

	var data models.Article
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}

	psh.repo.DeleteArticle(data)

}
