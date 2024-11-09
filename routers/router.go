package router

import (
	"log"
	"main/controller/catagories"
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
				r.Get("/", catagories.CategoryHandler)
				r.Post("/", catagories.CategoryHandler)
			})

			// r.Route("/book", func(r chi.Router) {
			// 	r.Get("/discount", bookstore.DiscountBookHandler)
			// 	r.Route("/{bookID}", func(r chi.Router) {
			// 		r.Get("/edit", bookstore.EditBookHandler)
			// 		r.Post("/edit", bookstore.EditBookHandler)
			// 		r.Get("/delete", bookstore.DeleteBookHandler)
			// 	})
			// })

			// r.Route("/order", func(r chi.Router) {
			// 	r.Get("/", orders.OrderListHandler)
			// 	r.Route("/{orderID}", func(r chi.Router) {
			// 		r.Get("/detail", orders.OrderDetailHandler)
			// 	})
			// })
		})
	})

	r.MethodNotAllowed(MethodNotAllowedHandler)

	log.Println("server started on http://localhost:8080")
	http.ListenAndServe(":8080", r)
}

func MethodNotAllowedHandler(w http.ResponseWriter, r *http.Request) {
	library.StrucToJson(w, library.MethodNotAllowed)
}
