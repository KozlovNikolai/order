package app

import (
	"fmt"

	"github.com/KozlovNikolai/order/internal/endpoint"
	timecalc "github.com/KozlovNikolai/order/internal/service"
)

type App struct {
	e *endpoint.Endpoint
	s *timecalc.TimeCalc
}

func New() (*App, error) {
	a := &App{}
	a.e = endpoint.New()
	a.s = timecalc.New()
	return a, nil
}

func (a *App) Run() error {
	fmt.Println("App running.")

	return nil
}
