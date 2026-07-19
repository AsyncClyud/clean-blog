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
	"os"

	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Overload(); err != nil {
		log.Println(".Env not found")
	}

	cfg := config.Load()

	db := storage.ConnectDataBase(os.Getenv("DATABASE_URL"))
	defer db.Close()

	postDB := storage.NewPostRepo(db)
	userDB := storage.NewUserRepo(db)
	postService := service.NewPostService(*postDB)
	authService := service.NewAuthService(*userDB, []byte(cfg.JWTSecret))
	postHandler := handler.NewPostHandler(*postService, *authService, *cfg)
	userHandler := handler.NewUserHandler(*authService, *cfg)
	middleware := middleware.NewAuthMiddleware(*authService)

	router := router.Router(*postDB, *userDB, *postHandler, *userHandler, *authService, *middleware)

	log.Println("Server is started...")
	log.Printf("Go to http://localhost:%v", cfg.Port)

	err := http.ListenAndServe("localhost:"+cfg.Port, router)
	if err != nil {
		log.Fatalln(err)
	}

}
