package main

import (
	"database/sql"
	"fmt"
	"log"
	"my-gram/app"
	"my-gram/controller"
	"my-gram/middleware"
	"my-gram/repository"
	"my-gram/router"
	"my-gram/service"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/ilyakaznacheev/cleanenv"
	_ "github.com/lib/pq"
)

var cfg app.Config

func main() {
	_ = cleanenv.ReadConfig(".env", &cfg)
	app.Db, app.Err = sql.Open("postgres", ConnectDbPsql(
		cfg.Db_Host,
		cfg.Db_Dbname,
		cfg.Db_Username,
		cfg.Db_Password,
		cfg.Db_Port,
	))
	defer app.Db.Close()
	if app.Err != nil {
		panic(app.Err)
	}
	app.Err = app.Db.Ping()
	if app.Err != nil {
		panic(app.Err)
	}
	fmt.Println("Successfully Connect to Database")

	userRepository := repository.NewUserRepository()
	photoRepository := repository.NewPhotoRepository()
	commentRepository := repository.NewCommentRepository()
	socialMediaRepository := repository.NewSocialMediaRepository()

	userService := service.NewUserService(userRepository, app.Db)
	photoService := service.NewPhotoService(photoRepository,app.Db)
	commentService := service.NewCommentService(commentRepository, app.Db)
	socialMediaService := service.NewSocialMediaService(socialMediaRepository, app.Db)

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

func ConnectDbPsql(host, user, password, dbname string, port int) string {
	_ = cleanenv.ReadConfig(".env", &cfg)
	psqlInfo := fmt.Sprintf("host= %s port= %d user= %s "+
		" password= %s dbname= %s sslmode=disable",
		cfg.Db_Host,
		cfg.Db_Port,
		cfg.Db_Username,
		cfg.Db_Password,
		cfg.Db_Dbname)
	return psqlInfo
}