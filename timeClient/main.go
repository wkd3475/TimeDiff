package main

import (
	"fmt"
	"log"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)

var (
	total_o int64
)

func main() {
	total_o = 0
	// u := url.URL{Scheme: "ws", Host: "localhost:9090", Path: "/time"}
	u := url.URL{Scheme: "ws", Host: "192.168.80.148:9090", Path: "/time"}
	fmt.Printf("connecting to %s\n", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			t4 := time.Now().UnixNano() / 1e3
			if err != nil {
				fmt.Println("read:", err)
				return
			}
			s := strings.Split(string(message), ",")
			t1, _ := strconv.ParseInt(s[1], 10, 64)
			t2, _ := strconv.ParseInt(s[2], 10, 64)
			t3, _ := strconv.ParseInt(s[3], 10, 64)

			o := (t2 + t3 - t1 - t4) / 2
			total_o += o / 1e3
		}
	}()

	for i := 1; i <= 100; i++ {
		time.Sleep(time.Millisecond * 500)
		datas := []byte(strconv.Itoa(i) + "," + strconv.FormatInt(time.Now().UnixNano()/1e3, 10))
		err := c.WriteMessage(websocket.TextMessage, datas)
		fmt.Println(string(datas))
		if err != nil {
			fmt.Println("write:", err)
			return
		}
	}
	time.Sleep(time.Millisecond * 1000)

	fmt.Println(total_o / 100)
}
