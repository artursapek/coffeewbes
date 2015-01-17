package server

import (
	"html/template"
	"net/http"
)

type Pic struct {
	Path, Caption string
}

type Thought struct {
	Body    string
	Thinker string
	Pics    []Pic
}

type Review struct {
	Title    string
	Location string
	Score    float32
	Thoughts []Thought
}

var reviewTemplate, _ = template.ParseFiles("templates/review.html")

func reviewHandler(w http.ResponseWriter, r *http.Request) {

	review := Review{
		Title:    "Coffee review #1",
		Location: "NYC",
		Score:    4.5,
		Thoughts: []Thought{
			Thought{
				Body:    "thought #1",
				Thinker: "mb",
			},
			Thought{
				Body:    "thought #2",
				Thinker: "mb",
			},
		},
	}

	reviewTemplate.Execute(w, review)
}
