package internal

import (
	links "cortico/internal/home"
	"fmt"

	"github.com/leapkit/leapkit/core/server"
	_ "github.com/mattn/go-sqlite3"
)

// AddServices is a function that will be called by the server
// to inject services in the context.
func AddServices(r server.Router) error {
	db, err := DB()
	if err != nil {
		return fmt.Errorf("connecting to the database!: %w", err)
	}

	// Services that will be injected in the context
	r.Use(server.InCtxMiddleware("links", links.NewService(db)))

	return nil
}
