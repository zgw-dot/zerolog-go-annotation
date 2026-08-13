package zerolog

import "testing"

func TestParseLevelWarningAlias(t *testing.T) {
	for _, input := range []string{"warning", "WARNING"} {
		level, err := ParseLevel(input)
		if err != nil || level != WarnLevel {
			t.Fatalf("ParseLevel(%q) = %v, %v; want WarnLevel", input, level, err)
		}
	}
}
