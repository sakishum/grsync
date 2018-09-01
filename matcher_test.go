package grsync

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMatcherMatch(t *testing.T) {
	t.Run("should match number", func(t *testing.T) {
		const strToMatch = "hello 1232"
		m := newMatcher(`(\d+)`)
		assert.True(t, m.Match(strToMatch))
	})

	t.Run("should match substring", func(t *testing.T) {
		const strToMatch = "hello Denis!"
		m := newMatcher(`hello (\w+)!`)
		assert.True(t, m.Match(strToMatch))
	})

	t.Run("should match speed", func(t *testing.T) {
		const strToMatch = "          1.05M 100%  659.30kB/s    0:00:01 (xfr#5, ir-chk=3641/3679)"
		m := newMatcher(`(\d+\.\d+.{2}/s)`)
		assert.True(t, m.Match(strToMatch))
	})
}

func TestMatcherExtract(t *testing.T) {
	t.Run("should extract number", func(t *testing.T) {
		const strToExtract = "hello 1234"
		m := newMatcher(`(\d+)`)
		extractedText := m.Extract(strToExtract)
		assert.Equal(t, extractedText, "1234")
	})

	t.Run("should extract substring", func(t *testing.T) {
		const strToExtract = "hello Denis!"
		m := newMatcher(`hello (\w+)!`)
		extractedText := m.Extract(strToExtract)
		assert.Equal(t, extractedText, "Denis")
	})

	t.Run("should extract speed", func(t *testing.T) {
		const strToMatch = "1.05M 100%    2.81MB/s    0:00:00 (xfr#1, ir-chk=3984/4068)"
		m := newMatcher(`(\d+\.\d+.{2}/s)`)
		extractedText := m.Extract(strToMatch)
		assert.Equal(t, extractedText, "2.81MB/s")
	})
}
