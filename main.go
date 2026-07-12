package main

import (
	"blog/internal/handler"
	"blog/internal/middleware"
	"blog/internal/router"
	"blog/internal/service"
	"blog/internal/storage"
	"log"
	"net/http"
)

func main() {

	db := storage.ConnectDataBase("user=myuser password=Hydra1234 dbname=sqldb sslmode=disable")
	defer db.Close()

	postDB := storage.NewPostRepo(db)
	userDB := storage.NewUserRepo(db)
	postService := service.NewPostService(*postDB)
	authService := service.NewAuthService(*userDB)
	postHandler := handler.NewPostHandler(*postDB, *postService)
	userHandler := handler.NewUserHandler(*authService)
	middleware := middleware.NewAuthMiddleware(*authService)

	router := router.Router(*postDB, *userDB, *postHandler, *userHandler, *authService, *middleware)

	log.Println("Server is started...")
	log.Println("Go to http://localhost:8080")

	err := http.ListenAndServe("127.0.0.1:8080", router)
	if err != nil {
		log.Fatalln(err)
	}

}
