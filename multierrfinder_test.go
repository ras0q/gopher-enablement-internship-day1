package multierrfinder_test

import (
	"testing"

	"github.com/gostaticanalysis/testutil"
	"github.com/ras0q/multierrfinder"
	"golang.org/x/tools/go/analysis/analysistest"
)

// TestAnalyzer is a test for Analyzer.
func TestAnalyzer(t *testing.T) {
	testdata := testutil.WithModules(t, analysistest.TestData(), nil)
	t.Run("a", func(t *testing.T) {
		analysistest.Run(t, testdata, multierrfinder.Analyzer, "a")
	})

	t.Run("b", func(t *testing.T) {
		analysistest.Run(t, testdata, multierrfinder.Analyzer, "b")
	})

	t.Run("rename", func(t *testing.T) {
		analysistest.Run(t, testdata, multierrfinder.Analyzer, "rename")
	})
}
