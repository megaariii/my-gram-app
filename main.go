package main

import (
	"fmt"
	"log"
	"my-gram/app"
	"my-gram/controller"
	"my-gram/exception"
	"my-gram/middleware"
	"my-gram/repository"
	"my-gram/router"
	"my-gram/service"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/julienschmidt/httprouter"
)

func main() {
	db := app.NewDB()

	userRepository := repository.NewUserRepository()
	photoRepository := repository.NewPhotoRepository()
	commentRepository := repository.NewCommentRepository()
	socialMediaRepository := repository.NewSocialMediaRepository()

	userService := service.NewUserService(userRepository, db)
	photoService := service.NewPhotoService(photoRepository, commentRepository, db)
	commentService := service.NewCommentService(commentRepository, db)
	socialMediaService := service.NewSocialMediaService(socialMediaRepository, db)

	userController := controller.NewUserController(userService)
	photoController := controller.NewPhotoController(photoService)
	commentController := controller.NewCommentController(commentService)
	socialMediaController := controller.NewSocialMediaController(socialMediaService)

	r := mux.NewRouter()
	r.Use(middleware.LoggingMiddleware)

	router.UserRouter(r, userController)
	router.PhotoRouter(r, photoController)
	router.CommentRouter(r, commentController)
	router.SocialMediaRouter(r, socialMediaController)

	routerError := httprouter.New()
	routerError.PanicHandler = exception.ErrorHandler

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8080",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Println("Listening on port -> 127.0.0.1:8080")

	log.Fatal(srv.ListenAndServe())

}