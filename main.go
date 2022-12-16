package main

import (
	"net/http"
	"zamannow/go-rest-api/db"
	"zamannow/go-rest-api/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	dbInstance := db.InitDB()
	defer dbInstance.Close()

	routers.RoutesHandler(r)

	s := &http.Server{
		Addr:    ":3000",
		Handler: r,
	}

	s.ListenAndServe()
}
