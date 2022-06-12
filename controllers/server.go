package controllers

import (
	"golang-sample/config"
	"golang-sample/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func StartEchoServer() {
	e := echo.New()
	e.Use(middleware.CORS())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/users", getAllUsers)
	e.GET("/users/:id", getUser)
	e.POST("/users", createUser)
	e.PUT("/users/:id", updateUser)
	e.DELETE("/users/:id", deleteUser)
	e.POST("/books", createBook)

	e.Logger.Fatal(e.Start(":" + config.Config.ServerPort))
}

func createUser(c echo.Context) error {
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return err
	}
	models.Db.Debug().Create(user)
	return c.JSON(http.StatusOK, user)
}

func getAllUsers(c echo.Context) error {
	users := []models.User{}
	models.Db.Debug().Find(&users)
	return c.JSON(http.StatusOK, users)
}

func getUser(c echo.Context) error {
	id := c.Param("id")
	user := new(models.User)
	models.Db.Debug().Preload("Books").Table("users").Find(&user, id)
	return c.JSON(http.StatusOK, user)
}

func deleteUser(c echo.Context) error {
	id := c.Param("id")
	user := new(models.User)
	models.Db.Debug().Delete(user, id)
	return c.JSON(http.StatusOK, user)
}

func updateUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return err
	}
	user.ID = id
	models.Db.Debug().Save(user)
	return c.JSON(http.StatusOK, user)
}

func createBook(c echo.Context) error {
	book := new(models.Book)
	if err := c.Bind(book); err != nil {
		return err
	}
	models.Db.Debug().Create(book)
	return c.JSON(http.StatusOK, book)
}
