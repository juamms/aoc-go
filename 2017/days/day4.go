package days

import (
	"sort"
	"strings"

	"github.com/juamms/aoc-go/2017/utils"
)

type Day4 struct{}

func (day Day4) Part1() interface{} {
	validPasswords := 0

	utils.ScanInputFile(4, func(password string) {
		cache := make(map[string]bool)
		words := strings.Split(password, " ")

		for _, word := range words {
			if _, ok := cache[word]; ok {
				return
			}

			cache[word] = true
		}

		validPasswords++
	})

	return validPasswords
}

// sorting functions from: https://siongui.github.io/2017/05/07/go-sort-string-slice-of-rune/
func stringToRuneSlice(s string) []rune {
	var r []rune
	for _, runeValue := range s {
		r = append(r, runeValue)
	}
	return r
}

func sortStringByCharacter(s string) string {
	r := stringToRuneSlice(s)
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})
	return string(r)
}

func (day Day4) Part2() interface{} {
	validPasswords := 0

	utils.ScanInputFile(4, func(password string) {
		cache := make(map[string]bool)
		words := strings.Split(password, " ")

		for _, word := range words {
			token := sortStringByCharacter(word)

			if _, ok := cache[token]; ok {
				return
			}

			cache[token] = true
		}

		validPasswords++
	})

	return validPasswords
}
