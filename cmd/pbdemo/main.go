package main

import (
	"encoding/json"
	"fmt"

	"github.com/ethanfrey/proto/simple"
	"github.com/gogo/protobuf/proto"
)

func main() {
	p := &simple.Person{
		Name:  "John Doe",
		Age:   42,
		Email: "john.doe@aol.com",
	}
	data, err := proto.Marshal(p)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Encoded Size: %d\n", len(data))

	in := new(simple.Person)
	err = proto.Unmarshal(data, in)
	if err != nil {
		panic(err)
	}

	if in.Name != p.Name || in.Age != p.Age || in.Email != p.Email {
		fmt.Printf("Input data incorrect: %#v\n", in)
		return
	}

	js, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(js))
}
