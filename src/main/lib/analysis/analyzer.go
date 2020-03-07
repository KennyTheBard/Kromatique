package analysis

import (
	"image"
)

// Analyzer is a function that encapsulates a behaviour applied on
// a point of an image and saves the result or results in a given map
type Analyze func(image.Image, int, int, map[string]interface{})

// AnalyzerRunner encapsulates the logic and data needed to run
// analyzer functions on an image and retrieves the obtained data
type AnalyzerRunner struct {
	analyzers []Analyze
}

// Run applies each analyzer function on each pixel of the image
// and returns a map of values obtained by each analyzer
func (r *AnalyzerRunner) Run(img image.Image) map[string]interface{} {
	dict := make(map[string]interface{})
	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			for _, analyzer := range r.analyzers {
				analyzer(img, x, y, dict)
			}
		}
	}

	return dict
}

// NewAnalyzerRunner creates a new AnalyzerRunner with given analyzers
func NewAnalyzerRunner(analyzers []Analyze) *AnalyzerRunner {
	runner := new(AnalyzerRunner)
	runner.analyzers = analyzers

	return runner
}
