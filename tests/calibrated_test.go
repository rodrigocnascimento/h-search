package tests

import (
	"os"
	"strings"
	"testing"

	"github.com/rodrigocnascimento/h-search/detector"
)

func TestCalibration(t *testing.T) {
	d := detector.New(detector.WithThreshold(50))

	benign := loadDomains("data/benign.txt")
	malicious := loadDomains("data/malicious.txt")

	var tp, tn, fp, fn int

	// malicious → esperado true
	for _, domain := range malicious {
		res := d.Analyze(domain)
		if res.IsSuspicious {
			tp++
		} else {
			fn++
		}
	}

	// benign → esperado false
	for _, domain := range benign {
		res := d.Analyze(domain)
		if res.IsSuspicious {
			fp++
		} else {
			tn++
		}
	}

	t.Logf("TP=%d FN=%d FP=%d TN=%d", tp, fn, fp, tn)

	precision := float64(tp) / float64(tp+fp)
	recall := float64(tp) / float64(tp+fn)

	t.Logf("Precision=%.2f Recall=%.2f", precision, recall)
}

func loadDomains(path string) []string {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(data), "\n")

	var domains []string
	for _, l := range lines {
		l = strings.TrimSpace(l)
		if l != "" {
			domains = append(domains, l)
		}
	}
	return domains
}
