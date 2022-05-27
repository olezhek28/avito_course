package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/udhos/equalfile"
)

func TestCopy(t *testing.T) {
	var (
		inputFile       = "testdata/input.txt"
		outputFile      = "tmp.txt"
		unsupportedFile = "/dev/urandom"

		cmp = equalfile.New(nil, equalfile.Options{}) // compare using single mode
	)

	t.Run("copy full file (offset = 0; limit = 0)", func(t *testing.T) {
		err := Copy(inputFile, outputFile, 0, 0)
		require.NoError(t, err)

		equal, err := cmp.CompareFile(outputFile, "testdata/out_offset0_limit0.txt")
		require.NoError(t, err)
		require.True(t, equal)

		os.Remove(outputFile)
	})

	t.Run("copy with (offset = 0; limit = 10)", func(t *testing.T) {
		err := Copy(inputFile, outputFile, 0, 10)
		require.NoError(t, err)

		equal, err := cmp.CompareFile(outputFile, "testdata/out_offset0_limit10.txt")
		require.NoError(t, err)
		require.True(t, equal)

		os.Remove(outputFile)
	})

	t.Run("copy with (offset = 0; limit = 1000)", func(t *testing.T) {
		err := Copy(inputFile, outputFile, 0, 1000)
		require.NoError(t, err)

		equal, err := cmp.CompareFile(outputFile, "testdata/out_offset0_limit1000.txt")
		require.NoError(t, err)
		require.True(t, equal)

		os.Remove(outputFile)
	})

	t.Run("copy with (offset = 0; limit = 10'000)", func(t *testing.T) {
		err := Copy(inputFile, outputFile, 0, 10_000)
		require.NoError(t, err)

		equal, err := cmp.CompareFile(outputFile, "testdata/out_offset0_limit10000.txt")
		require.NoError(t, err)
		require.True(t, equal)

		os.Remove(outputFile)
	})

	t.Run("copy with (offset = 100; limit = 1000)", func(t *testing.T) {
		err := Copy(inputFile, outputFile, 100, 1000)
		require.NoError(t, err)

		equal, err := cmp.CompareFile(outputFile, "testdata/out_offset100_limit1000.txt")
		require.NoError(t, err)
		require.True(t, equal)

		os.Remove(outputFile)
	})

	t.Run("copy with (offset = 6000; limit = 1000)", func(t *testing.T) {
		err := Copy(inputFile, outputFile, 6000, 1000)
		require.NoError(t, err)

		equal, err := cmp.CompareFile(outputFile, "testdata/out_offset6000_limit1000.txt")
		require.NoError(t, err)
		require.True(t, equal)

		os.Remove(outputFile)
	})

	t.Run("copy with offset great that file size", func(t *testing.T) {
		fileInfo, err := os.Stat(inputFile)
		require.NoError(t, err)

		err = Copy(inputFile, outputFile, fileInfo.Size()+1, 1000)
		require.Error(t, err)
		require.Equal(t, ErrOffsetExceedsFileSize, err)
	})

	t.Run("copy unsupported file", func(t *testing.T) {
		err := Copy(unsupportedFile, outputFile, 0, 0)
		require.Error(t, err)
		require.Equal(t, ErrUnsupportedFile, err)
	})
}
