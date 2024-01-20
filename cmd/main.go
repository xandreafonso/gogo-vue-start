package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("./web"))))

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)

		index := filepath.Join("./web/html", "index.html")
		parsedTemplate, _ := template.ParseFiles(index)

		parsedTemplate.Execute(w, nil)
	})

	serverPort := os.Getenv("PORT")

	if serverPort == "" {
		serverPort = "3000"
	}

	err := http.ListenAndServe(":"+serverPort, router)

	log.Println("Aplicação encerrada às", time.Now().Format(time.RFC3339))

	log.Fatal(err)
}
