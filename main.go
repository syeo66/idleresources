package main

import (
	"flag"
	"fmt"
	"time"

	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/syeo66/idleresources/gamestate"
	"golang.org/x/net/websocket"
)

var addr = flag.String("addr", ":8080", "http service address")

var gameState = gamestate.GameState{
	Resources: []gamestate.Resource{
		gamestate.NewWater(),
	},
	Tools: []gamestate.Tool{},
}

func websocketHandler(c echo.Context) error {
	websocket.Handler(func(ws *websocket.Conn) {
		defer ws.Close()
		for {
			// Write
			json, err := json.Marshal(gameState)
			if err != nil {
				c.Logger().Error(err)
				return
			}

			err = websocket.Message.Send(ws, string(json))
			if err != nil {
				c.Logger().Error(err)
				return
			}

			// Read
			msg := ""
			err = websocket.Message.Receive(ws, &msg)
			if err != nil {
				c.Logger().Error(err)
				return
			}
			fmt.Printf("%s\n", msg)
		}
	}).ServeHTTP(c.Response(), c.Request())
	return nil
}

func main() {
	flag.Parse()
	ticker := time.NewTicker(1 * time.Second)
	quit := make(chan struct{})

	go func() {
		for {
			select {
			case <-ticker.C:
				gameState.Tick()
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Static("/", "./public")
	e.GET("/ws", websocketHandler)
	e.Logger.Fatal(e.Start(*addr))
}
