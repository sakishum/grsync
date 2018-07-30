package rsync

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseArguments(t *testing.T) {
	t.Run("--verbose", func(t *testing.T) {
		args := getArguments(Options{
			Verbose: true,
		})
		assert.Contains(t, args, "--verbose")
	})

	t.Run("--checksum", func(t *testing.T) {
		args := getArguments(Options{
			Checksum: true,
		})
		assert.Contains(t, args, "--checksum")
	})

	t.Run("--quiet", func(t *testing.T) {
		args := getArguments(Options{
			Quiet: true,
		})
		assert.Contains(t, args, "--quiet")
	})

	t.Run("--archive", func(t *testing.T) {
		args := getArguments(Options{
			Archive: true,
		})
		assert.Contains(t, args, "--archive")
	})

	t.Run("--recursive", func(t *testing.T) {
		args := getArguments(Options{
			Recursive: true,
		})
		assert.Contains(t, args, "--recursive")
	})

	t.Run("--relative", func(t *testing.T) {
		args := getArguments(Options{
			Relative: true,
		})
		assert.Contains(t, args, "--relative")
	})

	t.Run("--no-implied-dirs", func(t *testing.T) {
		args := getArguments(Options{
			NoImpliedDirs: true,
		})
		assert.Contains(t, args, "--no-implied-dirs")
	})

	t.Run("--update", func(t *testing.T) {
		args := getArguments(Options{
			Update: true,
		})
		assert.Contains(t, args, "--update")
	})

	t.Run("--inplace", func(t *testing.T) {
		args := getArguments(Options{
			Inplace: true,
		})
		assert.Contains(t, args, "--inplace")
	})

	t.Run("--append", func(t *testing.T) {
		args := getArguments(Options{
			Append: true,
		})
		assert.Contains(t, args, "--append")
	})
}
