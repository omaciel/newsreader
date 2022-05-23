package routes

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
)


type Item struct {
	Title string `json:"title"`
	URL   string `json:"url"`
}

type Articles struct {
	Items []Item
}

type Getter interface {
	GetAll() []Item
}

type Adder interface {
	Add(item Item)
}

func New() *Articles {
	return &Articles{
		Items: []Item{},
	}
}

func (a *Articles) Add(item Item) {
	a.Items = append(a.Items, item)
}

func (a *Articles) GetAll(w http.ResponseWriter, r *http.Request) {
	items := a.Items
	json.NewEncoder(w).Encode(items)
}

func (a *Articles) ArticleRoutes() chi.Router {
	router := chi.NewRouter()
	router.Get("/", a.GetAll)

	return router
}
