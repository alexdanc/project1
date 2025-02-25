package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type RequestBody struct {
	Task string `json:"task"`
}

type Response struct {
	Status      string `json:"status"`
	RequestBody string `json:"body"`
}

var task []RequestBody

func GetHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, &task)
}

func PostHandler(c echo.Context) error {
	var tasks RequestBody
	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status:      "failed",
			RequestBody: " Don't Body",
		})
	}
	task = append(task, tasks)
	return c.JSON(http.StatusOK, Response{
		Status:      "success",
		RequestBody: "Body Added",
	})
}

func main() {
	e := echo.New()
	e.GET("/task", GetHandler)
	e.POST("/task", PostHandler)
	e.Logger.Fatal(e.Start(":8080"))
}
