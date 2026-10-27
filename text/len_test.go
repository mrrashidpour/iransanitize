package text

import (
	"testing"
)

func TestCountChars(t *testing.T) {
	tests := []struct {
		name         string
		input        string
		keepNewlines bool
		expected     int
	}{
		{
			name:     "empty string",
			input:    "",
			expected: 0,
		},
		{
			name:     "simple text",
			input:    "Hello World",
			expected: 11,
		},
		{
			name:     "text with spaces",
			input:    "این برای تست میباشد",
			expected: 16,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CountChars(tt.input)
			if result != tt.expected {
				t.Errorf("CountChars() = %d, want %d , name %s", result, tt.expected, tt.name)
			}
		})
	}
}
