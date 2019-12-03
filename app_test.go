package app

import (
	"testing"

	"gotest.tools/assert"
)

func TestInit(t *testing.T) {
	Init()
}

func TestVersion(t *testing.T) {
	version()
}

func TestFullName(t *testing.T) {
	expected := "grace-development-empty"
	actual := FullName()
	assert.Equal(t, actual, expected)
}

func TestName(t *testing.T) {
	expected := "empty"
	actual := Name()
	assert.Equal(t, actual, expected)
}

func TestEnv(t *testing.T) {
	expected := "development"
	actual := Env()
	assert.Equal(t, actual, expected)
}
