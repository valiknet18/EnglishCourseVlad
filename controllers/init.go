package controllers

import (
    // "fmt"
)

type Controllers struct {
  UsersController *UsersController
}

func New() Controllers {
  var controllers Controllers

  controllers.UsersController = NewUsersController()
  
  return controllers
}
