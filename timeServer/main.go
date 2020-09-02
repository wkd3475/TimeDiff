package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	ws "github.com/gorilla/websocket"
)

var (
	timeseverConnCtx = context.Background()
)

var upgrader = ws.Upgrader{
	CheckOrigin: func(_ *http.Request) (ok bool) {
		ok = true
		return
	},
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	var (
		err    error
		c      *ws.Conn
		cancel context.CancelFunc
	)

	timeseverConnCtx, cancel = context.WithCancel(context.Background())
	defer cancel()

	if c, err = upgrader.Upgrade(w, r, nil); err != nil {
		log.Println("failed to upgrade", err)
		return
	}
	defer c.Close()

	for {
		// receive data
		_, data, err := c.ReadMessage()
		t2 := strconv.FormatInt(time.Now().UnixNano()/1e3, 10)
		s := strings.Split(string(data), ",")
		t1 := string(s[1])
		if err != nil {
			log.Println("failed to read")
			break
		}
		datas := []byte(s[0] + "," + t1 + "," + t2 + "," + strconv.FormatInt(time.Now().UnixNano()/1e3, 10))
		if err := c.WriteMessage(ws.TextMessage, datas); err != nil {
			log.Println("failed to write")
			break
		}
		fmt.Println(string(datas))
	}

	cancel()
}

func setupRoutes() {
	http.HandleFunc("/time", timeHandler)
	if err := http.ListenAndServe(":9090", nil); err != nil {
		log.Println("failed to listend and serve")
	}
}

func main() {
	setupRoutes()
}
