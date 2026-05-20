package iransanitize

import (
	"github.com/mrrashidpour/iransanitize/date"
	"github.com/mrrashidpour/iransanitize/mobile"
	"github.com/mrrashidpour/iransanitize/text"
)

// ============ Mobile Functions ============

func SanitizeMobile(mobileStr string) string {
	return mobile.Sanitize(mobileStr)
}

func MaskMobile(mobileStr string) string {
	return mobile.Mask(mobileStr)
}

func CompareMobile(mobile1, mobile2 string) bool {
	return mobile.Compare(mobile1, mobile2)
}

// ============ Date Functions ============

// SanitizeDate تاریخ را به فرمت استاندارد YYYY-MM-DD تبدیل می‌کند
func SanitizeDate(dateStr string, options ...date.Option) string {

	var option date.Option

	if len(options) == 1 {
		option = options[0]
	}

	return date.SanitizeDate(dateStr, &option)
}

// ============ Text Functions ============

// SanitizeText پاکسازی متن
func SanitizeText(dateStr string, keepNewlines bool) string {
	return text.Sanitize(dateStr, keepNewlines)
}
