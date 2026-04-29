package date

import (
	"fmt"
	"strconv"
	"strings"
)

// SanitizeDate تاریخ را به فرمت استاندارد YYYY-MM-DD تبدیل می‌کند
func SanitizeDate(date string) string {
	if date == "" {
		return ""
	}

	// مرحله 1: پاکسازی اولیه
	raw := strings.TrimSpace(date)
	raw = normalizeDigits(raw)
	raw = strings.ReplaceAll(raw, "/", "-")

	// مرحله 2: استخراج اعداد
	numbers := extractNumbers(raw)
	if len(numbers) != 3 {
		return ""
	}

	// مرحله 3: تشخیص الگو
	year, month, day := detectPatternStrict(numbers)
	if year == 0 {
		return ""
	}

	// مرحله 4: اعتبارسنجی
	if !isValidDateParts(year, month, day) {
		return ""
	}

	// مرحله 5: برگرداندن به فرمت استاندارد
	return fmt.Sprintf("%04d-%02d-%02d", year, month, day)
}

// extractNumbers استخراج تمام اعداد از رشته
func extractNumbers(s string) []int {
	var numbers []int
	var current string

	for _, ch := range s {
		if ch >= '0' && ch <= '9' {
			current += string(ch)
		} else if current != "" {
			num, _ := strconv.Atoi(current)
			numbers = append(numbers, num)
			current = ""
		}
	}

	if current != "" {
		num, _ := strconv.Atoi(current)
		numbers = append(numbers, num)
	}

	return numbers
}

// detectPatternStrict تشخیص دقیق الگوی تاریخ
func detectPatternStrict(numbers []int) (year, month, day int) {
	if len(numbers) != 3 {
		return 0, 0, 0
	}

	a, b, c := numbers[0], numbers[1], numbers[2]

	// گزینه 1: YYYY-MM-DD (سال در ابتدا)
	if a > 31 && b >= 1 && b <= 12 && c >= 1 && c <= 31 {
		// باید مطمئن شویم که b و c معتبر هستند
		if isValidMonthDay(b, c, a) {
			return a, b, c
		}
	}

	// گزینه 2: DD-MM-YYYY (سال در انتها)
	if c > 31 && b >= 1 && b <= 12 && a >= 1 && a <= 31 {
		if isValidMonthDay(b, a, c) {
			return c, b, a
		}
	}

	// گزینه 3: YYYY-DD-MM (نادر - سال در ابتدا، روز در وسط)
	if a > 31 && c >= 1 && c <= 12 && b >= 1 && b <= 31 {
		if isValidMonthDay(c, b, a) {
			return a, c, b
		}
	}

	// گزینه 4: MM-DD-YYYY (ماه در ابتدا)
	if b <= 12 && c > 31 && a <= 31 {
		if isValidMonthDay(a, b, c) {
			return c, a, b
		}
	}

	return 0, 0, 0
}

// isValidMonthDay بررسی اعتبار ماه و روز
func isValidMonthDay(month, day, year int) bool {
	if month < 1 || month > 12 {
		return false
	}
	if day < 1 || day > 31 {
		return false
	}

	// تعیین حداکثر روز بر اساس نوع تقویم
	var maxDay int

	// سال شمسی
	if year >= 1300 && year <= 1499 {
		jalaliDays := []int{31, 31, 31, 31, 31, 31, 30, 30, 30, 30, 30, 29}
		maxDay = jalaliDays[month-1]

		// اصلاح برای اسفند سال کبیسه
		if month == 12 && isJalaliLeapYear(year) {
			maxDay = 30
		}
	} else if year >= 1500 && year <= 2100 {
		// سال میلادی
		gregorianDays := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
		maxDay = gregorianDays[month-1]

		// اصلاح برای فوریه سال کبیسه
		if month == 2 && isGregorianLeapYear(year) {
			maxDay = 29
		}
	} else {
		// سال خارج از محدوده
		return false
	}

	return day <= maxDay
}

// isValidDateParts اعتبارسنجی اجزای تاریخ
func isValidDateParts(year, month, day int) bool {
	// بررسی محدوده سال
	if year < 1300 || year > 2100 {
		return false
	}

	// بررسی ماه
	if month < 1 || month > 12 {
		return false
	}

	// بررسی روز
	if day < 1 || day > 31 {
		return false
	}

	// بررسی دقیق با در نظر گرفتن نوع تقویم
	return isValidMonthDay(month, day, year)
}

// isJalaliLeapYear بررسی کبیسه بودن سال شمسی
func isJalaliLeapYear(year int) bool {
	rem := year % 33
	return rem == 1 || rem == 5 || rem == 9 || rem == 13 || rem == 17 || rem == 22 || rem == 26 || rem == 30
}

// isGregorianLeapYear بررسی کبیسه بودن سال میلادی
func isGregorianLeapYear(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}

// normalizeDigits تبدیل اعداد فارسی/عربی به انگلیسی
func normalizeDigits(s string) string {
	replacer := strings.NewReplacer(
		"۰", "0", "۱", "1", "۲", "2", "۳", "3", "۴", "4",
		"۵", "5", "۶", "6", "۷", "7", "۸", "8", "۹", "9",
		"٠", "0", "١", "1", "٢", "2", "٣", "3", "٤", "4",
		"٥", "5", "٦", "6", "٧", "7", "٨", "8", "٩", "9",
	)
	return replacer.Replace(s)
}

// IsValidDate بررسی اعتبار تاریخ
func IsValidDate(date string) bool {
	return SanitizeDate(date) != ""
}
