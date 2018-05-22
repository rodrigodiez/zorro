package middleware

import (
	"testing"

	"github.com/rodrigodiez/zorro/lib/mocks"

	"github.com/rodrigodiez/zorro/pkg/service"
)

func TestNewMetricsImplementsMiddleWare(t *testing.T) {
	var _ MiddleWare = NewMetrics()
}

func TestMetricsIncreasesMaskCountOnCount(t *testing.T) {
	var zorro service.Zorro = &mocks.Zorro{}
	var middleware MiddleWare = NewMetrics()
	var service = middleware(zorro)

}
