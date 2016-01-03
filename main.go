package main

import (
		"fmt"
		"net/http"
		"github.com/valiknet18/EnglishCourseVlad/routes"
)

func main() {
		router := routes.Init();

		fmt.Println("Server running on port :8000")
		http.ListenAndServe(":8000", router)
}
