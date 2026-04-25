package detector

import (
	"github.com/yourusername/dga-detector/internal/features"
	"github.com/yourusername/dga-detector/internal/scoring"
)

type Detector struct {
	threshold float64
}

func New(opts ...Option) *Detector {
	d := &Detector{
		threshold: 60,
	}
	for _, opt := range opts {
		opt(d)
	}
	return d
}

func (d *Detector) Analyze(domain string) Result {
	f := features.Extract(domain)
	score := scoring.Calculate(domain, f)

	return Result{
		Domain:       domain,
		Score:        score,
		IsSuspicious: score >= d.threshold,
		Features:     f,
	}
}
