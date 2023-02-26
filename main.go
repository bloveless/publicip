package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type server struct {
	httpClient http.Client
}

func (s server) hello(c echo.Context) error {
	resp, err := http.Get("https://ipinfo.io")
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	ip := struct {
		Ip string `json:"ip"`
	}{}
	err = json.NewDecoder(resp.Body).Decode(&ip)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.String(http.StatusOK, ip.Ip)
}

func main() {
	s := server{
		httpClient: http.Client{Timeout: 1 * time.Minute},
	}

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", s.hello)

	e.Logger.Fatal(e.Start(":8090"))
}
