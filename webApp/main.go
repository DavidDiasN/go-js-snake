package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/DavidDiasN/htmx-snake"
)

func main() {

	http.HandleFunc("/snake", func(w http.ResponseWriter, r *http.Request) {
		quit := make(chan bool)
		output := make(chan []byte)
		connectionBoard := board.NewGame(25, 25)
		// add more channels to catch errors
		go connectionBoard.FrameSender(quit, output)
		select {
		case res := <-output:
			fmt.Fprint(w, res)
		}
		err := connectionBoard.MoveListener(quit)
		arr := []rune{'d', 's', 'a', 'w', 'd', 's', 'a'}
		for i := 0; i < 10; i++ {
			if i < 6 {
				connectionBoard.WriteUserInput(arr[i])
			}
			time.Sleep(50 * time.Millisecond)
		}
		fmt.Println("I am a loser that blocks")

		if err == board.UserClosedGame {
			fmt.Println("User closed the game")
		}

	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

type boardGame struct {
	snakeBoard board.Board
	http.Handler
	template *template.Template
}
