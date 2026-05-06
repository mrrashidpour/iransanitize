package text

import (
	"testing"
)

func TestSanitize(t *testing.T) {
	tests := []struct {
		name         string
		input        interface{}
		keepNewlines bool
		expected     string
	}{
		{
			name:         "empty string",
			input:        "",
			keepNewlines: false,
			expected:     "",
		},
		{
			name:         "nil input",
			input:        nil,
			keepNewlines: false,
			expected:     "",
		},
		{
			name:         "simple text",
			input:        "Hello World",
			keepNewlines: false,
			expected:     "Hello World",
		},
		{
			name:         "text with spaces",
			input:        "  Hello   World  ",
			keepNewlines: false,
			expected:     "Hello World",
		},
		{
			name:         "text with tabs and newlines - collapse",
			input:        "Hello\t\tWorld\n\nTest",
			keepNewlines: false,
			expected:     "Hello World Test",
		},
		{
			name:         "text with tabs and newlines - preserve",
			input:        "Hello\t\tWorld\n\nTest",
			keepNewlines: true,
			expected:     "Hello  World\n\nTest",
		},
		{
			name:         "basic HTML tags",
			input:        "<p>Hello <strong>World</strong></p>",
			keepNewlines: false,
			expected:     "Hello World",
		},
		{
			name:         "script tag removal",
			input:        "Hello <script>alert('xss')</script> World",
			keepNewlines: false,
			expected:     "Hello World", // Script content removed, single space
		},
		{
			name:         "style tag removal",
			input:        "Hello <style>.test{color:red}</style> World",
			keepNewlines: false,
			expected:     "Hello World",
		},
		{
			name:         "nested HTML tags",
			input:        "<div><p>Line 1</p><p>Line 2</p></div>",
			keepNewlines: false,
			expected:     "Line 1Line 2",
		},
		{
			name:         "unclosed HTML tag - less than sign escaped",
			input:        "Hello <strong World",
			keepNewlines: false,
			expected:     "Hello &lt;strong World",
		},
		{
			name:         "percent encoded characters - only space preserved",
			input:        "Hello%20World%21%3F",
			keepNewlines: false,
			expected:     "Hello World",
		},
		{
			name:         "multiple percent encoded removed",
			input:        "%48%65%6C%6C%6F",
			keepNewlines: false,
			expected:     "",
		},
		{
			name:         "mixed percent and text",
			input:        "Test%20%41%42%20Text",
			keepNewlines: false,
			expected:     "Test Text", // Single space between Test and Text
		},
		{
			name:         "valid UTF-8 with special characters",
			input:        "Hello 世界 🌍",
			keepNewlines: false,
			expected:     "Hello 世界 🌍",
		},
		{
			name:         "map input returns empty",
			input:        map[string]interface{}{"key": "value"},
			keepNewlines: false,
			expected:     "",
		},
		{
			name:         "slice input returns empty",
			input:        []interface{}{"a", "b", "c"},
			keepNewlines: false,
			expected:     "",
		},
		{
			name:         "int input returns empty",
			input:        123,
			keepNewlines: false,
			expected:     "",
		},
		{
			name:         "HTML with line breaks preserved",
			input:        "<p>Line 1</p>\n<p>Line 2</p>\n<p>Line 3</p>",
			keepNewlines: true,
			expected:     "Line 1 Line 2 Line 3",
		},
		{
			name:         "HTML with line breaks collapsed",
			input:        "<p>Line 1</p>\n<p>Line 2</p>\n<p>Line 3</p>",
			keepNewlines: false,
			expected:     "Line 1 Line 2 Line 3",
		},
		{
			name:         "complex HTML with attributes",
			input:        `<a href="http://example.com" class="link">Click here</a> <img src="image.jpg" alt="image">`,
			keepNewlines: false,
			expected:     "Click here", // img tag has no text content
		},
		{
			name:         "single unclosed angle bracket",
			input:        "Hello < World",
			keepNewlines: false,
			expected:     "Hello &lt; World",
		},
		{
			name:         "multiple angle brackets",
			input:        "a < b > c",
			keepNewlines: false,
			expected:     "a c",
		},
		{
			name:         "preserve newlines with multiple line breaks",
			input:        "Line1\n\nLine2\n\n\nLine3",
			keepNewlines: true,
			expected:     "Line1\n\nLine2\n\n\nLine3",
		},
		{
			name:         "carriage return handling in collapse mode",
			input:        "Hello\r\nWorld\rTest",
			keepNewlines: false,
			expected:     "Hello World Test",
		},
		{
			name:         "carriage return preserve",
			input:        "Hello\r\nWorld\rTest",
			keepNewlines: true,
			expected:     "Hello\r\nWorld\rTest",
		},
		{
			name:         "percent with incomplete encoding preserved",
			input:        "Hello%2",
			keepNewlines: false,
			expected:     "Hello%2",
		},
		{
			name:         "email address preserved",
			input:        "user@example.com",
			keepNewlines: false,
			expected:     "user@example.com",
		},
		{
			name:         "URL preserved",
			input:        "https://example.com/path?q=value",
			keepNewlines: false,
			expected:     "https://example.com/path?q=value",
		},
		{
			name:         "HTML comment removed",
			input:        "Hello <!-- comment --> World",
			keepNewlines: false,
			expected:     "Hello World",
		},
		{
			name:         "self-closing HTML tag",
			input:        "Hello <br/> World",
			keepNewlines: false,
			expected:     "Hello World",
		},
		{
			name:         "HTML with attributes and quotes",
			input:        "<div class=\"test\" id='main'>Content</div>",
			keepNewlines: false,
			expected:     "Content",
		},
		{
			name:         "mix of valid and invalid HTML",
			input:        "<p>Valid</p> <invalid <a href='#'>Link</a>",
			keepNewlines: false,
			expected:     "Valid Link",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SanitizeText(tt.input, tt.keepNewlines)
			if result != tt.expected {
				t.Errorf("SanitizeText() = %q, want %q , name %q", result, tt.expected, tt.name)
			}
		})
	}
}

// TestSanitizePercentEncoding tests percent encoding removal
func TestSanitizePercentEncoding(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "single percent encoding",
			input:    "Hello%20World",
			expected: "Hello World",
		},
		{
			name:     "multiple percent encodings",
			input:    "Hello%20%57%6F%72%6C%64",
			expected: "Hello",
		},
		{
			name:     "lowercase percent",
			input:    "Hello%20world%21",
			expected: "Hello world",
		},
		{
			name:     "mixed case percent",
			input:    "Hello%20%57%6f%72%6c%64",
			expected: "Hello",
		},
		{
			name:     "invalid percent encoding preserved",
			input:    "Hello%2GWorld",
			expected: "Hello%2GWorld",
		},
		{
			name:     "percent without digits preserved",
			input:    "Hello% World",
			expected: "Hello% World",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SanitizeText(tt.input, false)
			if result != tt.expected {
				t.Errorf("SanitizeText() = %q, want %q ,name %q", result, tt.expected, tt.name)
			}
		})
	}
}

// TestSanitizeXSS attempts to test XSS prevention
func TestSanitizeXSS(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "javascript protocol",
			input:    "<a href=\"javascript:alert('xss')\">Click</a>",
			expected: "Click",
		},
		{
			name:     "onerror attribute",
			input:    "<img src=x onerror=alert('xss')>",
			expected: "",
		},
		{
			name:     "onload attribute",
			input:    "<body onload=alert('xss')>",
			expected: "",
		},
		{
			name:     "script tag with encoded content",
			input:    "<script>alert('xss')</script>",
			expected: "",
		},
		{
			name:     "nested script tags",
			input:    "<scr<script>ipt>alert('xss')</scr</script>ipt>",
			expected: "",
		},
		{
			name:     "event handler",
			input:    "<div onclick=\"alert('xss')\">Click</div>",
			expected: "Click",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SanitizeText(tt.input, false)
			if result != tt.expected {
				t.Errorf("SanitizeText() = %q, want %q", result, tt.expected)
			}
		})
	}
}

// TestSanitizeInvalidUTF8 tests invalid UTF-8 handling
func TestSanitizeInvalidUTF8(t *testing.T) {
	// Create invalid UTF-8 string
	invalidUTF8 := string([]byte{0x48, 0x65, 0x6C, 0x6C, 0x6F, 0xFF, 0x20, 0x57, 0x6F, 0x72, 0x6C, 0x64})

	result := SanitizeText(invalidUTF8, false)
	if result != "" {
		t.Errorf("Expected empty string for invalid UTF-8, got %q", result)
	}
}

// TestSanitizeWithNewlines tests newline preservation specifically
func TestSanitizeWithNewlines(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "single newline",
			input:    "Hello\nWorld",
			expected: "Hello\nWorld",
		},
		{
			name:     "multiple newlines",
			input:    "Hello\n\n\nWorld",
			expected: "Hello\n\n\nWorld",
		},
		{
			name:     "newlines with spaces",
			input:    "Hello \n World",
			expected: "Hello \n World",
		},
		{
			name:     "newlines with tabs",
			input:    "Hello\t\n\tWorld",
			expected: "Hello \n World", // Tabs become spaces
		},
		{
			name:     "Windows line endings",
			input:    "Hello\r\nWorld",
			expected: "Hello\r\nWorld",
		},
		{
			name:     "mixed line endings",
			input:    "Hello\r\nWorld\nTest\rLine",
			expected: "Hello\r\nWorld\nTest\rLine",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SanitizeText(tt.input, true)
			if result != tt.expected {
				t.Errorf("SanitizeText() = %q, want %q", result, tt.expected)
			}
		})
	}
}

// Benchmark tests
func BenchmarkSanitizeText(b *testing.B) {
	input := "<p>Hello <strong>World</strong> this is a <a href='#'>test</a> string with %20percent%20encoding</p>"

	b.Run("without newlines", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			SanitizeText(input, false)
		}
	})

	b.Run("with newlines", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			SanitizeText(input, true)
		}
	})

	b.Run("simple text", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			SanitizeText("Hello World", false)
		}
	})
}
