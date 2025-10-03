package booty

import (
	"bytes"
	"strings"
	"testing"
)

func TestCalculateShares(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		total int
		crew  int
		share int
		rest  int
	}{
		"zero total, one crew":          {total: 0, crew: 1, share: 0, rest: 0},
		"odd with captain double share": {total: 1, crew: 1, share: 0, rest: 1},
		"exact division with captain":   {total: 2, crew: 1, share: 1, rest: 0},
		"remainder with captain":        {total: 3, crew: 1, share: 1, rest: 1},
		"two crew plus captain":         {total: 10, crew: 2, share: 3, rest: 1},
		"three crew plus captain":       {total: 10, crew: 3, share: 2, rest: 2},
		"larger crew":                   {total: 100, crew: 10, share: 9, rest: 1},
		"large numbers":                 {total: 1_000_000, crew: 1, share: 500_000, rest: 0},
		"zero total many crew":          {total: 0, crew: 5, share: 0, rest: 0},
		"prime totals":                  {total: 13, crew: 7, share: 1, rest: 5},
		"small total larger crew":       {total: 1, crew: 2, share: 0, rest: 1},
	}

	for name, tc := range tests {
		name, tc := name, tc // capture for parallel subtests
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			gotShare, gotRest := CalculateShares(tc.total, tc.crew)
			if gotShare != tc.share || gotRest != tc.rest {
				t.Errorf("CalculateShares(%d,%d) = (%d,%d), want (%d,%d)", tc.total, tc.crew, gotShare, gotRest, tc.share, tc.rest)
			}
		})
	}
}

func TestAskIntValid(t *testing.T) {
	t.Parallel()
	in := strings.NewReader("42\n")
	var out bytes.Buffer
	v, err := AskInt(in, &out, "How many pieces?")
	if err != nil {
		t.Fatalf("AskInt returned error: %v", err)
	}
	if v != 42 {
		t.Fatalf("AskInt returned %d, want 42", v)
	}
	if got := out.String(); got != "How many pieces? " {
		t.Fatalf("prompt output = %q, want %q", got, "How many pieces? ")
	}
}

func TestAskIntInvalid(t *testing.T) {
	t.Parallel()
	cases := map[string]string{
		"not a number":  "abc\n",
		"below minimum": "0\n",
		"negative":      "-3\n",
		"empty":         "\n",
	}
	for name, input := range cases {
		name, input := name, input // capture for parallel subtests
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			in := strings.NewReader(input)
			var out bytes.Buffer
			_, err := AskInt(in, &out, "Enter a number:")
			if err == nil {
				t.Fatalf("AskInt(%q) returned nil error, want non-nil", input)
			}
		})
	}
}
