package main

import (
	"errors"
	"fmt"
	"io"
	"math"
	"os"
	"strings"
)

var (
	ErrFileDoesNotExists     = errors.New("file does not exists")
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

func Copy(fromPath string, toPath string, offset, limit int64) error {
	info, err := os.Stat(fromPath)
	if os.IsNotExist(err) || info.IsDir() {
		return ErrFileDoesNotExists
	}

	if info.Size() == 0 {
		return ErrUnsupportedFile
	}

	if offset > info.Size() {
		return ErrOffsetExceedsFileSize
	}

	fFrom, err := os.Open(fromPath)
	if err != nil {
		return err
	}

	fTo, err := os.Create(toPath)
	if err != nil {
		return err
	}

	if limit == 0 {
		limit = info.Size()
	}

	_, err = fFrom.Seek(offset, 0)
	if err != nil {
		return err
	}

	copySize := info.Size() - offset

	if limit < copySize {
		copySize = limit
	}

	pReader := &progressReader{R: fFrom, CopySize: copySize}

	_, err = io.CopyN(fTo, pReader, limit)
	if err != nil && err != io.EOF {
		return err
	}

	return nil
}

type progressReader struct {
	R        io.Reader
	CopySize int64
	readed   int64
}

func (pr *progressReader) Read(p []byte) (int, error) {
	n, err := pr.R.Read(p)

	if err == nil {
		pr.printProgress(n)
	}

	return n, err
}

func (pr *progressReader) printProgress(rs int) {
	pr.readed += int64(rs)
	percent := int(math.Round(float64(pr.readed) * 100 / float64(pr.CopySize)))

	bar := fmt.Sprintf(
		"\r|%s%s%s| %d%%",
		strings.Repeat("=", percent),
		">",
		strings.Repeat("-", 100-percent),
		percent,
	)

	fmt.Print(bar)

	if pr.readed == pr.CopySize {
		fmt.Println()
	}
}
