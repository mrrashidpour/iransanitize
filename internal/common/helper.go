package common

import (
	"strings"
	"unicode"
)

// ConvertPersianArabicToEnglish converts Persian/Arabic digits to English
func ConvertPersianArabicToEnglish(s string) string {
	persian := []rune{'۰', '۱', '۲', '۳', '۴', '۵', '۶', '۷', '۸', '۹'}
	arabic := []rune{'٠', '١', '٢', '٣', '٤', '٥', '٦', '٧', '٨', '٩'}

	result := make([]rune, 0, len(s))
	for _, ch := range s {
		converted := false
		for j, p := range persian {
			if ch == p {
				result = append(result, rune('0'+j))
				converted = true
				break
			}
		}
		if !converted {
			for j, a := range arabic {
				if ch == a {
					result = append(result, rune('0'+j))
					converted = true
					break
				}
			}
		}
		if !converted {
			result = append(result, ch)
		}
	}
	return string(result)
}

// RemoveSpecialChars removes special characters from string
func RemoveSpecialChars(s string) string {
	replacer := strings.NewReplacer(
		" ", "",
		"-", "",
		"+", "",
		"(", "",
		")", "",
		"_", "",
		".", "",
		"،", "",
		"؛", "",
	)
	return replacer.Replace(s)
}

// IsAllDigits checks if all characters are digits
func IsAllDigits(s string) bool {
	if s == "" {
		return false
	}
	for _, ch := range s {
		if !unicode.IsDigit(ch) {
			return false
		}
	}
	return true
}

// RemoveNonDigits removes all non-digit characters
func RemoveNonDigits(s string) string {
	result := make([]rune, 0, len(s))
	for _, ch := range s {
		if unicode.IsDigit(ch) {
			result = append(result, ch)
		}
	}
	return string(result)
}

// IsAllEnglishDigits بررسی می‌کند که همه کاراکترها ارقام انگلیسی باشند
func IsAllEnglishDigits(s string) bool {
	if s == "" {
		return false
	}
	for _, ch := range s {
		if ch < '0' || ch > '9' {
			return false
		}
	}
	return true
}

// HasPersianOrArabicDigits بررسی می‌کند که آیا رشته دارای اعداد فارسی یا عربی است
func HasPersianOrArabicDigits(s string) bool {
	persian := []rune{'۰', '۱', '۲', '۳', '۴', '۵', '۶', '۷', '۸', '۹'}
	arabic := []rune{'٠', '١', '٢', '٣', '٤', '٥', '٦', '٧', '٨', '٩'}

	for _, ch := range s {
		for _, p := range persian {
			if ch == p {
				return true
			}
		}
		for _, a := range arabic {
			if ch == a {
				return true
			}
		}
	}
	return false
}
