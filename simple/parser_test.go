package simple

import (
	"testing"

	"github.com/gogo/protobuf/proto"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func ParseStruct(bz []byte, pb proto.Message) error {
	field, err := ParseBytesField(bz)
	if err != nil {
		return err
	}
	return proto.Unmarshal(field, pb)
}

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

	field, err = ExtractPath(bz, 1)
	assert.NoError(t, err)
	title, err = ParseString(field)
	assert.Equal(t, e.Title, title)

	// get the age, then the email
	field, err = ExtractPath(bz, 2, 2)
	assert.NoError(t, err)
	age, _, err := ParseInt32(field)
	assert.Equal(t, e.Person.Age, age)

	field, err = ExtractPath(bz, 2, 3)
	assert.NoError(t, err)
	email, err := ParseString(field)
	assert.Equal(t, e.Person.Email, email)

}
