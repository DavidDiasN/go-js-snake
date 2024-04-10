package main

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/DavidDiasN/htmx-snake"
	"github.com/gorilla/websocket"
)

var wsUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {

	http.HandleFunc("/snake", func(w http.ResponseWriter, r *http.Request) {
		myBoard := make([][]rune, 25)
		for i := range myBoard {
			myBoard[i] = make([]rune, 25)
		}
		component := squares(myBoard)
		component.Render(context.Background(), w)
		fmt.Println("Done making snake screen")

	})

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Successful websocket connection")

		conn, err := wsUpgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Printf("Problem upgrading connection to webSockets %v\n", err)
		}

		_, mess, err := conn.ReadMessage()

		conn.WriteJSON("Hello")
		fmt.Println(mess)


    wrapperConn := 
		connectionBoard := board.NewGame(25, 25)

		quit := make(chan bool)
		// Give Frame Sender a wrtier dependency
		// create a wrapper write function that you can give to FrameSender

		go connectionBoard.FrameSender(quit)

		go func() {
			select {
			case res := <-output:
				conn.WriteJSON(res)
			}
		}()

		go func() {

			_, newMessage, err := conn.ReadMessage()
			if err != nil {
				fmt.Println("error")
			} else {

				fmt.Println(newMessage[1])
			}
			// Give move listener a reader dependency
			// create a wrapper read function that you can give to MoveListener
			err = connectionBoard.MoveListener(quit)

			if err == board.UserClosedGame {
				fmt.Println("User closed the game")
			}
		}()

		for i := 0; i < 23; i++ {
		}
		fmt.Println("Connection terminated")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

type boardGame struct {
	http.Handler
	template   *template.Template
	snakeBoard board.Board
}

type wrapperConn struct {
  realConn *websocket.Conn  
}

func (s *simpleConn) Read() {
  return realConn.ReadMessage(readBuffer)
}

func (s *simpleConn) Write(data any) {

}
