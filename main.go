package main

import (
	"flag"
	"time"

	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/syeo66/idleresources/gamestate"
	"golang.org/x/net/websocket"
)

var addr = flag.String("addr", ":8080", "http service address")

var gameState = gamestate.GameState{
	Resources: []gamestate.ResourceInterface{
		gamestate.NewResource("water"),
		// gamestate.NewWater(),
	},
	Tools: []gamestate.Tool{},
}

func websocketHandler(c echo.Context) error {
	websocket.Handler(func(ws *websocket.Conn) {
		defer ws.Close()

		go func() {
			for gs := range gameState.C {
				jsonData, err := json.Marshal(gs)
				if err != nil {
					c.Logger().Error(err)
					return
				}

				err = websocket.Message.Send(ws, string(jsonData))
				if err != nil {
					c.Logger().Error(err)
					return
				}
			}
		}()

		gameState.Init()

		for {
			// Read
			msg := ""
			err := websocket.Message.Receive(ws, &msg)
			if err != nil {
				c.Logger().Error(err)
				return
			}

			var cmd map[string]interface{}
			err = json.Unmarshal([]byte(msg), &cmd)
			if err != nil {
				c.Logger().Error(err)
				return
			}
			gameState.HandleCommand(cmd)
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
