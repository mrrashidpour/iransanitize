package date

import (
	"testing"
)

func TestSanitizeDate(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		// تاریخ‌های شمسی معتبر
		{"شمسی کامل", "1374-09-02", "1374-09-02"},
		{"شمسی با اسلش", "1374/09/02", "1374-09-02"},
		{"شمسی بدون صفر", "1374-9-2", "1374-09-02"},
		{"شمسی با اعداد فارسی", "۱۳۷۴-۰۹-۰۲", "1374-09-02"},
		{"شمسی فرمت معکوس", "02/09/1374", "1374-09-02"},

		// تاریخ‌های میلادی معتبر
		{"میلادی کامل", "1994-11-23", "1994-11-23"},
		{"میلادی با اسلش", "1994/11/23", "1994-11-23"},
		{"میلادی بدون صفر", "1994-11-3", "1994-11-03"},
		{"میلادی فرمت معکوس", "23/11/1994", "1994-11-23"},

		// تاریخ‌های نامعتبر
		{"ماه نامعتبر میلادی", "1994-13-23", ""},
		{"روز نامعتبر شمسی", "1374-09-32", ""},
		{"روز نامعتبر میلادی", "1994-11-32", ""},
		{"فرمت نامعتبر", "13740902", ""},
		{"خالی", "", ""},
		{"رشته تصادفی", "abc", ""},
		{"سال کمتر از 1300", "1200-01-01", ""},
		{"سال بیشتر از 2100", "2200-01-01", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SanitizeDate(tt.input)
			if result != tt.expected {
				t.Errorf("input: %q, expected: %q, got: %q", tt.input, tt.expected, result)
			}
		})
	}
}

func TestIsValidDate(t *testing.T) {
	// تست‌های معتبر
	validDates := []string{
		"1374-09-02",
		"1994-11-23",
		"1400-12-29",
		"2024-02-28",
	}

	for _, date := range validDates {
		if !IsValidDate(date) {
			t.Errorf("expected valid: %s", date)
		}
	}

	// تست‌های نامعتبر
	invalidDates := []string{
		"1994-13-23",
		"1374-09-32",
		"1994-11-32",
		"1400-12-30",
		"2024-02-30",
	}

	for _, date := range invalidDates {
		if IsValidDate(date) {
			t.Errorf("expected invalid: %s", date)
		}
	}
}

func TestEdgeCases(t *testing.T) {
	// سال کبیسه شمسی
	if !IsValidDate("1399-12-30") {
		t.Error("1399-12-30 should be valid (leap year)")
	}

	// سال غیر کبیسه شمسی
	if IsValidDate("1400-12-30") {
		t.Error("1400-12-30 should be invalid (not leap year)")
	}

	// سال کبیسه میلادی
	if !IsValidDate("2020-02-29") {
		t.Error("2020-02-29 should be valid (leap year)")
	}

	// سال غیر کبیسه میلادی
	if IsValidDate("2023-02-29") {
		t.Error("2023-02-29 should be invalid (not leap year)")
	}
}

func BenchmarkSanitizeDate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SanitizeDate("1374-09-02")
	}
}
