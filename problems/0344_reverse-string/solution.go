package main

func reverseString(s []byte) {
	leftP, rightP := 0, len(s)-1
	for leftP < rightP {
		s[leftP], s[rightP] = s[rightP], s[leftP]

		leftP++
		rightP--
	}
}
