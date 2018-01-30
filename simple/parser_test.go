package simple

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestExtractField(t *testing.T) {
	// make some data
	e := Employee{
		Title: "Chaos Engineer",
		Person: &Person{
			Name:  "Melco Vich",
			Age:   67,
			Email: "chaos@fall.down",
		},
	}

	bz, err := e.Marshal()
	require.NoError(t, err)
	fmt.Printf("bz: %x\n", bz[:20])

	// now extract
	foo, err := extractField(bz, 3)
	assert.Error(t, err)
	assert.Nil(t, foo)

	field, err := extractField(bz, 1)
	assert.NoError(t, err)
	title, err := parseString(field)
	assert.Equal(t, e.Title, title)

	// var p Person
	// per, err := extractField(bz, 2)
	// assert.NoError(t, err)
	// err = proto.Unmarshal(per, &p)
	// assert.NoError(t, err)
	// assert.Equal(t, 67, p.Age)
	// assert.Equal(t, e.Person.Name, p.Name)
}
