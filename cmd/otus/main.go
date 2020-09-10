package main

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main(){
	logrus.Info("Hello World")

	port := os.Getenv("PORT")
	if port == ""{
		logrus.Fatal("Port is not set")
	}


	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	serv := &http.Server{
		Addr: net.JoinHostPort("",port),
		Handler: r,
	}

	go serv.ListenAndServe()

	interrupt := make(chan os.Signal,1)
	signal.Notify(interrupt,os.Interrupt,syscall.SIGTERM)

	<-interrupt
	timeout, cancelFunc := context.WithTimeout(context.Background(),5*time.Second)
	defer cancelFunc()
	serv.Shutdown(timeout)
}
