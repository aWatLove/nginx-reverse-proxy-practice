package main

import (
	"context"
	"nginxproxy/http"

	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	handler := http.NewHandler()
	srv := new(http.Server)

	go func() {
		if err := srv.Run("3000", handler.InitRoutes()); err != nil {
			log.Fatal(err)
		}
	}()
	log.Println("server started")

	//graceful shuttdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Println("server shutting down")
	if err := srv.Shutdown(context.Background()); err != nil {
		log.Printf("error occured on server shutting down: %s", err.Error())
	}

}
