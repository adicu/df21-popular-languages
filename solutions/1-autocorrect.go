package main

import "fmt"

// Source: https://github.com/agnivade/levenshtein/blob/master/levenshtein.go
func ComputeDistance(a, b string) int {
	if len(a) == 0 {
		return len(b)
	}

	if len(b) == 0 {
		return len(a)
	}

	if a == b {
		return 0
	}

	// We need to convert to []rune if the strings are non-ASCII.
	// This could be avoided by using utf8.RuneCountInString
	// and then doing some juggling with rune indices,
	// but leads to far more bounds checks. It is a reasonable trade-off.
	s1 := []rune(a)
	s2 := []rune(b)

	// swap to save some memory O(min(a,b)) instead of O(a)
	if len(s1) > len(s2) {
		s1, s2 = s2, s1
	}
	lenS1 := len(s1)
	lenS2 := len(s2)

	// init the row
	x := make([]uint16, lenS1+1)
	// we start from 1 because index 0 is already 0.
	for i := 1; i < len(x); i++ {
		x[i] = uint16(i)
	}

	// make a dummy bounds check to prevent the 2 bounds check down below.
	// The one inside the loop is particularly costly.
	_ = x[lenS1]
	// fill in the rest
	for i := 1; i <= lenS2; i++ {
		prev := uint16(i)
		for j := 1; j <= lenS1; j++ {
			current := x[j-1] // match
			if s2[i-1] != s1[j-1] {
				current = min(min(x[j-1]+1, prev+1), x[j]+1)
			}
			x[j-1] = prev
			prev = current
		}
		x[lenS1] = prev
	}
	return int(x[lenS1])
}

func min(a, b uint16) uint16 {
	if a < b {
		return a
	}
	return b
}

var dictionary = [...]string{
	"banana",
	"orange",
	"apple",
	"pear",
	"mango",
	"watermelon",
	"pineapple",
}

func Autocorrect(word string) string {
	shortestDistance := len(word)
	correctWord := word

	for _, dictionaryWord := range dictionary {
		currentDistance := ComputeDistance(dictionaryWord, word)
		if currentDistance < shortestDistance {
			shortestDistance = currentDistance
			correctWord = dictionaryWord
		}
	}

	return correctWord
}

func main() {
	fmt.Println(ComputeDistance("bar","bat"))
	fmt.Println(Autocorrect("babana"))
}
