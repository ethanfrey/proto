package simple

import (
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

	// now extract
	foo, err := ExtractField(bz, 3)
	assert.Error(t, err)
	assert.Nil(t, foo)

	field, err := ExtractField(bz, 1)
	assert.NoError(t, err)
	title, err := ParseString(field)
	assert.Equal(t, e.Title, title)

	var p Person
	field, err = ExtractField(bz, 2)
	require.NoError(t, err)
	err = ParseStruct(field, &p)
	assert.NoError(t, err)
	assert.Equal(t, int32(67), p.Age)
	assert.Equal(t, e.Person.Name, p.Name)
}
