package context

import (
	"context"
	"fmt"
	"net/http"
)

// Store interface with a Fetch function
type Store interface {
	Fetch(ctx context.Context) (string, error)
}

// Server returns string from store
func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := store.Fetch(r.Context())

		if err != nil {
			return
		}

		fmt.Fprint(w, data)
	}
}
