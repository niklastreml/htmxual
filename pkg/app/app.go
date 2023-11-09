package app

import "githbu.com/niklastreml/htmxual/pkg/services"

type App struct {
	counter services.CounterService
}

func New() *App {
	return &App{
		counter: *services.NewCounterService(),
	}
}
