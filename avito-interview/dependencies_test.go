package avito_interview

import "testing"

func TestPrintDeps(t *testing.T) {

	tests := []struct {
		name string
		deps Dependencies
	}{
		{
			name: "gsd",
			deps: Dependencies{
				"tensorflow": {"nvcc", "gpu", "linux"},
				"pyscc":      {"linux"},
				"linux":      {"core"},
				"mylib":      {"tensorflow"},
				"mylib2":     {"requests"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			PrintDeps(tt.deps)
		})
	}
}
