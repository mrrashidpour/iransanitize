package mobile

// Mask شماره موبایل را ماسک می‌کند (برای نمایش در لاگ‌ها یا UI)
// مثال: 0912***6789
func Mask(mobile string) string {
	cleaned := Sanitize(mobile)
	if cleaned == "" {
		return ""
	}

	if len(cleaned) != 11 {
		return cleaned
	}

	return cleaned[:4] + "***" + cleaned[7:]
}

// Compare دو شماره موبایل را بدون در نظر گرفتن فرمت مقایسه می‌کند
func Compare(mobile1, mobile2 string) bool {
	clean1 := Sanitize(mobile1)
	clean2 := Sanitize(mobile2)

	if clean1 == "" || clean2 == "" {
		return false
	}

	return clean1 == clean2
}

// NormalizeBatch پاکسازی چندین شماره موبایل به صورت همزمان
func NormalizeBatch(mobiles []string, opts ...Option) []string {
	result := make([]string, 0, len(mobiles))
	option := DefaultOptions
	if len(opts) > 0 {
		option = opts[0]
	}

	for _, m := range mobiles {
		cleaned := Sanitize(m, option)
		if cleaned != "" {
			result = append(result, cleaned)
		}
	}

	return result
}

// ExtractUnique استخراج شماره‌های یکتا از یک لیست
func ExtractUnique(mobiles []string, opts ...Option) []string {
	normalized := NormalizeBatch(mobiles, opts...)

	seen := make(map[string]bool)
	unique := make([]string, 0, len(normalized))

	for _, m := range normalized {
		if !seen[m] {
			seen[m] = true
			unique = append(unique, m)
		}
	}

	return unique
}

// FormatWithDash شماره موبایل را با خط تیره فرمت می‌کند
// مثال: 0912-345-6789
func FormatWithDash(mobile string) string {
	cleaned := Sanitize(mobile)
	if cleaned == "" {
		return ""
	}

	if len(cleaned) != 11 {
		return cleaned
	}

	return cleaned[:4] + "-" + cleaned[4:7] + "-" + cleaned[7:]
}

// FormatWithSpace با فاصله فرمت می‌کند
// مثال: 0912 345 6789
func FormatWithSpace(mobile string) string {
	cleaned := Sanitize(mobile)
	if cleaned == "" {
		return ""
	}

	if len(cleaned) != 11 {
		return cleaned
	}

	return cleaned[:4] + " " + cleaned[4:7] + " " + cleaned[7:]
}

// IsIranCell بررسی می‌کند که شماره متعلق به ایرانسل باشد
func IsIranCell(mobile string) bool {
	operator := ExtractOperator(mobile)
	return operator == "ایرانسل"
}

// IsHamrahAvval بررسی می‌کند که شماره متعلق به همراه اول باشد
func IsHamrahAvval(mobile string) bool {
	operator := ExtractOperator(mobile)
	return operator == "همراه اول"
}

// IsRightel بررسی می‌کند که شماره متعلق به رایتل باشد
func IsRightel(mobile string) bool {
	operator := ExtractOperator(mobile)
	return operator == "رایتل"
}
