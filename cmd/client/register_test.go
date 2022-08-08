package main //nolint:testpackage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {
	t.Parallel()

	assert := assert.New(t)

	regCmd := NewRegistration()

	assert.NotNil(regCmd)
}

func TestRegisterReq(t *testing.T) {
	t.Parallel()

	assert := assert.New(t)

	regcmd := NewRegistration()
	rootcmd := New()
	args := []string{"registration", "test-case"}

	rootcmd.AddCommand(regcmd)

	regreq := RegisterReq(rootcmd, args)

	assert.Nil(regreq)
}
