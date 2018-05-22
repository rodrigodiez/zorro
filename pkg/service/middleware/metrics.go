package middleware

import (
	"github.com/rodrigodiez/zorro/pkg/service"
)

type callMetrics struct {
	service.Zorro

	next        service.Zorro
	maskCounter Counter
}

func (m *callMetrics) Mask(key string) string {
	m.maskCounter.Add(int64(1))

	return m.next.Mask(key)
}

type Counter interface {
	Add(int64)
}

func NewCallMetrics(maskCounter Counter) MiddleWare {
	return func(service service.Zorro) service.Zorro {
		return &callMetrics{
			next:        service,
			maskCounter: maskCounter,
		}
	}
}
