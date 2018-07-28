package main

import (
	"fmt"
	"strings"
)

func main() {
	// Contains(s, substr string) bool
	fmt.Println(strings.Contains("test","es"))
	// => true

	// func Count(s, sep string) int
	fmt.Println(strings.Count("test","t"))
	// => 2

	// func HasPrefix(s, prefix string) bool
	fmt.Println(strings.HasPrefix("test","te"))
	// => true

	// func HasSuffix(s, suffix string) bool
	fmt.Println(strings.HasSuffix("test","st"))
	// => true

	// func Index(s, sep string) int
	fmt.Println(strings.Index("test","e"))
	// => 1

	// func Join(a []string, sep string) string
	fmt.Println(strings.Join([]string{"a","b"}, "-"))
	// => "a-b"

	// func Repeat(s string, count int) string
	fmt.Println(strings.Repeat("a", 5))
	// => "aaaaa"

	// func Replace(s, old, new string, n int) string
	fmt.Println(strings.Replace("aaaa","a","b",2))
	// => "bbaa"

	// func Split(s, sep string) []string
	fmt.Println(strings.Split("a-b-c-d-e", "-"))
	// => []string{"a","b","c","d","e"}

	// func ToLower(s string) string
	fmt.Println(strings.ToLower("TEST"))
	// => "test"

	// func ToUpper(s string) string
	fmt.Println(strings.ToUpper("test"))
	// => "TEST"

	arr := []byte("test")
	fmt.Println(arr)
	
	str := string([]byte{'t','e','s','t'})
	fmt.Println(str)
}