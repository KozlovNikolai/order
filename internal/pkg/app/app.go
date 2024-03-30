package app

import (
	"fmt"
	"log"

	"github.com/KozlovNikolai/order/internal/endpoint"
	"github.com/KozlovNikolai/order/internal/service/timecalc"
)

type App struct {
	e  *endpoint.Endpoint
	tc *timecalc.TimeCalc
}

func New() (*App, error) {
	a := &App{}
	a.tc = timecalc.New()
	a.e = endpoint.New(a.tc)
	return a, nil
}

func (a *App) Run() error {
	fmt.Println("App running.")
	err := a.e.Status()
	if err != nil {
		log.Fatalln(err)
	}
	return nil
}
