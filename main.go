package main

import (
		"fmt"
		"net/http"
		"github.com/valiknet18/EnglishCourseVlad/routes"
)

func main() {
		router := routes.Init();

		http.ListenAndServe(":8000", router)
}
