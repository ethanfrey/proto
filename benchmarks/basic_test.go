package benchmarks

import (
	"testing"

	"github.com/gogo/protobuf/proto"

	wire "github.com/tendermint/go-wire"

	"github.com/ethanfrey/proto/options"
	"github.com/ethanfrey/proto/simple"
)

func makePerson() *simple.Person {
	return &simple.Person{
		Name:  "John Doe",
		Age:   42,
		Email: "john.doe@aol.com",
	}
}

func makeBook() *simple.Book {
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

	return &simple.Book{
		Phones: []*simple.PhoneNumber{&p1, &p2, &p3},
	}
}

func makeResponse() *options.Response {
	return &options.Response{
		Error: false,
		Data:  &options.Bytes{0x42, 0x00, 0xCA, 0xFE, 0x00},
		Log:   "drink some mocha",
	}
}

func benchmarkWireUnmarshal(b *testing.B, in, out interface{}) {
	data := wire.BinaryBytes(in)
	for i := 0; i < b.N; i++ {
		err := wire.ReadBinaryBytes(data, out)
		if err != nil {
			panic(err)
		}
	}
}

func benchmarkProtoUnmarshal(b *testing.B, in, out proto.Message) {
	data, err := proto.Marshal(in)
	if err != nil {
		panic(err)
	}

	for i := 0; i < b.N; i++ {
		err := proto.Unmarshal(data, out)
		if err != nil {
			panic(err)
		}
	}
}

func benchmarkWireMarshal(b *testing.B, in interface{}) {
	for i := 0; i < b.N; i++ {
		data := wire.BinaryBytes(in)
		if len(data) < 3 {
			panic("no data")
		}
	}
}

func benchmarkProtoMarshal(b *testing.B, in proto.Message) {
	for i := 0; i < b.N; i++ {
		data, err := proto.Marshal(in)
		if err != nil {
			panic(err)
		}
		if len(data) < 3 {
			panic("no data")
		}
	}
}

func BenchmarkUnmarshal(b *testing.B) {
	cases := []struct {
		name    string
		in, out proto.Message
		wire    interface{}
	}{
		// cuz, of course, wire is "special" with unmarshalling pointers
		{"person", makePerson(), new(simple.Person), new(*simple.Person)},
		{"book", makeBook(), new(simple.Book), new(*simple.Book)},
		{"response", makeResponse(), new(options.Response), new(*options.Response)},
	}

	for _, tc := range cases {
		b.Run(tc.name+"-proto", func(sub *testing.B) {
			benchmarkProtoUnmarshal(sub, tc.in, tc.out)
		})
		b.Run(tc.name+"-wire", func(sub *testing.B) {
			benchmarkWireUnmarshal(sub, tc.in, tc.wire)
		})
	}
}

func BenchmarkMarshal(b *testing.B) {
	cases := []struct {
		name string
		in   proto.Message
	}{
		// cuz, of course, wire is "special" with unmarshalling pointers
		{"person", makePerson()},
		{"book", makeBook()},
		{"response", makeResponse()},
	}

	for _, tc := range cases {
		b.Run(tc.name+"-proto", func(sub *testing.B) {
			benchmarkProtoMarshal(sub, tc.in)
		})
		b.Run(tc.name+"-wire", func(sub *testing.B) {
			benchmarkWireMarshal(sub, tc.in)
		})
	}
}
