package router

import (
	"log"
	"main/controller/catagories"
	"main/controller/invesment"
	"main/controller/items"
	"main/controller/reminder"
	usershandler "main/controller/userHandler"
	"main/library"
	middlewaree "main/middleware"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func InitRoute() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/", func(r chi.Router) {
		r.Post("/login", usershandler.LoginHandler)
		r.Post("/register", usershandler.RegisterHandler)

		r.With(middlewaree.TokenMiddleware).Route("/api", func(r chi.Router) {
			r.Route("/categories", func(r chi.Router) {
				r.Get("/{id}", catagories.GetCategoryByIdHandler)
				r.Put("/{id}", catagories.PutCategoryByIdHandler)
				r.Delete("/{id}", catagories.DeleteCategoryByIdHandler)
				r.Get("/", catagories.CategoryHandler)
				r.Post("/", catagories.CategoryHandler)
			})

			r.Route("/items/", func(r chi.Router) {
				r.Get("/{id}", items.GetInventoryItemByIdHandler)
				r.Put("/{id}", items.UpdateInventoryItemByIdHandler)
				r.Delete("/{id}", items.DeleteInventoryItemByIdHandler)
				r.Post("/", items.AddInventoryItemHandler)
				r.Get("/{page}/{limit}", items.GetItemsPaginated)
			})

			r.Route("/items/investment", func(r chi.Router) {
				r.Get("/{id}", invesment.GetItemsInvesmentByIdHandler)
				r.Get("/", invesment.GetItemsInvesmentHandler)
			})

			r.Route("/items/replacement-needed", func(r chi.Router) {
				r.Get("/", reminder.GetItemsReplacementHandler)
			})
		})
	})

	r.MethodNotAllowed(MethodNotAllowedHandler)

	log.Println("server started on http://localhost:8080")
	http.ListenAndServe(":8080", r)
}

func MethodNotAllowedHandler(w http.ResponseWriter, r *http.Request) {
	library.StrucToJson(w, library.MethodNotAllowed)
}
