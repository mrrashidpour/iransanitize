package date

import (
	"fmt"
	"strconv"
	"strings"

	jalaali "github.com/jalaali/go-jalaali"
	"github.com/mrrashidpour/iransanitize/internal/common"
)

// SanitizeDate تاریخ را به فرمت استاندارد YYYY-MM-DD (میلادی) تبدیل می‌کند
func SanitizeDate(date string) string {
	if date == "" {
		return ""
	}

	// مرحله 1: پاکسازی اولیه
	raw := strings.TrimSpace(date)
	raw = common.ConvertToEnglishDigits(raw)
	raw = strings.ReplaceAll(raw, "/", "-")

	// مرحله 2: جدا کردن بخش‌ها
	parts := strings.Split(raw, "-")
	if len(parts) != 3 {
		return ""
	}

	// مرحله 3: تبدیل به اعداد
	a, _ := strconv.Atoi(parts[0])
	b, _ := strconv.Atoi(parts[1])
	c, _ := strconv.Atoi(parts[2])

	// مرحله 4: تشخیص الگوی سال-ماه-روز یا روز-ماه-سال
	var year, month, day int
	switch {
	case a > 31:
		year, month, day = a, b, c
	case c > 31:
		year, month, day = c, b, a
	default:
		return ""
	}

	// مرحله 6: اگر سال کمتر از 1300 است، تاریخ شمسی را به میلادی تبدیل کن
	if year < 1300 {
		return ""
	}

	// مرحله 5: اگر سال کمتر از 1500 است، تاریخ شمسی را به میلادی تبدیل کن
	if year < 1500 {
		gYear, gMonth, gDay, err := jalaali.ToGregorian(year, jalaali.Month(month), day)
		if err != nil {
			return ""
		}
		if gYear < 1821 || gYear > 2100 {
			return ""
		}
		return fmt.Sprintf("%04d-%02d-%02d", gYear, gMonth, gDay)
	}

	// مرحله 6: اعتبارسنجی تاریخ میلادی
	if !isValidGregorian(year, month, day) {
		return ""
	}

	// مرحله 7: برگرداندن تاریخ میلادی به فرمت استاندارد
	return fmt.Sprintf("%04d-%02d-%02d", year, month, day)
}

// isValidGregorian بررسی اعتبار تاریخ میلادی
func isValidGregorian(year, month, day int) bool {

	if year < 1821 || year > 2100 {
		return false
	}
	if month < 1 || month > 12 {
		return false
	}
	if day < 1 || day > 31 {
		return false
	}

	// بررسی روزهای هر ماه
	monthDays := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

	// اصلاح برای سال کبیسه میلادی
	if month == 2 && isGregorianLeapYear(year) {
		monthDays[1] = 29
	}

	return day <= monthDays[month-1]
}

// isGregorianLeapYear بررسی کبیسه بودن سال میلادی
func isGregorianLeapYear(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}

// IsValidDate بررسی اعتبار تاریخ
func IsValidDate(date string) bool {
	return SanitizeDate(date) != ""
}
