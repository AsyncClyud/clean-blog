package handler

import (
	"blog/internal/models"
	"blog/internal/service"
	"encoding/json"
	"html/template"
	"net/http"
	"strconv"
)

type PostHandler struct {
	postService service.PostService
	authService service.AuthService
}

func NewPostHandler(postservice service.PostService, auth service.AuthService) *PostHandler {
	return &PostHandler{postService: postservice, authService: auth}
}

func (psh PostHandler) PrivacyPageHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("web/privacy.html")
	if err != nil {
		http.Error(w, "Invalid HTML file", http.StatusBadGateway)
		return
	}
	tmpl.Execute(w, nil)
}

func (psh PostHandler) TermsPageHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("web/terms.html")
	if err != nil {
		http.Error(w, "Invalid HTML file", http.StatusBadGateway)
		return
	}
	tmpl.Execute(w, nil)
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

	articles := psh.postService.GetArticles()

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
	article := psh.postService.GetArticleById(Id)

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

	var article models.Article
	err := json.NewDecoder(r.Body).Decode(&article)
	if err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}
	cookie, err := r.Cookie("jwt-token")
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	claims, err := psh.authService.Validate_Token(cookie.Value)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	status_code := psh.postService.InsertArticle(article, claims)
	ResponseArticle(status_code, w, r)

}

func (psh PostHandler) UpdateArticleHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Accept", "application/json")

	var article models.Article
	err := json.NewDecoder(r.Body).Decode(&article)
	if err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}

	cookie, err := r.Cookie("jwt-token")
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	claims, err := psh.authService.Validate_Token(cookie.Value)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	if article.Author != claims {
		ResponseArticle(http.StatusForbidden, w, r)
		return
	}
	status_code := psh.postService.UpdateArticle(article)
	ResponseArticle(status_code, w, r)

}

func (psh PostHandler) DeleteArticleHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Accept", "application/json")

	var article models.Article
	err := json.NewDecoder(r.Body).Decode(&article)
	if err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}

	cookie, err := r.Cookie("jwt-token")
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	claims, err := psh.authService.Validate_Token(cookie.Value)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	if article.Author != claims {
		ResponseArticle(http.StatusForbidden, w, r)
		return
	}

	psh.postService.DeleteArticle(article)

}
