package str

import (
	"strings"
)

func CaseEqual(a, b string) bool {
	return len(a) == len(b) && strings.ToLower(a) == strings.ToLower(b)
}

func CaseHasPrefix(s, prefix string) (match bool, tail string) {
	l := len(prefix)
	if len(s) < l {
		return false, s
	}
	head, tail := s[:l], s[l:]
	return strings.ToLower(head) == strings.ToLower(prefix), tail
}

func CaseHasSuffix(s, suffix string) (match bool, head string) {
	l := len(suffix)
	if len(s) < l {
		return false, s
	}
	head, tail := s[:l], s[l:]
	return strings.ToLower(tail) == strings.ToLower(suffix), head
}

func Simp(complex string) (simple string) {
	clen := len(complex)
	simp := make([]byte, clen)
	slen := 0
	for cidx := 0; cidx < clen; cidx++ {
		codeunit := complex[cidx]
		switch {
		case codeunit >= 'A' && codeunit <= 'Z':
			codeunit += 'a' - 'A'
			fallthrough
		case codeunit >= 'a' && codeunit <= 'z':
			fallthrough
		case codeunit >= '0' && codeunit <= '9':
			fallthrough
		case codeunit == '-' || codeunit == '/' || codeunit == '_':
			simp[slen] = codeunit
			slen++
		}
	}
	return string(simp[:slen])
}
