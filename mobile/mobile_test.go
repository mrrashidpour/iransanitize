package mobile

import (
	"testing"
)

func TestSanitize(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"شماره استاندارد", "09034665213", "09034665213"},
		{"شماره استاندارد", "09123456789", "09123456789"},
		{"بدون صفر ابتدایی", "9123456789", "09123456789"},
		{"با فاصله", "0912 345 6789", "09123456789"},
		{"با خط تیره", "0912-345-6789", "09123456789"},
		{"با کد بین‌المللی", "+989123456789", "09123456789"},
		{"با کد 98", "989123456789", "09123456789"},
		{"با کد 0098", "00989123456789", "09123456789"},
		{"اعداد فارسی", "۰۹۱۲۳۴۵۶۷۸۹", "09123456789"},
		{"اعداد عربی", "٠٩١٢٣٤٥٦٧٨٩", "09123456789"},
		{"شماره نامعتبر", "1234567890", ""},
		{"شماره کوتاه", "09123", ""},
		{"شماره بلند", "091234567890", ""},
		{"شماره با حروف", "0912abc6789", ""},
		{"خالی", "", ""},
		{"فقط صفر", "0", ""},
		{"پیشوند نامعتبر", "09999999999", "09999999999"}, // پیشوند 099 معتبر است
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Sanitize(tt.input)
			if result != tt.expected {
				t.Errorf("expected %s, got %s", tt.expected, result)
			}
		})
	}
}

func TestConvertToInternational(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"09123456789", "+989123456789"},
		{"9123456789", "+989123456789"},
		{"invalid", ""},
	}

	for _, tt := range tests {
		result := ConvertToInternational(tt.input)
		if result != tt.expected {
			t.Errorf("for %s expected %s, got %s", tt.input, tt.expected, result)
		}
	}
}

func TestExtractOperator(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"09123456789", "همراه اول"},
		{"09301234567", "همراه اول"},
		{"09901234567", "ایرانسل"},
		{"09201234567", "رایتل"},
		{"invalid", ""},
	}

	for _, tt := range tests {
		result := ExtractOperator(tt.input)
		if result != tt.expected {
			t.Errorf("for %s expected %s, got %s", tt.input, tt.expected, result)
		}
	}
}

func TestMask(t *testing.T) {
	result := Mask("09123456789")
	expected := "0912***6789"
	if result != expected {
		t.Errorf("expected %s, got %s", expected, result)
	}
}

func TestCompare(t *testing.T) {
	if !Compare("09123456789", "9123456789") {
		t.Error("should be equal")
	}

	if Compare("09123456789", "09123456780") {
		t.Error("should not be equal")
	}
}

func TestFormatWithDash(t *testing.T) {
	result := FormatWithDash("09123456789")
	expected := "0912-345-6789"
	if result != expected {
		t.Errorf("expected %s, got %s", expected, result)
	}
}

func TestStrictMode(t *testing.T) {
	result := Sanitize("0912۳۴۵۶۷۸۹", StrictOptions)
	if result != "" {
		t.Error("strict mode should reject Persian numbers")
	}
}

func BenchmarkSanitize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sanitize("0912-345-6789")
	}
}
