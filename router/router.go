package router

import (
	"my-gram/controller"
	"my-gram/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

func UserRouter(r *mux.Router, uc controller.UserController) {
	r.HandleFunc("/users/register", uc.Register).Methods("POST")
	r.HandleFunc("/users/login", uc.Login).Methods("POST")
	r.Handle("/users/profile", middleware.Authentication(http.HandlerFunc(uc.GetUserById))).Methods("GET")
	r.Handle("/users", middleware.Authentication(http.HandlerFunc(uc.Update))).Methods("PUT")
	r.Handle("/users", middleware.Authentication(http.HandlerFunc(uc.Delete))).Methods("DELETE")
}

func PhotoRouter(r *mux.Router, pc controller.PhotoController) {
	r.Handle("/photo", middleware.Authentication(http.HandlerFunc(pc.CreatePhoto))).Methods("POST")
	r.Handle("/photos", middleware.Authentication(http.HandlerFunc(pc.GetPhotos))).Methods("GET")
	r.Handle("/photo/{id}", middleware.Authentication(http.HandlerFunc(pc.GetPhotoById))).Methods("GET")
	r.Handle("/photo/{id}", middleware.Authentication(http.HandlerFunc(pc.UpdatePhoto))).Methods("PUT")
	r.Handle("/photo/{id}", middleware.Authentication(http.HandlerFunc(pc.DeletePhoto))).Methods("DELETE")
}

func CommentRouter(r *mux.Router, cc controller.CommentController) {
	r.Handle("/comment", middleware.Authentication(http.HandlerFunc(cc.AddComment))).Methods("POST")
	r.Handle("/comments", middleware.Authentication(http.HandlerFunc(cc.GetAllComment))).Methods("GET")
	r.Handle("/comment/{id}", middleware.Authentication(http.HandlerFunc(cc.GetCommentById))).Methods("GET")
	r.Handle("/comment/{id}", middleware.Authentication(http.HandlerFunc(cc.UpdateComment))).Methods("PUT")
	r.Handle("/comment/{id}", middleware.Authentication(http.HandlerFunc(cc.DeleteComment))).Methods("DELETE")
}

func SocialMediaRouter(r *mux.Router, smc controller.SocialMediaController) {
	r.Handle("/social-media", middleware.Authentication(http.HandlerFunc(smc.CreateSocialMedia))).Methods("POST")
	r.Handle("/social-medias", middleware.Authentication(http.HandlerFunc(smc.GetAllSocialMedia))).Methods("GET")
	r.Handle("/social-media/{id}", middleware.Authentication(http.HandlerFunc(smc.GetAllSocialMedia))).Methods("GET")
	r.Handle("/social-media/{id}", middleware.Authentication(http.HandlerFunc(smc.UpdateSocialMedia))).Methods("PUT")
	r.Handle("/social-media/{id}", middleware.Authentication(http.HandlerFunc(smc.DeleteSocialMedia))).Methods("DELETE")
}