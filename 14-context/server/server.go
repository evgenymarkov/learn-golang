package server

import (
	"fmt"
	"net/http"

	"github.com/evgenymarkov/learn-golang/14-context/storage"
)

func CreateServer(store storage.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := store.Fetch(r.Context())

		if err != nil {
			return
		}

		fmt.Fprint(w, data)
	}
}
