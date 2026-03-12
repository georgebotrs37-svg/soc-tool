package api

import (
	"fmt"
	"net/http"
)

func StartServer(db string) {
	fmt.Println("🌐 API Server running on http://localhost:8080")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "SOC Tool is Active!")
	})
	http.ListenAndServe(":8080", nil)
}
