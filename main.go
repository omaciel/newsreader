package main

import (
	"net/http"
	"github.com/omaciel/newsreader/pkg/routes"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

func main() {
	port := ":8080"
	feed := routes.New()

	feed.Add(routes.Item{
		Title: "How to use Chi",
		URL:   "https://howto.com",
	})

	r := chi.NewRouter()
	r.Use(render.SetContentType(render.ContentTypeJSON))
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	r.Route("/feed", func(s chi.Router) {
		s.Mount("/", feed.ArticleRoutes())
	})

	http.ListenAndServe(port, r)
}
