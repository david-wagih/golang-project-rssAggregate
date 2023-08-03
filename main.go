package main

import (
	"log"
	"os"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"net/http"
)

func main() {

	godotenv.Load(".env")

	portString := os.Getenv("PORT")
	if(portString == ""){
		log.Fatal("Port is not found in the environment")
	}

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
		ExposedHeaders: []string{"Link"},
		AllowCredentials: false,
		MaxAge: 300,
	}))



	srv := &http.Server{
		Handler: router,
		Addr: ":" + portString,
	}

	log.Printf("Server is listening on port %v", portString)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}