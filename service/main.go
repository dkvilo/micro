package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/dkvilo/micro/service/handlers"
	"github.com/go-openapi/runtime/middleware"
	"github.com/julienschmidt/httprouter"
)

var fport string
var fenv string

func init() {
	flag.StringVar(&fport, "port", ":8080", "port number")
	flag.StringVar(&fenv, "env", "development", "application environment")
}

func main() {

	flag.Parse()

	// Custom Logger
	l := log.New(os.Stdout, "/message-api/", log.LstdFlags)
	
	// Message Handlers
	message := handlers.NewMessenger(l, fport)

	// Router
	sm := httprouter.New()
	
	sm.GET("/message", message.Handler)
	
	// host Redoc if app is on development mode
	// default env mode: development
	// to change env mode simply add -env=production or --env production
	if fenv == "development" {
		if _, err := os.Stat("./swagger.yaml"); os.IsNotExist(err) {
			l.Fatal("swagger.yaml not found, you may want to build it first! run 'make swagger' to do so")
			os.Exit(1)
		}
		sm.Handler("GET", "/swagger.yaml", http.FileServer(http.Dir("./")))
		sm.Handler("GET", "/docs", middleware.Redoc(middleware.RedocOpts{SpecURL: "/swagger.yaml"}, nil));
	}

	// Custom Server
	s := http.Server {
		Addr: 				fport,
		Handler: 			sm,
		ErrorLog: 		l,
		ReadTimeout: 	5 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	// Run the server in the goroutine
	go func() {
		log.Println("Running Service")
		log.Println(fmt.Sprintf(" - Port: %s", fport))
		log.Println(fmt.Sprintf(" - Environment: %s", fenv))
		if fenv == "development" {
			log.Println(fmt.Sprintf(" - Documentation: http://localhost%s/docs", fport))
		}

		err := s.ListenAndServe()
		if err != nil {
			log.Println(err.Error())
			os.Exit(1)
		}
	}()
	
	// Create os Signal Chanel 
	ch := make(chan os.Signal, 1)
	
	// Notify the chanel
	signal.Notify(ch, os.Interrupt)
	signal.Notify(ch, os.Kill)

	// Wait for signal
	log.Println("Received Signal:", <-ch)

	// and finlay, shutdown the server, wait for operation to be finished
	ctx, c := context.WithTimeout(context.Background(), 30 * time.Second)
	defer c()
	
	s.Shutdown(ctx)
}


