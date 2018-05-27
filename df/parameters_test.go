package df

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestMakeParameters(t *testing.T) {
	ps := MakeParameters("name", "thai")

	assert.Equal(t, `{"name":"thai"}`, string(ps))
}

func TestAppendParameter1(t *testing.T) {
	ps := MakeParameters("name", "thai")

	new := AppendString(ps, "friend", "rainer")

	assert.Equal(t, `{"name":"thai", "friend":"rainer"}`, string(new))
}

func TestAppendParameter2(t *testing.T) {
	ps := MakeParameters("name", "thai")

	new := AppendNonString(ps, "friends", "2.00")

	assert.Equal(t, `{"name":"thai", "friends":2.00}`, string(new))
}

func TestAppendParameter3(t *testing.T) {
	ps := MakeParameters("name", "thai")

	new := AppendNonString(ps, "friends", "2")

	assert.Equal(t, `{"name":"thai", "friends":2}`, string(new))
}

func TestAppendParameter4(t *testing.T) {
	ps := MakeParameters("name", "test")

	new := AppendNonString(ps, "list", "[0.12, 12.99]")

	assert.Equal(t, `{"name":"test", "list":[0.12, 12.99]}`, string(new))
}
