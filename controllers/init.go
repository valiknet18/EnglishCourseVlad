package controllers

type Controller interface {

}

var controllers map[string]Controller

func New() *controllers {
  controllers["UsersController"] = NewUsersController()

  return controllers
}
