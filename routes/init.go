package routes

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"github.com/valiknet18/EnglishCourseVlad/controllers"
)

func Init() http.Handler {
	controllersMap := controllers.New()
	router := httprouter.New()
	router.GET("/api/users/", controllersMap.UsersController.Index)

	return router
}
