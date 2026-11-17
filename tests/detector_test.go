package tests

import (
	"testing"

	"github.com/rodrigocnascimento/h-search/detector"
)

func TestDetector(t *testing.T) {
	d := detector.New()

	tests := []struct {
		domain        string
		expectSuspect bool
	}{
		{"xkqzrt123.xyz", true},
		{"alquvsuyqnanqremddxhvtyk.com", true},
		{"google.com", false},
		{"facebook.com", false},
		{"cade.com.br", false},
		{"ajsdklqweoi123.biz", true},
	}

	for _, tt := range tests {
		res := d.Analyze(tt.domain)

		t.Logf("domain=%s score=%.2f features=%+v", tt.domain, res.Score, res.Features)

		if res.IsSuspicious != tt.expectSuspect {
			t.Errorf(
				"domain=%s expected=%v got=%v score=%.2f",
				tt.domain,
				tt.expectSuspect,
				res.IsSuspicious,
				res.Score,
			)
		}
	}
}
