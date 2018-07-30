package grsync

import (
	"grsync/rsync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTask(t *testing.T) {
	t.Run("create new empty Task", func(t *testing.T) {
		createdTask := New("a", "b", rsync.Options{})

		assert.Empty(t, createdTask.Log(), "Task log should return empty string")
		assert.Empty(t, createdTask.State(), "Task should inited with empty state")
	})
}
