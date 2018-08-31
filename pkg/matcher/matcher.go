package matcher

import (
	"regexp"
)

type Matcher struct {
	regExp *regexp.Regexp
}

func (m Matcher) Match(data string) bool {
	return m.regExp.Match([]byte(data))
}

func (m Matcher) Extract(data string) string {
	const submatchCount = 1
	matches := m.regExp.FindAllStringSubmatch(data, submatchCount)
	if len(matches) == 0 || len(matches[0]) < 2 {
		return ""
	}

	return matches[0][1]
}

func (m Matcher) ExtractAllStringSubmatch(data string, submatchCount int) [][]string {
	return m.regExp.FindAllStringSubmatch(data, submatchCount)
}

func New(regExpString string) *Matcher {
	return &Matcher{
		regExp: regexp.MustCompile(regExpString),
	}
}
