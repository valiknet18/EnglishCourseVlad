package routes

import (
	"github.com/julienschmidt/httprouter"
	"github.com/valiknet18/EnglishCourseVlad/controllers"
)

func Init() *interface{} {
	controllersMap := controllers.New()
	router := httprouter.New()
	router.Get("/api/users/", controllersMap["UsersController"].Index)

	return router
}
