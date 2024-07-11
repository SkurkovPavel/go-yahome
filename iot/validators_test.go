package iot

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGroupValidator_New(t *testing.T) {

	g := Group{}

	t.Run("EMPTY ID", func(t *testing.T) {
		err := g.Validate()
		require.Error(t, err)
	})
	t.Run("EMPTY ACTION", func(t *testing.T) {
		g.Id = "test"
		err := g.Validate()
		require.Error(t, err)
	})
	t.Run("OK", func(t *testing.T) {
		g.Actions = []Capability{{
			Type: "test",
		}}
		err := g.Validate()
		require.NoError(t, err)
	})
}

func TestDeviceValidator_New(t *testing.T) {

	d := Device{}
	t.Run("EMPTY ID", func(t *testing.T) {
		err := d.Validate()
		require.Error(t, err)
	})
	t.Run("EMPTY ACTION", func(t *testing.T) {
		d.Id = "test"
		err := d.Validate()
		require.Error(t, err)
	})
	t.Run("EMPTY TYPE", func(t *testing.T) {
		d.Actions = []Capability{{
			Type: "test",
		}}
		err := d.Validate()
		require.Error(t, err)
	})
	t.Run("EMPTY STATE", func(t *testing.T) {
		d.Type = "test"
		err := d.Validate()
		require.Error(t, err)
	})
	t.Run("OK", func(t *testing.T) {
		d.State = "test"
		err := d.Validate()
		require.NoError(t, err)
	})
}

func TestErrReturner_New(t *testing.T) {

	method := "test"
	tErr := errors.New("test")

	expErr := ErrorReturn(method, tErr)

	require.Error(t, expErr)
	assert.Equal(t, expErr, fmt.Errorf("%s error: %s", method, tErr.Error()))
}
