package problems

import "testing"

func TestIsValidParentheses(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"()", true},
		{"()[]{}", true},
		{"(]", false},
		{"([)]", false},
		{"{[]}", true},
		{"", true},
	}

	for _, tt := range tests {
		if got := IsValidParentheses(tt.input); got != tt.want {
			t.Fatalf("input %q: expected %v, got %v", tt.input, tt.want, got)
		}
	}
}
