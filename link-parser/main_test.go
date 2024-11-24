package main

import (
	"reflect"
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func TestExtractText(t *testing.T) {
	var testsCases = []struct {
		input    string
		expected string
	}{
		{input: "<a href='/lost'>Lost? Need help?</a>", expected: "Lost? Need help?"},
		{input: "<a href='https://github.com/gophercises'>Gophercises is on <strong>Github</strong>!</a>", expected: "Gophercises is on Github!"},
	}

	for _, test := range testsCases {
		doc, err := html.Parse(strings.NewReader(test.input))
		if err != nil {
			t.Fatal(err)
		}

		actual := extractText(doc)
		if actual != test.expected {
			t.Errorf("expected %q but got %q", test.expected, actual)
		}
	}
}

func TestProcessFile(t *testing.T) {

	var testsCases = []struct {
		input    string
		expected []Link
	}{
		{input: "ex1.html", expected: []Link{
			{Href: "/other-page", Text: "A link to another page "},
		}},
		{input: "ex2.html", expected: []Link{
			{Href: "https://www.twitter.com/joncalhoun", Text: "Check me out on twitter"},
			{Href: "https://github.com/gophercises", Text: "Gophercises is on Github!"},
		}},
		{input: "ex3.html", expected: []Link{
			{Href: "#", Text: "Login"},
			{Href: "https://twitter.com/marcusolsson", Text: "@marcusolsson"},
			{Href: "/lost", Text: "Lost? Need help?"},
		}},
	}

	for _, test := range testsCases {
		doc, err := ReadFileAndParseFile(test.input)
		if err != nil {
			t.Fatal(err)
		}

		actual := ProcessFile(doc)
		if !reflect.DeepEqual(actual, test.expected) {
			t.Errorf("expected %+v but got %+v", test.expected, actual)
		}
	}
}
