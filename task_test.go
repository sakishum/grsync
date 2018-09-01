package grsync

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTask(t *testing.T) {
	t.Run("create new empty Task", func(t *testing.T) {
		createdTask := NewTask("a", "b", RsyncOptions{})

		assert.Empty(t, createdTask.Log(), "Task log should return empty string")
		assert.Empty(t, createdTask.State(), "Task should inited with empty state")
	})
}

func TestTaskProgressParse(t *testing.T) {
	progressMatcher := newMatcher(`\(.+-chk=(\d+.\d+)`)
	const taskInfoString = `999,999 99%  999.99kB/s    0:00:59 (xfr#9, to-chk=999/9999)`
	remain, total := getTaskProgress(progressMatcher.Extract(taskInfoString))

	assert.Equal(t, remain, 999)
	assert.Equal(t, total, 9999)
}

func TestTaskProgressWithDifferentChkID(t *testing.T) {
	progressMatcher := newMatcher(`\(.+-chk=(\d+.\d+)`)
	const taskInfoString = `999,999 99%  999.99kB/s    0:00:59 (xfr#9, ir-chk=999/9999)`
	remain, total := getTaskProgress(progressMatcher.Extract(taskInfoString))

	assert.Equal(t, remain, 999)
	assert.Equal(t, total, 9999)
}

func TestTaskSpeedParse(t *testing.T) {
	speedMatcher := newMatcher(`(\d+\.\d+.{2}\/s)`)
	const taskInfoString = `0.00kB/s \n 999,999 99%  999.99kB/s    0:00:59 (xfr#9, ir-chk=999/9999)`
	speed := getTaskSpeed(speedMatcher.ExtractAllStringSubmatch(taskInfoString, 2))
	assert.Equal(t, "999.99kB/s", speed)
}
