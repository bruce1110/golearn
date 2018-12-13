package main

import (
	"sync"
	"github.com/gin-gonic/gin"
	"github.com/fvbock/endless"
	"bruce.net/Servers/middleware"
	"bruce.net/Servers/router"
	"bruce.net/Servers/run"
)

func init() {
	run.Init()
}

func main() {
	w := sync.WaitGroup{}
	w.Add(1)
	go func() {
		frontServer()
		w.Done()
	}()
	w.Wait()
}

func frontServer() {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())
	r.Use(middleware.Base)
	router.InitRoute(r)
	serevr := endless.NewServer(":8080", r)
	serevr.ListenAndServe()
}
