package utils

import (
	"slices"
	"strings"
	"unicode"
)

// func main() {

// 	log.Println(ReverseStringNew("hello world"))
// 	log.Println(ReverseString("hello world"))
// 	log.Println(ToUpper("hello world"))
// }

// ReverseString 反转字符串，需要在 1.21 以上版本生效
func ReverseStringNew(s string) string {
	r := []rune(s)
	slices.Reverse(r)
	return string(r)
}

// ReverseString 反转字符串
func ReverseString(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func ToUpperNew(s string) string {
	return strings.ToUpper(s)
}

// ToUpper uppercases all the runes in its argument string.
func ToUpper(s string) string {
	r := []rune(s)
	for i := range r {
		r[i] = unicode.ToUpper(r[i])
	}
	return string(r)
}
