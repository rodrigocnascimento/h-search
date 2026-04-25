package scoring

import "github.com/rodrigocnascimento/h-search/internal/features"

func Calculate(domain string, f features.Features) float64 {
	score := 0.0

	if f.Entropy > 3.8 {
		score += (f.Entropy - 3.8) * 25
	}

	if f.Length > 20 {
		score += float64(f.Length-20) * 1.5
	}

	score += f.DigitRatio * 30

	if f.VowelRatio < 0.25 {
		score += (0.25 - f.VowelRatio) * 40
	}

	if f.ConsonantClusters >= 4 {
		score += float64(f.ConsonantClusters) * 5
	}

	if f.UniqueCharRatio > 0.7 {
		score += (f.UniqueCharRatio - 0.7) * 50
	}

	return score
}
