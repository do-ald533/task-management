package controllers

import (
	"goapi/internal/modules/books/controllers"
	"goapi/test/units/books/mocks"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/valyala/fasthttp"

	"github.com/gofiber/fiber/v2"
)

type BooksControllerTestSuite struct {
	ctx        *fiber.Ctx
	controller *controllers.BooksController
	service    *mocks.BookServiceMock
	repository *mocks.BookRepositoryMock
	suite.Suite
}

func (suite *BooksControllerTestSuite) SetupTest() {
	suite.service = &mocks.BookServiceMock{}
	suite.repository = &mocks.BookRepositoryMock{}
	suite.controller = controllers.NewBooksController(suite.service)

	app := fiber.New()

	suite.ctx = app.AcquireCtx(&fasthttp.RequestCtx{})
}

func TestBooksControllerSuite(t *testing.T) {
	suite.Run(t, new(BooksControllerTestSuite))
	// ...
}
