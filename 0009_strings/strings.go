package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	// Strings in go are enclosed within double quotes and utf-8 encoded
	s1 := "Hi!"
	fmt.Println(s1)
	// If you have double quotes within the string, either escape the
	// enclosed quotes or use backticks
	s2 := "\"Hello\""
	s3 := `"hello"?`
	fmt.Println(s2)
	fmt.Println(s3)
	// Note that strings within backticks are raw strings and the print
	// escape characters literally
	fmt.Println(`"hi" \n "hello"`)
	// Strings can be concatenated using the "+" operator. Note that
	// go generates a new string each time you concatenate and hence this
	// may not be efficient.
	s4 := "foo"
	s5 := "bar"
	s6 := s4 + "_" + s5
	fmt.Println(s6)

	// When you access the elements of a string by index you get the ascii
	// value since a string is basically an array of bytes
	s7 := "abcd"
	fmt.Println(s7[0]) // prints 97

	// Strings in golang are immutable and the following line is incorrect
	// s7[4] = "x"

	// %s is formate specifier for string & %q for quoted string
	s8 := "string"
	fmt.Printf("%s\n", s8)
	fmt.Printf("%q\n", s8)

	// go doesn't have a char datatype, it instead uses bytes and runes.
	// bytes (uint8) represents ascii characters
	// rune (int32) represents characters in utf-8 encoding
	// runes are enclosed in single quotes
	var r1, r2 = 'a', 'b'
	fmt.Printf("r1 is of type %T, byte value %d and char value %c\n", r1, r1, r1)
	fmt.Printf("r2 is of type %T, byte value %d and char value %c\n", r2, r2, r2)

	// Length of string is sum of length of runes and runes may be of size
	// 1 - 4 bytes depending on the utf-8 character it's representing
	unicodeStr := "a¥z"
	fmt.Println(len(unicodeStr))                    // len = 4 even though we see 3 chars
	fmt.Println(utf8.RuneCountInString(unicodeStr)) // prints 3

	// If we iterate over indices of string we endup printing the bytes and no
	// the actual string
	for i := 0; i < len(unicodeStr); i++ {
		fmt.Printf("%c", unicodeStr[i])
	}
	fmt.Println()

	// If we print rune by rune we get the actual string
	for i := 0; i < len(unicodeStr); {
		utf8Rune, size := utf8.DecodeRuneInString(unicodeStr[i:]) // rune, size := utf8.DecodeRuneInString(string[idx:])
		fmt.Printf("%c", utf8Rune)
		i += size
	}
	fmt.Println()

	// To decode the rune automatically we use range
	for _, utf8Rune := range unicodeStr {
		fmt.Printf("%c", utf8Rune)
	}
	fmt.Println()

	// Slicing a string is efficient since it uses same backing array
	// By default returns slice of bytes and not slice of runes
	fmt.Println(unicodeStr[0:2])
	// To obtain rune slice, first convert to rune slice and convert
	// the runes in slice to string. This however is not efficient since
	// converting between string and runes slice leads to new backing
	// array being created
	runeSlice := []rune(unicodeStr)
	fmt.Println(string(runeSlice[0:2]))

	///// Functions in strings package

	// strings.Contains
	sentence := "This is the main string"
	subStr := "main"
	fmt.Println(strings.Contains(sentence, subStr))
	subStr = "mains"
	fmt.Println(strings.Contains(sentence, subStr))

	// strings.ContainsAny return true if any unicode character is presnet
	fmt.Println(strings.ContainsAny(sentence, "oooooooooh")) // true: 'h' is in sentence

	// strings.ContainsRune checks if rune is in string
	fmt.Println(strings.ContainsRune(sentence, '¥'))   // false
	fmt.Println(strings.ContainsRune(unicodeStr, '¥')) // true

	// strints.Count returns num of occurances of substring.
	// if substring is empty returns 1 + len(string)
	fmt.Println(strings.Count("cheese", "e")) // 3
	fmt.Println(strings.Count("cheese", ""))  // 7

	// strings.ToUpper returns all chars in upper
	newS := "aBc"
	fmt.Println(strings.ToUpper(newS))
	// strings.ToLower returns all chars in lower
	fmt.Println(strings.ToLower(newS))

	// since strings are immutable, ToUpper & ToLower return new strings
	// and hence may be inefficient. For example Comparing strings in go
	// while excluding case can be done inefficiently as
	fmt.Println(strings.ToUpper("go") == strings.ToUpper("Go"))
	fmt.Println(strings.ToLower("go") == strings.ToLower("Go"))
	// We can do this efficiently using strings.EqualFold which does a
	// case insensitive matching. Unlike ToUpper & ToLower, instead of
	// creating 2 news slices and comparing them, strings.EqualFold
	// compares runes of the two strings and does a case insensitive
	// comparision if the runes do not match.
	fmt.Println(strings.EqualFold("go", "Go"))

	// strings.Repeat returns new string containing 'n' copies of the string
	repeater := "#"
	fmt.Println(strings.Repeat(repeater, 20))

	// strings.Replace repalces substring s1 with substring s2 'n' times
	newStr := "\\Users\\Dowloads\\file"
	fmt.Println(strings.Replace(newStr, "\\", "/", 2))
	// strings.ReplaceAll replaces all instances of substring s1 with substring s2
	fmt.Println(strings.ReplaceAll(newStr, "\\", "/"))

	// strings.Split slices string based on separater and returns the slices
	ip := "10.46.64.103"
	octets := strings.Split(ip, ".")
	fmt.Printf("octets if of type %T and value %#v\n", octets, octets)
	// if separater is empty Split function splits after each rune literal
	chars := strings.Split("Golang¥", "")
	fmt.Printf("chars are %#v\n", chars)

	// strings.Join joins a slice of strings by inserting separater between each
	// element of the slice
	ip = strings.Join(octets, ".")
	fmt.Println("IP is", ip)

	// strings.Fields splits a string by whitespaces & newlines
	newString := "Hello ?\n Who's there?"
	fields := strings.Fields(newString)
	fmt.Printf("fields is of type %T and value %#v\n", fields, fields)

	// strings.TrimSpace remove leading & trailing spaces
	spacedStr := "  There is a cow How do you cross?\t"
	fmt.Printf("%q\n", strings.TrimSpace(spacedStr))

	// strings.Trim removes other leading & trailing characters
	leadTrailStr := ".........aaaaaaaaaaaa........"
	fmt.Printf("%q\n", strings.Trim(leadTrailStr, "."))
}
