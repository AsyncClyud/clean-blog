package main

import (
	"blog/internal/config"
	"blog/internal/handler"
	"blog/internal/middleware"
	"blog/internal/router"
	"blog/internal/service"
	"blog/internal/storage"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Overload(); err != nil {
		log.Println(".Env not found")
	}

	cfg := config.Load()

	db := storage.ConnectDataBase(cfg.DSN())
	defer db.Close()

	postDB := storage.NewPostRepo(db)
	userDB := storage.NewUserRepo(db)
	postService := service.NewPostService(*postDB)
	authService := service.NewAuthService(*userDB, []byte(cfg.JWTSecret))
	postHandler := handler.NewPostHandler(*postService, *authService)
	userHandler := handler.NewUserHandler(*authService)
	middleware := middleware.NewAuthMiddleware(*authService)

	router := router.Router(*postDB, *userDB, *postHandler, *userHandler, *authService, *middleware)

	log.Println("Server is started...")
	log.Printf("Go to http://localhost:%v", cfg.Port)

	err := http.ListenAndServe(":"+cfg.Port, router)
	if err != nil {
		log.Fatalln(err)
	}

}
