package main

import (
	"encoding/json"
	"fmt"

	"github.com/ethanfrey/proto/simple"
	"github.com/gogo/protobuf/proto"
	wire "github.com/tendermint/go-wire"
)

// Equaler lets us find objects that can be compared
type Equaler interface {
	proto.Message
	Equal(that interface{}) bool
}

func TrialEncodings(in, out Equaler) error {
	data, err := proto.Marshal(in)
	if err != nil {
		return err
	}
	fmt.Printf("Protobuf Size: %d\n", len(data))

	wdata := wire.BinaryBytes(in)
	fmt.Printf("Wire Size: %d\n", len(wdata))

	err = proto.Unmarshal(data, out)
	if err != nil {
		return err
	}

	if !in.Equal(out) {
		return fmt.Errorf("Loaded data doesn't match input")
	}

	js, err := json.MarshalIndent(out, "", "  ")
	if err != nil {
		return err
	}
	fmt.Println(string(js))
	return nil
}

func main() {
	in := &simple.Person{
		Name:  "John Doe",
		Age:   42,
		Email: "john.doe@aol.com",
	}
	out := new(simple.Person)

	fmt.Println("\n--> simple.Person")
	err := TrialEncodings(in, out)
	if err != nil {
		panic(err)
	}

	p1 := simple.PhoneNumber{
		Number: "1234567890",
		Type:   simple.MOBILE,
	}
	p2 := simple.PhoneNumber{
		Number: "abcdefg",
		Type:   simple.HOME,
	}
	p3 := simple.PhoneNumber{
		Number: "678ftw890",
		Type:   simple.WORK,
	}

	bin := &simple.Book{
		Phones: []*simple.PhoneNumber{&p1, &p2, &p3},
	}
	bout := new(simple.Book)

	fmt.Println("\n--> simple.Book")
	err = TrialEncodings(bin, bout)
	if err != nil {
		panic(err)
	}

}
