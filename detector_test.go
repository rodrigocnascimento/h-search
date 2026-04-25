package detector_test

import (
	"testing"
	"github.com/yourusername/dga-detector/detector"
)

func TestDetector(t *testing.T) {
	d := detector.New()

	domains := []string{
		"xkqzrt123.xyz",
		"google.com",
	}

	for _, domain := range domains {
		res := d.Analyze(domain)
		if domain == "google.com" && res.IsSuspicious {
			t.Errorf("false positive: %s", domain)
		}
	}
}
