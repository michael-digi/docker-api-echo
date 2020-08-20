package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"docker.io/go-docker"
	"docker.io/go-docker/api/types"
	"github.com/labstack/echo"
)

// Container holds docker container info
type Container struct {
	ID      string `json:"Id"`
	Image   string
	ImageID string
	Command string
	Created int64
	State   string
	Status  string
}

func checkAPIKey(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		apiKey := c.Request().Header.Get("X-Api-Key")
		fmt.Println(apiKey, "this is apikey")

		if apiKey == "" {
			return nil
		}
		return next(c)
	}
}

func loggingMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		log.Println("Context", c)

		return next(c)
	}
}

func testing(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		log.Println("this is testing")

		return next(c)
	}
}

func getContainers(c echo.Context) error {
	fmt.Println("You hit containers")
	cli, err := docker.NewEnvClient()
	dockerJSON := []Container{}

	if err != nil {
		panic(err)
	}

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})

	if err != nil {
		panic(err)
	}

	for _, c := range containers {
		dockerJSON = append(dockerJSON, Container{c.ID[:10], c.Image, c.ImageID, c.Command, c.Created, c.State, c.Status})
	}

	response := map[string][]Container{
		"data": dockerJSON,
	}
	return c.JSON(http.StatusOK, response)
}

//main function
func main() {
	// create a new echo instance
	e := echo.New()

	e.GET("/containers", getContainers, checkAPIKey, loggingMiddleware, testing)

	e.Logger.Fatal(e.Start(":8000"))
}
