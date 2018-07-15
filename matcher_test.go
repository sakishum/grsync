package grsync

import (
	"testing"
)

func TestMatcherMatch(t *testing.T) {
	t.Run("should match number", func(t *testing.T) {
		const strToMatch = "hello 1232"
		m := newMatcher(`(\d+)`)
		if !m.Match(strToMatch) {
			t.Errorf("should match string: %s, but don't", strToMatch)
		}
	})

	t.Run("should match substring", func(t *testing.T) {
		const strToMatch = "hello Denis!"
		m := newMatcher(`hello (\w+)!`)
		if !m.Match(strToMatch) {
			t.Errorf("should match string: %s, but don't", strToMatch)
		}
	})
}

func TestMatcherExtract(t *testing.T) {
	t.Run("should extract number", func(t *testing.T) {
		const strToExtract = "hello 1234"
		m := newMatcher(`(\d+)`)
		extractedText := m.Extract(strToExtract)
		if extractedText != "1234" {
			t.Errorf("expected `1234`, but found %s", extractedText)
		}
	})

	t.Run("should extract substring", func(t *testing.T) {
		const strToExtract = "hello Denis!"
		m := newMatcher(`hello (\w+)!`)
		extractedText := m.Extract(strToExtract)
		if extractedText != "Denis" {
			t.Errorf("expected `Denis`, but found %s", extractedText)
		}
	})
}
