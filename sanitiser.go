package iransanitize

import (
	"github.com/mrrashidpour/iransanitize/date"
	"github.com/mrrashidpour/iransanitize/mobile"
)

// ============ Mobile Functions ============
type MobileOption = mobile.Option

var (
	DefaultMobileOptions = mobile.DefaultOptions
	LenientMobileOptions = mobile.LenientOptions
	StrictMobileOptions  = mobile.StrictOptions
)

func SanitizeMobile(mobileStr string, opts ...MobileOption) string {
	return mobile.Sanitize(mobileStr, opts...)
}

func ConvertMobileToInternational(mobileStr string) string {
	return mobile.ConvertToInternational(mobileStr)
}

func IsValidMobile(mobileStr string) bool {
	return mobile.IsValid(mobileStr)
}

func ExtractMobileOperator(mobileStr string) string {
	return mobile.ExtractOperator(mobileStr)
}

func MaskMobile(mobileStr string) string {
	return mobile.Mask(mobileStr)
}

func CompareMobile(mobile1, mobile2 string) bool {
	return mobile.Compare(mobile1, mobile2)
}

// ============ Date Functions ============
// SanitizeDate تاریخ را به فرمت استاندارد YYYY-MM-DD تبدیل می‌کند
func SanitizeDate(dateStr string) string {
	return date.SanitizeDate(dateStr)
}

// IsValidDate بررسی اعتبار تاریخ
func IsValidDate(dateStr string) bool {
	return date.IsValidDate(dateStr)
}

// ConvertDateToJalali تبدیل تاریخ میلادی به شمسی (اختیاری)
func ConvertDateToJalali(dateStr string) string {
	// این تابع را بعداً اضافه می‌کنیم
	// فعلاً کامنت شده
	return ""
}
