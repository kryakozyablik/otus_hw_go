package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadDir(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
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
	})

	t.Run("with dir", func(t *testing.T) {
		envList, err := ReadDir("testdata/env_with_dir")
		require.Error(t, err, ErrDirExists)
		require.Nil(t, envList)
	})

	t.Run("with dir", func(t *testing.T) {
		envList, err := ReadDir("testdata/env_with_invalid_file")
		require.Error(t, err, ErrInvalidFile)
		require.Nil(t, envList)
	})
}
