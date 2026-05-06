package mobile

import (
	"regexp"
	"strings"

	"github.com/mrrashidpour/iransanitize/internal/common"
)

func Sanitize(mobile string) string {
	mobile = strings.TrimSpace(mobile)

	mobile = common.ConvertToEnglishDigits(mobile)
	mobile = common.RemoveSpecialChars(mobile)

	if strings.HasPrefix(mobile, ".") {
		mobile = "0" + mobile[1:]
	}

	mobile = removePrefixes(mobile)

	if !strings.HasPrefix(mobile, "0") {
		mobile = "0" + mobile
	}

	if !common.IsAllDigits(mobile) {
		return ""
	}

	if len(mobile) != 11 {
		return ""
	}

	pattern := regexp.MustCompile(`^(\+98|0)?9\d{9}$`)
	if !pattern.MatchString(mobile) {
		return ""
	}

	return mobile
}

// removePrefixes پیشوندهای مختلف را حذف می‌کند
func removePrefixes(mobile string) string {
	switch {
	case len(mobile) == 14 && strings.HasPrefix(mobile, "00989"):
		return mobile[4:]
	case len(mobile) == 13 && strings.HasPrefix(mobile, "0989"):
		return mobile[3:]
	case len(mobile) == 13 && strings.HasPrefix(mobile, "+989"):
		return mobile[3:]
	case len(mobile) == 12 && strings.HasPrefix(mobile, "989"):
		return mobile[2:]
	default:
		return mobile
	}
}
