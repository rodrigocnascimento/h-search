package detector

import "github.com/rodrigocnascimento/h-search/internal/features"

type Result struct {
	Domain       string
	Score        float64
	IsSuspicious bool
	Features     features.Features
}
