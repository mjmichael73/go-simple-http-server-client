package main

import (
	"net/http"

	"github.com/mjmichael73/go-simple-http-server-client/server/pkg/router"
)

func main() {
	r := router.InitializeRouter()
	http.ListenAndServe(":3000", r)
}
