package delivery

import (
	"arch-linter-test/internal/delivery/handlers"
)

type Server struct {
	handler handlers.Handler
}
