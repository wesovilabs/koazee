package stream_test

import (
	"strings"
	"testing"

	"github.com/wesovilabs/koazee/stream"

	"github.com/stretchr/testify/assert"
)

func TestStream_Map(t *testing.T) {
	s := stream.New([]string{"HomE", "welcome", "SoftWare", "tech", "GEEK"})
	out := s.Map(func(element string) string {
		return strings.ToUpper(element)
	}).Out().Val()

	assert.Equal(t, []string{"HOME", "WELCOME", "SOFTWARE", "TECH", "GEEK"}, out)
	s = stream.New([]*person{
		{
			firstName: "John",
			age:       35,
		},
		{
			firstName: "Jane",
			age:       60,
		},
		{
			firstName: "Jean",
			age:       23,
		},
	})
	out = s.Map(func(p *person) string {
		return p.firstName
	}).Out().Val()
	assert.Equal(t, []string{"John", "Jane", "Jean"}, out)
}
