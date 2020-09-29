package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadDir(t *testing.T) {
	envList, err := ReadDir("testdata/env")

	expectedEnvList := Environment{
		"BAR":   "bar",
		"FOO":   "   foo\nwith new line",
		"HELLO": "\"hello\"",
		"UNSET": "",
	}
	require.NoError(t, err)

	equal := reflect.DeepEqual(expectedEnvList, envList)
	require.True(t, equal)
}
