package main

import (
	//"fmt"
	"github.com/DavidDiasN/tcp_server_app"
	"html/template"
	"log"
	"net/http"
)

type Test struct {
	HTML     string
	SafeHTML template.HTML
	Title    string
	Path     string
	Dog      Dog
	Map      map[string]string
}

type Dog struct {
	Name string
	Age  int
}

func main() {

	http.HandleFunc("/snake", func(w http.ResponseWriter, r *http.Request) {
		quit := make(chan bool)

		connectionBoard := board.NewGame(25, 25)
		// add more channels to catch errors
		go connectionBoard.FrameSender(quit)

		err := connectionBoard.MoveListener(quit)
		if err == board.UserClosedGame {
			fmt.Println("User closed the game")
		}

		t, err := template.ParseFiles("templates/context.gohtml")
		if err != nil {
			panic(err)
		}

		data := Test{
			HTML:     "<h1>A header!</h1>",
			SafeHTML: template.HTML("<h1>A Safe header</h1>"),
			Title:    "Backslash! an in depth look at the \"\\\" character.",
			Path:     "/dashboard/settings",
			Dog:      Dog{"Fido", 6},
			Map: map[string]string{
				"key":       "value",
				"other_key": "other_value",
			},
		}

		err = t.Execute(w, data)
		if err != nil {
			panic(err)
		}

	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
