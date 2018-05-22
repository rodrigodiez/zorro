package middleware

import (
	"github.com/rodrigodiez/zorro/pkg/service"
)

type metrics struct {
	service.Zorro
	next service.Zorro
}

func NewMetrics() MiddleWare {
	return func(service service.Zorro) service.Zorro {
		return &metrics{
			next: service,
		}
	}
}
