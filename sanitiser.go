package iransanitize

import (
	"github.com/mrrashidpour/iransanitize/mobile"
)

// Mobile Sanitizer Functions
type MobileOption = mobile.Option

var (
	DefaultMobileOptions = mobile.DefaultOptions
	LenientMobileOptions = mobile.LenientOptions
	StrictMobileOptions  = mobile.StrictOptions
)

// SanitizeMobile شماره موبایل را پاکسازی می‌کند
func SanitizeMobile(mobileStr string, opts ...MobileOption) string {
	return mobile.Sanitize(mobileStr, opts...)
}

// ConvertMobileToInternational تبدیل به فرمت بین‌المللی
func ConvertMobileToInternational(mobileStr string) string {
	return mobile.ConvertToInternational(mobileStr)
}

// IsValidMobile بررسی اعتبار شماره موبایل
func IsValidMobile(mobileStr string) bool {
	return mobile.IsValid(mobileStr)
}

// ExtractMobileOperator استخراج اپراتور
func ExtractMobileOperator(mobileStr string) string {
	return mobile.ExtractOperator(mobileStr)
}

// MaskMobile ماسک کردن شماره موبایل
func MaskMobile(mobileStr string) string {
	return mobile.Mask(mobileStr)
}

// CompareMobile مقایسه دو شماره موبایل
func CompareMobile(mobile1, mobile2 string) bool {
	return mobile.Compare(mobile1, mobile2)
}
