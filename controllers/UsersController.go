package controllers

import (
    "fmt"
)

type UsersController struct {

}

func NewUsersController() *UsersController {
    return &UsersController{}
}

func (u UsersController) Index(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    fmt.Println("Hello world")
}
