package mobile

import (
	"github.com/mrrashidpour/iransanitize/internal/common"
	"regexp"
	"strings"
)

// Sanitize شماره موبایل را پاکسازی و استاندارد می‌کند
// خروجی: شماره موبایل ۱۱ رقمی با فرمت 09XXXXXXXXX
// در صورت نامعتبر بودن: رشته خالی برگشت داده می‌شود
func Sanitize(mobile string, opts ...Option) string {
	if mobile == "" {
		return ""
	}

	// اعمال گزینه‌ها
	option := DefaultOptions
	if len(opts) > 0 {
		option = opts[0]
	}

	// مرحله 1: پاکسازی اولیه
	mobile = strings.TrimSpace(mobile)
	mobile = common.RemoveSpecialChars(mobile)
	mobile = common.ConvertPersianArabicToEnglish(mobile)

	// مرحله 2: حذف پیشوندهای بین‌المللی
	mobile = removeInternationalPrefixes(mobile)

	// مرحله 3: اصلاح نقطه ابتدایی
	if strings.HasPrefix(mobile, ".") {
		mobile = "0" + mobile[1:]
	}

	// مرحله 4: افزودن صفر ابتدایی در صورت نیاز
	if option.AddZeroPrefix && !strings.HasPrefix(mobile, "0") {
		mobile = "0" + mobile
	}

	// مرحله 5: اعتبارسنجی نهایی
	if !isValidFormat(mobile, option) {
		return ""
	}

	return mobile
}

// ConvertToInternational تبدیل شماره موبایل به فرمت بین‌المللی
// خروجی: +989XXXXXXXXX
func ConvertToInternational(mobile string) string {
	cleaned := Sanitize(mobile)
	if cleaned == "" {
		return ""
	}
	return "+98" + cleaned[1:]
}

// ConvertToZeroFormat تبدیل به فرمت با صفر ابتدایی
func ConvertToZeroFormat(mobile string) string {
	return Sanitize(mobile)
}

// IsValid بررسی اعتبار شماره موبایل
func IsValid(mobile string) bool {
	return Sanitize(mobile) != ""
}

// ExtractOperator استخراج اپراتور از شماره موبایل
func ExtractOperator(mobile string) string {
	cleaned := Sanitize(mobile)
	if cleaned == "" {
		return ""
	}

	prefix := cleaned[1:4] // سه رقم بعد از 0

	switch prefix {
	case "910", "911", "912", "913", "914", "915", "916", "917", "918", "919":
		return "همراه اول"
	case "990", "991", "992", "993", "994", "995", "996", "997", "998", "999":
		return "ایرانسل"
	case "920", "921", "922", "923", "924", "925", "926", "927", "928", "929":
		return "رایتل"
	case "930", "931", "932", "933", "934", "935", "936", "937", "938", "939":
		return "همراه اول"
	case "901", "902", "903", "904", "905", "906", "907", "908", "909":
		return "شاتل موبایل"
	default:
		return "نامشخص"
	}
}

// removeInternationalPrefixes حذف پیشوندهای بین‌المللی
func removeInternationalPrefixes(mobile string) string {
	switch {
	case strings.HasPrefix(mobile, "0098"):
		return mobile[4:]
	case strings.HasPrefix(mobile, "098"):
		return mobile[3:]
	case strings.HasPrefix(mobile, "98"):
		return mobile[2:]
	default:
		return mobile
	}
}

// isValidFormat بررسی فرمت نهایی شماره موبایل
func isValidFormat(mobile string, option Option) bool {
	// بررسی تمام رقم بودن
	if !common.IsAllDigits(mobile) {
		return false
	}

	// بررسی طول
	if len(mobile) != 11 {
		return false
	}

	// بررسی فرمت با regex
	var pattern *regexp.Regexp
	if option.StrictMode {
		pattern = regexp.MustCompile(`^09[0-9]{9}$`)
	} else {
		pattern = regexp.MustCompile(`^09\d{9}$`)
	}

	if !pattern.MatchString(mobile) {
		return false
	}

	// بررسی معتبر بودن سه رقم اول
	prefix := mobile[0:3]
	validPrefixes := map[string]bool{
		"091": true, "092": true, "093": true, "094": true,
		"095": true, "096": true, "097": true, "098": true, "099": true,
	}

	if option.ValidatePrefix && !validPrefixes[prefix] {
		return false
	}

	return true
}
