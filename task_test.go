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

func TestTaskProgressParse(t *testing.T) {
	const taskInfoString = `999,999 99%  999.99kB/s    0:00:59 (xfr#9, to-chk=999/9999)`
	remain, total := getTaskProgress(progressMatcher.Extract(taskInfoString))

	assert.Equal(t, remain, 999)
	assert.Equal(t, total, 9999)
}
