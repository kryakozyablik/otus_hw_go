package main

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRunCmd(t *testing.T) {
	expectedEnvList := Environment{"PATH": "replaced_path", "NEW_ENV": "new_env_data"}
	cmd := []string{"env"}

	oldStdOut := os.Stdout
	rStdO, wStdO, err := os.Pipe()
	require.NoError(t, err)
	os.Stdout = wStdO

	returnCode := RunCmd(cmd, expectedEnvList)
	require.Equal(t, 0, returnCode)

	err = wStdO.Close()
	require.NoError(t, err)

	os.Stdout = oldStdOut

	result, err := ioutil.ReadAll(rStdO)
	require.NoError(t, err)
	require.NotEmpty(t, result)

	envList := parseEnvString(string(result))

	require.Equal(t, expectedEnvList["PATH"], envList["PATH"])
	require.Equal(t, expectedEnvList["NEW_ENV"], envList["NEW_ENV"])
}

func parseEnvString(envString string) map[string]string {
	envString = strings.TrimRight(envString, "\n")
	envList := make(map[string]string)
	for _, envLine := range strings.Split(envString, "\n") {
		envMap := strings.Split(envLine, "=")
		envList[envMap[0]] = envMap[1]
	}
	return envList
}
