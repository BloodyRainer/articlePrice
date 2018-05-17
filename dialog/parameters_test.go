package dialog

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

	new := AppendParameter(ps, "friend", "rainer")

	assert.Equal(t, `{"name":"thai", "friend":"rainer"}`, string(new))
}

func TestAppendParameter2(t *testing.T) {
	ps := MakeParameters("name", "thai")

	new := AppendParameter(ps, "friends", "2.00")

	assert.Equal(t, `{"name":"thai", "friends":2.00}`, string(new))
}

func TestAppendParameter3(t *testing.T) {
	ps := MakeParameters("name", "thai")

	new := AppendParameter(ps, "friends", "2")

	assert.Equal(t, `{"name":"thai", "friends":2}`, string(new))
}
