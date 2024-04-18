package main

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	board "github.com/DavidDiasN/htmx-snake"
	"github.com/gorilla/websocket"
)

var wsUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		component := mainPage()
		component.Render(context.Background(), w)
	})

	http.HandleFunc("/snake", func(w http.ResponseWriter, r *http.Request) {
		component := squares(25)
		component.Render(context.Background(), w)
		fmt.Println("Done making snake screen")

	})

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Successful websocket connection")

		conn, err := wsUpgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Printf("Problem upgrading connection to webSockets %v\n", err)
		}

		wrapperConn := wrapperConn{conn}
		connectionBoard := board.NewGame(25, 25, wrapperConn)

		quit := make(chan bool)

		go connectionBoard.FrameSender(quit)

		go func() {

			err = connectionBoard.MoveListener(quit)

			if err == board.UserClosedGame {
				fmt.Println("User closed the game")
			}
		}()
	})

	http.HandleFunc("/scripts/snake.js", func(w http.ResponseWriter, r *http.Request) {
		file, err := os.Open("scripts/snake.js")

		if err != nil {
			fmt.Println("open error")
			panic(err)
		}

		info, err := file.Stat()

		if err != nil {
			fmt.Println("stat error")
			panic(err)
		}
		data := make([]byte, info.Size())
		_, err = file.Read(data)
		if err != nil {
			fmt.Println("Read error")
			panic(err)
		}

		w.Header().Add("Content-Type", "text/javascript")
		w.Write(data)

	})

	http.HandleFunc("/styles/snake.css", func(w http.ResponseWriter, r *http.Request) {

		file, err := os.Open("styles/snake.css")

		if err != nil {
			fmt.Println("open error")
			panic(err)
		}

		info, err := file.Stat()

		if err != nil {
			fmt.Println("stat error")
			panic(err)
		}
		data := make([]byte, info.Size())
		_, err = file.Read(data)
		if err != nil {
			fmt.Println("Read error")
			panic(err)
		}

		w.Header().Add("Content-Type", "text/css")
		w.Write(data)
	})

	http.HandleFunc("/styles/main.css", func(w http.ResponseWriter, r *http.Request) {

		file, err := os.Open("styles/main.css")

		if err != nil {
			fmt.Println("open error")
			panic(err)
		}

		info, err := file.Stat()

		if err != nil {
			fmt.Println("stat error")
			panic(err)
		}
		data := make([]byte, info.Size())
		_, err = file.Read(data)
		if err != nil {
			fmt.Println("Read error")
			panic(err)
		}

		http.HandleFunc("/assets/resume.pdf", func(w http.ResponseWriter, r *http.Request) {

			file, err := os.Open("/assets/resume.pdf")

			if err != nil {
				fmt.Println("open error")
				panic(err)
			}

			info, err := file.Stat()

			if err != nil {
				fmt.Println("stat error")
				panic(err)
			}
			data := make([]byte, info.Size())
			_, err = file.Read(data)
			if err != nil {
				fmt.Println("Read error")
				panic(err)
			}

			w.Header().Add("Content-Type", "application/pdf")
			w.Write(data)
		})

		w.Header().Add("Content-Type", "text/css")
		w.Write(data)
	})

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

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

func (w wrapperConn) Read() (messageLen int, message []byte, err error) {
	return w.realConn.ReadMessage()
}

func (w wrapperConn) Write(data interface{}) error {
	return w.realConn.WriteJSON(data)
}
