package main

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"
)

func TestSampleMeasurementFiles(t *testing.T) {
	// Get the list of input files in the samples directory.
	paths, err := filepath.Glob(filepath.Join("samples", "*.txt"))
	if err != nil {
		t.Fatal(err)
	}

	for _, path := range paths {
		_, filename := filepath.Split(path)
		testname := filename[:len(filename)-len(filepath.Ext(path))]

		// Each path turns into a test: the test name is the filename without the
		// extension.
		t.Run(testname, func(t *testing.T) {
			// >>> This is the actual code under test.
			result, err := process(path)
			if err != nil {
				t.Fatal("error processing:", err)
			}
			// <<<

			// Each input file is expected to have a "golden output" file, with the
			// same path except the .input extension is replaced by .golden
			goldenfile := filepath.Join("samples", testname+".out")
			want, err := os.ReadFile(goldenfile)
			if err != nil {
				t.Fatal("error reading golden file:", err)
			}

			if !bytes.Equal([]byte(result), want) {
				t.Errorf("\n==== got:\n%s\n==== want:\n%s\n", result, want)
			}
		})
	}
}
