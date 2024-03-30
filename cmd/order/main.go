package main

import (
	"log"

	"github.com/KozlovNikolai/order/internal/pkg/app"
)

func main() {
	a, err := app.New()
	if err != nil {
		log.Fatalln(err)
	}

	err = a.Run()
	if err != nil {
		log.Fatalln(err)
	}
}
