package main

import (
	"github.com/DrLivsey00/ShopPraactice/internal/router"
	"log"
)

func main() {
	app := router.New()

	log.Fatal(app.Listen(":8080"))
}
