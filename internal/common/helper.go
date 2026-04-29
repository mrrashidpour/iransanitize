package common

import (
	"strings"
	"unicode"
)

// ConvertPersianArabicToEnglish اعداد فارسی و عربی را به انگلیسی تبدیل می‌کند
func ConvertPersianArabicToEnglish(s string) string {
	persian := []rune{'۰', '۱', '۲', '۳', '۴', '۵', '۶', '۷', '۸', '۹'}
	arabic := []rune{'٠', '١', '٢', '٣', '٤', '٥', '٦', '٧', '٨', '٩'}

	result := []rune(s)
	for i, ch := range result {
		for j, p := range persian {
			if ch == p {
				result[i] = rune('0' + j)
				break
			}
		}
		for j, a := range arabic {
			if ch == a {
				result[i] = rune('0' + j)
				break
			}
		}
	}
	return string(result)
}

// RemoveSpecialChars حذف کاراکترهای خاص (فاصله، خط تیره، پرانتز، ...)
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
		"،", "",
	)
	return replacer.Replace(s)
}

// IsAllDigits بررسی می‌کند که همه کاراکترها رقم باشند
func IsAllDigits(s string) bool {
	for _, ch := range s {
		if !unicode.IsDigit(ch) {
			return false
		}
	}
	return true
}

// RemoveNonDigits حذف کاراکترهای غیرعددی
func RemoveNonDigits(s string) string {
	result := make([]rune, 0, len(s))
	for _, ch := range s {
		if unicode.IsDigit(ch) {
			result = append(result, ch)
		}
	}
	return string(result)
}
