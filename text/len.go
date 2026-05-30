package text

// isPersianCharacter بررسی می‌کند کاراکتر فارسی است یا خیر
func isPersianCharacter(r rune) bool {
	return (r >= 0x0600 && r <= 0x06FF) || // عربی
		(r >= 0xFB50 && r <= 0xFDFF) || // فرم‌های نمایشی عربی - فارسی
		(r >= 0xFE70 && r <= 0xFEFF) || // فرم‌های نمایشی عربی - دیاکریتیک
		(r >= 0x0750 && r <= 0x077F) || //扩展 عربی
		(r == 0x200C || r == 0x200D) // نیم‌فاصله
}

// CountChars تعداد کاراکترهای فارسی را برمی‌گرداند
func CountChars(text string) int {
	count := 0
	for _, ch := range text {
		if isPersianCharacter(ch) {
			count++
		}
	}
	return count
}
