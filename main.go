package main

import (
	"fmt"
	"net/http"
	"log"
	"github.com/douglasmakey/oauth2-example/handlers"
)

func main() {
	
	os.Setenv("GOOGLE_OAUTH_CLIENT_ID", "308872887952-3fgbmida5f4rs5q11v63a5nmgmtv55di.apps.googleusercontent.com")
	os.Setenv("GOOGLE_OAUTH_CLIENT_SECRET","IxwMhR86rlSmUtRU-OB9eGsI")
	
	
	// We create a simple server using http.Server and run.
	server := &http.Server{
		Addr: fmt.Sprintf(":"+os.os.Getenv("PORT")),
		Handler: handlers.New(),
	}

	log.Printf("Starting HTTP Server. Listening at %q", server.Addr)
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Printf("%v", err)
	} else {
		log.Println("Server closed!")
	}
}
