package detector

type Result struct {
	Domain       string
	Score        float64
	IsSuspicious bool
	Features     Features
}

type Features struct {
	Entropy           float64
	Length            int
	DigitRatio        float64
	VowelRatio        float64
	ConsonantClusters int
	UniqueCharRatio   float64
}
