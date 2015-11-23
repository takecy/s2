package main

import (
	"fmt"
	"os"

	"github.com/takecy/s2"
)

type Fuga struct {
	City   string `json:"city"`
	Number int64
}

type Hoge struct {
	ID   string `json:"id"`
	Name string
	Age  int64
	Fuga Fuga
}

func main() {
	id := "id001"
	name := "name001"
	age := int64(88)
	city := "kasuya"
	number := int64(12345)

	s := Hoge{
		ID:   id,
		Name: name,
		Age:  age,
		Fuga: Fuga{
			City:   city,
			Number: number,
		},
	}

	m := exampleToMap(s)
	exampleFromMap(m)
}

func exampleToMap(s Hoge) map[string]interface{} {
	m, err := s2.ToMap("json", s)

	fmt.Fprintf(os.Stderr, "%+v\n", err)
	fmt.Fprintf(os.Stdout, "%+v\n", m)

	return m
}

func exampleFromMap(m map[string]interface{}) {
	s := &Hoge{}
	err := s2.FromMap(m, s)

	fmt.Fprintf(os.Stderr, "%+v\n", err)
	fmt.Fprintf(os.Stdout, "%+v\n", s)
}
