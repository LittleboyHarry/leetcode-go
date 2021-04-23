package main

import (
	"math"
	"unicode"
)

type state struct {
	skippingPrefix bool
	negative       bool
	num            int32
}

func newState() state {
	return state{
		true,
		false,
		0,
	}
}

func (s *state) next(ch rune) bool {
	if s.skippingPrefix {
		if ch == ' ' {
			return true
		}
		s.skippingPrefix = false
		if ch == '+' {
			return true
		}
		if ch == '-' {
			s.negative = true
			return true
		}
		if unicode.IsDigit(ch) {
			s.next(ch)
			return true
		}
		return false
	} else {
		if unicode.IsDigit(ch) {
			old := s.num
			s.num *= 10
			if s.num/10 != old {
				if s.negative {
					s.num = math.MinInt32
				} else {
					s.num = math.MaxInt32
				}
				return false
			}
			if s.negative {
				s.num -= ch - '0'
				if s.num > old {
					s.num = math.MinInt32
					return false
				}
			} else {
				s.num += ch - '0'
				if s.num < old {
					s.num = math.MaxInt32
					return false
				}
			}
			return true
		}
		return false
	}
}

func myAtoi(s string) int {
	st := newState()
	for _, ch := range s {
		if !st.next(ch) {
			break
		}
	}
	return int(st.num)
}
