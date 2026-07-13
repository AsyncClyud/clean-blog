package router

import (
	"blog/internal/handler"
	"blog/internal/middleware"
	"blog/internal/service"
	"blog/internal/storage"
	"net/http"
)

func Router(postDB storage.PostRepository, userDB storage.UserRepository, postHandler handler.PostHandler, userHandler handler.UserHandler, authUser service.AuthService, middleware middleware.AuthMiddleware) *http.ServeMux {
	mux := http.NewServeMux()

	FsMain := http.FileServer(http.Dir("web"))
	mux.Handle("/", FsMain)

	fsArticles := http.FileServer(http.Dir("web/article"))
	mux.Handle("GET /web/articles", http.StripPrefix("/web/articles/", fsArticles))

	fsAuth := http.FileServer(http.Dir("web/auth"))
	mux.Handle("GET /web/auth", http.StripPrefix("/web/auth/", fsAuth))

	fsProfile := http.FileServer(http.Dir("web/profile"))
	mux.Handle("GET /web/profile", http.StripPrefix("/web/profile/", fsProfile))

	mux.Handle("GET /article/{Id}", http.HandlerFunc(postHandler.ArticlePageHandler))
	mux.Handle("GET /api/articles/{Id}", http.HandlerFunc(postHandler.GetArticleByIdHandler))
	mux.Handle("GET /api/articles", http.HandlerFunc(postHandler.GetArticlesHandler))
	mux.Handle("GET /auth/register", http.HandlerFunc(userHandler.RegisterPageHandler))
	mux.Handle("POST /auth/register", http.HandlerFunc(userHandler.RegisterHandler))
	mux.Handle("GET /auth/login", http.HandlerFunc(userHandler.LoginPageHandler))
	mux.Handle("POST /auth/login", http.HandlerFunc(userHandler.LoginHandler))
	mux.Handle("GET /profile", http.HandlerFunc(userHandler.ProfilePageHandler))

	mux.Handle("GET /api/profile", middleware.RequireAuth(http.HandlerFunc(userHandler.ProfileHandler)))
	mux.Handle("POST /api/users", middleware.RequireAuth(http.HandlerFunc(userHandler.GetArticleAuthorHandler)))
	mux.Handle("POST /api/articles", middleware.RequireAuth(http.HandlerFunc(postHandler.InsertArticleHandler)))
	mux.Handle("PUT /api/articles", middleware.RequireAuth(http.HandlerFunc(postHandler.UpdateArticleHandler)))
	mux.Handle("DELETE /api/articles", middleware.RequireAuth(http.HandlerFunc(postHandler.DeleteArticleHandler)))
	mux.Handle("GET /article/create", middleware.RequireAuth(http.HandlerFunc(postHandler.CreateArticlePageHandler)))

	return mux

}
