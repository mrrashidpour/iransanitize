package date

// OutputFormat نوع فرمت خروجی
type OutputFormat int

const (
	Auto      OutputFormat = iota // خودکار (تشخیص بر اساس ورودی)
	Jalali                        // خروجی شمسی
	Gregorian                     // خروجی میلادی
)

// Option تنظیمات پیکربندی برای پاکسازی تاریخ
type Option struct {
	// OutputFormat فرمت خروجی
	OutputFormat OutputFormat

	// MinYear حداقل سال مجاز (پیش‌فرض: 1300)
	MinYear int

	// MaxYear حداکثر سال مجاز (پیش‌فرض: 1500 برای شمسی، 2100 برای میلادی)
	MaxYear int

	// StrictMode بررسی دقیق‌تر (فقط فرمت YYYY-MM-DD قبول شود)
	StrictMode bool

	// AcceptJalali قبول تاریخ شمسی (پیش‌فرض: true)
	AcceptJalali bool

	// AcceptGregorian قبول تاریخ میلادی (پیش‌فرض: true)
	AcceptGregorian bool
}

// DefaultOptions گزینه‌های پیش‌فرض
var DefaultOptions = Option{
	OutputFormat:    Auto,
	MinYear:         1300,
	MaxYear:         1500,
	StrictMode:      false,
	AcceptJalali:    true,
	AcceptGregorian: true,
}

// GregorianOnlyOptions فقط تاریخ میلادی
var GregorianOnlyOptions = Option{
	OutputFormat:    Gregorian,
	MinYear:         1900,
	MaxYear:         2100,
	StrictMode:      false,
	AcceptJalali:    false,
	AcceptGregorian: true,
}

// JalaliOnlyOptions فقط تاریخ شمسی
var JalaliOnlyOptions = Option{
	OutputFormat:    Jalali,
	MinYear:         1300,
	MaxYear:         1500,
	StrictMode:      false,
	AcceptJalali:    true,
	AcceptGregorian: false,
}

// StrictOptions گزینه‌های سخت‌گیرانه
var StrictOptions = Option{
	OutputFormat:    Auto,
	MinYear:         1300,
	MaxYear:         1500,
	StrictMode:      true,
	AcceptJalali:    true,
	AcceptGregorian: true,
}
