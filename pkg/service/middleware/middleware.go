package middleware

import (
	"github.com/rodrigodiez/zorro/pkg/service"
)

type MiddleWare func(service.Zorro) service.Zorro
