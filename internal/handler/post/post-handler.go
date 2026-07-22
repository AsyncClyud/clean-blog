package posthandler

import (
	"blog/internal/config"
	"blog/internal/models"
	postservice "blog/internal/service/post"
	userservice "blog/internal/service/user"
	captcha "blog/internal/turnstile"
	"encoding/json"
	"html/template"
	"net/http"
	"strconv"
)

type PostHandler struct {
	postService postservice.PostService
	authService userservice.AuthService
	Turnslite   captcha.Verifier
	Config      config.Config
}

func NewPostHandler(postservice postservice.PostService, auth userservice.AuthService, config config.Config) *PostHandler {
	return &PostHandler{postService: postservice, authService: auth, Turnslite: *captcha.NewVerifier(config), Config: config}
}

func (psh *PostHandler) PrivacyPageHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("web/privacy.html")
	if err != nil {
		http.Error(w, "Invalid HTML file", http.StatusBadGateway)
		return
	}
	tmpl.Execute(w, nil)
}

func (psh *PostHandler) TermsPageHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("web/terms.html")
	if err != nil {
		http.Error(w, "Invalid HTML file", http.StatusBadGateway)
		return
	}
	tmpl.Execute(w, nil)
}

func (psh *PostHandler) ArticlePageHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("web/article/article.html")
	if err != nil {
		http.Error(w, "Invalid HTML file", http.StatusBadGateway)
		return
	}
	tmpl.Execute(w, nil)
}

func (psh *PostHandler) CreateArticlePageHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("web/article/create_article.html")
	if err != nil {
		http.Error(w, "Invalid HTML file", http.StatusBadGateway)
		return
	}
	tmpl.Execute(w, nil)
}

func (psh *PostHandler) UpdateArticlePageHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("web/article/update_article.html")
	if err != nil {
		http.Error(w, "Invalid HTML file", http.StatusBadGateway)
		return
	}
	tmpl.Execute(w, nil)
}

func (psh *PostHandler) NotFoundPageHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("web/not_found.html")
	if err != nil {
		http.Error(w, "Invalid HTML file", http.StatusBadGateway)
	}
	tmpl.Execute(w, nil)
}

func (psh *PostHandler) GetArticlesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	articles := psh.postService.GetArticles()

	err := json.NewEncoder(w).Encode(articles)
	if err != nil {
		http.Error(w, "Invalid body", http.StatusBadGateway)
		return
	}
}

func (psh *PostHandler) GetArticleByIdHandler(w http.ResponseWriter, r *http.Request) {
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

func (psh *PostHandler) InsertArticleHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Accept", "application/json")

	var article models.Article
	err := json.NewDecoder(r.Body).Decode(&article)
	if err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}

	cfToken := article.Turnstile_token
	remoteAddr := r.RemoteAddr

	ok, err := psh.Turnslite.Verify(r.Context(), cfToken, remoteAddr)
	if err != nil || !ok {
		status_code := http.StatusForbidden
		ResponseArticle(status_code, w, r)
		return
	}

	cookie, err := r.Cookie("jwt-token")
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	userID, err := psh.authService.Validate_Token(cookie.Value)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	status_code := psh.postService.InsertArticle(article, userID)
	ResponseArticle(status_code, w, r)
}

func (psh *PostHandler) UpdateArticleHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Accept", "application/json")

	var article models.Article
	err := json.NewDecoder(r.Body).Decode(&article)
	if err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}

	cfToken := article.Turnstile_token
	remoteAddr := r.RemoteAddr

	ok, err := psh.Turnslite.Verify(r.Context(), cfToken, remoteAddr)
	if err != nil || !ok {
		status_code := http.StatusForbidden
		ResponseArticle(status_code, w, r)
		return
	}

	cookie, err := r.Cookie("jwt-token")
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	userID, err := psh.authService.Validate_Token(cookie.Value)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	if article.Author != userID {
		ResponseArticle(http.StatusUnauthorized, w, r)
		return
	}
	status_code := psh.postService.UpdateArticle(article)
	ResponseArticle(status_code, w, r)
}

func (psh *PostHandler) DeleteArticleHandler(w http.ResponseWriter, r *http.Request) {
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
	userID, err := psh.authService.Validate_Token(cookie.Value)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	if article.Author != userID {
		ResponseArticle(http.StatusForbidden, w, r)
		return
	}
	psh.postService.DeleteArticle(article)
}

func (psh *PostHandler) GetArticleComments(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, ok := strconv.Atoi(r.PathValue("Id"))
	if ok != nil {
		http.Error(w, "Invalid id", http.StatusNotFound)
	}

	comments := psh.postService.GetArticleCommentsById(id)

	err := json.NewEncoder(w).Encode(comments)
	if err != nil {
		http.Error(w, "Internal error", http.StatusBadGateway)
	}
}

func (psh *PostHandler) InsertCommentHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Accept", "application/json")

	var comment models.Comment
	err := json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}

	cfToken := comment.Turnstile_token
	remoteAddr := r.RemoteAddr

	ok, err := psh.Turnslite.Verify(r.Context(), cfToken, remoteAddr)
	if err != nil || !ok {
		status_code := http.StatusForbidden
		ResponseComment(status_code, w, r)
		return
	}

	cookie, err := r.Cookie("jwt-token")
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	userID, err := psh.authService.Validate_Token(cookie.Value)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	status_code := psh.postService.InsertComment(comment, userID)
	ResponseComment(status_code, w, r)
}
