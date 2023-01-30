package templaterule

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestURLQueryMap(t *testing.T) {
	s, err := urlquerymap(map[string]int{"a": 1, "b": 2})
	assert.NoError(t, err)
	assert.Equal(t, "a=1&b=2", s)

	var m map[string]string
	s, err = urlquerymap(m)
	assert.NoError(t, err)
	assert.Equal(t, "", s)

	_, err = urlquerymap(3)
	assert.EqualError(t, err, "urlquerymap expected a map but got (int)3")
}

func TestSelector(t *testing.T) {
	assert.Equal(t, "a=1,b=2,c=3", selector(map[string]int{"a": 1, "b": 2, "c": 3}))
	assert.Equal(t, "", selector(nil))
	assert.Equal(t, "", selector(map[int]int{}))
}

func TestKVMap(t *testing.T) {
	assert.Equal(t, map[string]any{"a": 1, "b": 2, "c": 3}, kvMap("a", 1, "b", 2, "c", 3))
	assert.Panics(t, func() { kvMap("x") })
	assert.Equal(t, map[string]any(nil), kvMap())
}
