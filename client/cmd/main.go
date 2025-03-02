package main

import (
	"net/http"

	"github.com/mjmichael73/go-simple-http-server-client/client/pkg/router"
)

func main() {
	r := router.InitializeRouter()
	http.ListenAndServe(":4000", r)
}
