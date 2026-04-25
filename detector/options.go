package detector

type Option func(*Detector)

func WithThreshold(t float64) Option {
	return func(d *Detector) {
		d.threshold = t
	}
}
