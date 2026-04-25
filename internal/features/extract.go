package features

import (
	"unicode"
	"strings"
	"math"
)

type Features struct {
	Entropy           float64
	Length            int
	DigitRatio        float64
	VowelRatio        float64
	ConsonantClusters int
	UniqueCharRatio   float64
}

func Extract(domain string) Features {
	s := normalize(domain)
	length := float64(len(s))

	if length == 0 {
		return Features{}
	}

	var digits, vowels, unique int
	seen := make(map[rune]bool)

	vowelSet := "aeiou"
	consecutive := 0
	maxCluster := 0

	for _, r := range s {
		if !seen[r] {
			seen[r] = true
			unique++
		}

		if unicode.IsDigit(r) {
			digits++
		}

		if strings.ContainsRune(vowelSet, r) {
			vowels++
			consecutive = 0
		} else if unicode.IsLetter(r) {
			consecutive++
			if consecutive > maxCluster {
				maxCluster = consecutive
			}
		}
	}

	return Features{
		Entropy:           entropy(s),
		Length:            int(length),
		DigitRatio:        float64(digits) / length,
		VowelRatio:        float64(vowels) / length,
		ConsonantClusters: maxCluster,
		UniqueCharRatio:   float64(unique) / length,
	}
}

func entropy(s string) float64 {
	m := make(map[rune]float64)
	for _, r := range s {
		m[r]++
	}
	var e float64
	l := float64(len(s))
	for _, c := range m {
		p := c / l
		e -= p * math.Log2(p)
	}
	return e
}
