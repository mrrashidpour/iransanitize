package mobile

// Option تنظیمات پیکربندی برای پاکسازی شماره موبایل
type Option struct {
	// AddZeroPrefix در صورت نبود صفر ابتدایی، به طور خودکار اضافه شود
	AddZeroPrefix bool

	// StrictMode بررسی دقیق‌تر (فقط اعداد انگلیسی مجاز)
	StrictMode bool

	// ValidatePrefix بررسی اعتبار پیشوند (مثل 091، 093، ...)
	ValidatePrefix bool

	// AcceptInternational پذیرش فرمت بین‌المللی در خروجی
	AcceptInternational bool
}

// DefaultOptions گزینه‌های پیش‌فرض
var DefaultOptions = Option{
	AddZeroPrefix:       true,
	StrictMode:          false,
	ValidatePrefix:      true,
	AcceptInternational: false,
}

// LenientOptions گزینه‌های宽松 (برای ورودی‌های نامطمئن)
var LenientOptions = Option{
	AddZeroPrefix:       true,
	StrictMode:          false,
	ValidatePrefix:      false,
	AcceptInternational: false,
}

// StrictOptions گزینه‌های سخت‌گیرانه
var StrictOptions = Option{
	AddZeroPrefix:       true,
	StrictMode:          true,
	ValidatePrefix:      true,
	AcceptInternational: false,
}
