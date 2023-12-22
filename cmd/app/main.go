package main

import (
	"github.com/imchiennb/acmex/internal/delivery/http"
	"github.com/imchiennb/acmex/pkg/database"
)

func main() {
	http.StartWebServer(database.Init())
}
