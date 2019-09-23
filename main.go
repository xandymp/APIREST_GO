package main

import (
	"github.com/gorilla/mux"
	"app"
	"os"
	"fmt"
	"net/http"
	"controllers"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/api/user/new", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/api/user/login", controllers.Authenticate).Methods("POST")
	router.HandleFunc("/api/funcionario/save", controllers.SaveFuncionario).Methods("POST")
	router.HandleFunc("/api/funcionario/save/", controllers.SaveFuncionario).Methods("POST")
	router.HandleFunc("/api/funcionario/delete/", controllers.DeleteFuncionario).Methods("DELETE")
	router.HandleFunc("/api/funcionario/", controllers.ListFuncionario).Methods("GET")
	router.HandleFunc("/api/funcionario", controllers.ListFuncionario).Methods("GET")
	router.HandleFunc("/api/funcionario/import", controllers.ImportFuncionario).Methods("POST")

	router.Use(app.JwtAuthentication)

	port := os.Getenv("PORT")
	if port == "" {
		port = "9002" //localhost
	}

	fmt.Println(port)

	err := http.ListenAndServe(":" + port, router)
	if err != nil {
		fmt.Print(err)
	}
}
