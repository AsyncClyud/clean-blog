package router

import (
	"blog/internal/handler"
	"blog/internal/middleware"
	"blog/internal/service"
	"blog/internal/storage"
	"net/http"
)

func Router(postDB storage.PostRepository, userDB storage.UserRepository, postHandler handler.PostHandler, userHandler handler.UserHandler, authUser service.AuthService, middleware middleware.Middleware) *http.ServeMux {
	mux := http.NewServeMux()

	FsMain := http.FileServer(http.Dir("web"))
	mux.Handle("/", middleware.SecureHeaders(FsMain))
	fsArticles := http.FileServer(http.Dir("web/article"))
	mux.Handle("GET /web/articles", middleware.SecureHeaders(http.StripPrefix("/web/articles/", fsArticles)))
	fsAuth := http.FileServer(http.Dir("web/auth"))
	mux.Handle("GET /web/auth", middleware.SecureHeaders(http.StripPrefix("/web/auth/", fsAuth)))
	fsProfile := http.FileServer(http.Dir("web/profile"))
	mux.Handle("GET /web/profile", middleware.SecureHeaders(http.StripPrefix("/web/profile/", fsProfile)))

	mux.Handle("GET /privacy", middleware.SecureHeaders(http.HandlerFunc(postHandler.PrivacyPageHandler)))
	mux.Handle("GET /terms", middleware.SecureHeaders(http.HandlerFunc(postHandler.TermsPageHandler)))
	mux.Handle("GET /not_found", middleware.SecureHeaders(http.HandlerFunc(postHandler.NotFoundPageHandler)))
	mux.Handle("GET /api/auth", middleware.SecureHeaders(http.HandlerFunc(userHandler.IsAuth)))
	mux.Handle("GET /auth/register", middleware.SecureHeaders(http.HandlerFunc(userHandler.RegisterPageHandler)))
	mux.Handle("POST /auth/register", middleware.SecureHeaders(http.HandlerFunc(userHandler.RegisterHandler)))
	mux.Handle("GET /auth/login", middleware.SecureHeaders(http.HandlerFunc(userHandler.LoginPageHandler)))
	mux.Handle("POST /auth/login", middleware.SecureHeaders(http.HandlerFunc(userHandler.LoginHandler)))
	mux.Handle("GET /article/{Id}", middleware.SecureHeaders(http.HandlerFunc(postHandler.ArticlePageHandler)))
	mux.Handle("GET /api/articles/{Id}", middleware.SecureHeaders(http.HandlerFunc(postHandler.GetArticleByIdHandler)))
	mux.Handle("GET /api/articles", middleware.SecureHeaders(http.HandlerFunc(postHandler.GetArticlesHandler)))
	mux.Handle("GET /api/comments/{Id}", middleware.SecureHeaders(http.HandlerFunc(postHandler.GetArticleComments)))
	mux.Handle("GET /profile", middleware.SecureHeaders(http.HandlerFunc(userHandler.MainProfilePageHandler)))
	mux.Handle("GET /user_profile/{Id}", middleware.SecureHeaders(http.HandlerFunc(userHandler.UserProfilePageHandler)))

	mux.Handle("POST /api/logout", middleware.RequireAuth(http.HandlerFunc(userHandler.LogoutHandler)))
	mux.Handle("POST /api/users", middleware.RequireAuth(http.HandlerFunc(userHandler.GetArticleAuthorHandler)))
	mux.Handle("GET /api/profile", middleware.RequireAuth(http.HandlerFunc(userHandler.ProfileHandler)))
	mux.Handle("PUT /api/profile/username", middleware.RequireAuth(http.HandlerFunc(userHandler.ChangeUsernameHandler)))
	mux.Handle("PUT /api/profile/password", middleware.RequireAuth(http.HandlerFunc(userHandler.ChangePasswordHandler)))
	mux.Handle("PUT /api/profile/bio", middleware.RequireAuth(http.HandlerFunc(userHandler.ChangeBioHandler)))
	mux.Handle("GET /profile/settings", middleware.RequireAuth(http.HandlerFunc(userHandler.SettingsPageHandler)))
	mux.Handle("POST /api/articles", middleware.RequireAuth(http.HandlerFunc(postHandler.InsertArticleHandler)))
	mux.Handle("PUT /api/articles", middleware.RequireAuth(http.HandlerFunc(postHandler.UpdateArticleHandler)))
	mux.Handle("DELETE /api/articles", middleware.RequireAuth(http.HandlerFunc(postHandler.DeleteArticleHandler)))
	mux.Handle("GET /article/create", middleware.RequireAuth(http.HandlerFunc(postHandler.CreateArticlePageHandler)))
	mux.Handle("GET /article/update/{Id}", middleware.RequireAuth(http.HandlerFunc(postHandler.UpdateArticlePageHandler)))
	mux.Handle("POST /api/comments", middleware.SecureHeaders(http.HandlerFunc(postHandler.InsertCommentHandler)))

	return mux

}
