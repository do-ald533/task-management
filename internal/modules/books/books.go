package books

import (
	"goapi/internal/middleware"
	"goapi/internal/modules/books/controllers"
	"goapi/internal/modules/books/repositories"
	"goapi/internal/modules/books/services"
	"goapi/pkg/aws/sns"
	"goapi/pkg/permissions"

	"github.com/gofiber/fiber/v2"
)

func PrivateRoutes(route fiber.Router) {
	bookRepositories, _ := repositories.NewBookRepository()
	snsService := sns.NewSNSService()
	bookServices := services.NewBookService(bookRepositories, snsService)

	controller := controllers.NewBooksController(
		bookServices,
	)

	route.Get("/books", middleware.Credentials(permissions.BookReadCredential), controller.Index)
	route.Get("/books/:id", middleware.Credentials(permissions.BookReadCredential), controller.Show)
	route.Post("/books", middleware.Credentials(permissions.BookCreateCredential), controller.Store)
	route.Patch("/books/:id", middleware.Credentials(permissions.BookUpdateCredential), controller.Update)
	route.Delete("/books/:id", middleware.Credentials(permissions.BookDeleteCredential), controller.Delete)
}
