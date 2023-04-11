package controllers

import (
	"encoding/json"
	"fmt"
	"goapi/test/units/books/mocks"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func (suite *BooksControllerTestSuite) TestSuccessStoreBook() {
	suite.ctx.Request().Header.Set("Authorization", fmt.Sprintf("Bearer %s", mocks.UserLogged.Access))
	suite.ctx.Request().Header.SetContentType(fiber.MIMEApplicationJSON)

	body, _ := json.Marshal(mocks.StoreBookMock)
	suite.ctx.Request().SetBody(body)

	suite.service.On("Store", mocks.UserId, &mocks.StoreBookMock).Return(&mocks.BookMock, nil)
	err := suite.controller.Store(suite.ctx)

	expectedResponse, _ := json.Marshal(mocks.BookMock)

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), fiber.StatusCreated, suite.ctx.Response().StatusCode())
	assert.JSONEq(suite.T(), string(expectedResponse), string(suite.ctx.Response().Body()))
}
