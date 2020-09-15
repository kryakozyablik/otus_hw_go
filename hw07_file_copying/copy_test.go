package main

import (
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

type test struct {
	offset int64
	limit  int64
}

func TestCopy(t *testing.T) {
	tmpFile, err := ioutil.TempFile("/tmp", "output")
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err := os.Remove(tmpFile.Name())
		if err != nil {
			log.Fatal(err)
		}
	}()

	t.Run("file copy test", func(t *testing.T) {
		for _, tst := range [...]test{
			{
				0,
				0,
			},
			{
				10,
				20,
			},
			{
				0,
				6603,
			},
			{
				6603,
				0,
			},
		} {

			err := Copy("testdata/input.txt", tmpFile.Name(), tst.offset, tst.limit)
			require.Nil(t, err)

			originalFile, err := ioutil.ReadFile("testdata/input.txt")
			require.Nil(t, err)

			testData := originalFile[tst.offset:]
			if tst.limit > 0 {
				testData = testData[:tst.limit]
			}

			fileCopy, err := ioutil.ReadFile(tmpFile.Name())
			require.Nil(t, err)

			require.Equal(t, testData, fileCopy)
		}
	})

	t.Run("copy stream", func(t *testing.T) {
		err := Copy("/dev/urandom", tmpFile.Name(), 0, 0)
		require.Equal(t, err, ErrUnsupportedFile)
	})

	t.Run("file does not exists", func(t *testing.T) {
		err := Copy("/not/exists", tmpFile.Name(), 0, 0)
		require.Equal(t, err, ErrFileDoesNotExists)
	})

	t.Run("offset out of range", func(t *testing.T) {
		err := Copy("testdata/out_offset0_limit10.txt", tmpFile.Name(), 10000, 0)
		require.Equal(t, err, ErrOffsetExceedsFileSize)
	})

	t.Run("limit out of range", func(t *testing.T) {
		err := Copy("testdata/input.txt", tmpFile.Name(), 0, 100000)

		originalFile, err := ioutil.ReadFile("testdata/input.txt")
		require.Nil(t, err)

		fileCopy, err := ioutil.ReadFile(tmpFile.Name())
		require.Nil(t, err)

		require.Equal(t, originalFile, fileCopy)
	})
}
