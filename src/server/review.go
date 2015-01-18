package server

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"gopkg.in/yaml.v2"
)

type Pic struct {
	Path    string `yaml:"Path"`
	Caption string `yaml:"Caption"`
}

type Thought struct {
	Body    template.HTML `yaml:"Body"`
	Thinker string        `yaml:"Thinker"`
	Pics    []Pic         `yaml:"Pics"`
}

type Review struct {
	Title    string    `yaml:"Title"`
	Location string    `yaml:"Location"`
	Score    float32   `yaml:"Score"`
	Thoughts []Thought `yaml:"Thoughts"`
}

var reviewTemplate, _ = template.ParseFiles("templates/review.html")

func reviewHandler(w http.ResponseWriter, r *http.Request) {
	theCoffee := strings.Split(r.URL.Path, "/")
	distributer := theCoffee[3]
	blend := theCoffee[4]
	review := getReview(distributer, blend)
	log.Println(review)
	reviewTemplate.Execute(w, review)
}

func getReview(distributer, blend string) Review {
	var review Review
	sickBytes, err := ioutil.ReadFile("./reviews/" + distributer + "/" + blend + ".yml")
	if err != nil {
		panic(err)
	}
	erry := yaml.Unmarshal(sickBytes, &review)
	if erry != nil {
		panic(erry)
	}
	return review
}
