# IranSanitize 🇮🇷

[![Go Reference](https://pkg.go.dev/badge/github.com/mrrashidpour/iransanitize.svg)](https://pkg.go.dev/github.com/mrrashidpour/iransanitize)
[![Go Report Card](https://goreportcard.com/badge/github.com/mrrashidpour/iransanitize)](https://goreportcard.com/report/github.com/mrrashidpour/iransanitize)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)
[![GitHub release](https://img.shields.io/github/v/release/mrrashidpour/iransanitize)](https://github.com/mrrashidpour/iransanitize/releases)

> پکیج قدرتمند Go برای پاکسازی، اعتبارسنجی و استانداردسازی اطلاعات ایرانی

**IranSanitize** یک کتابخانه کارآمد برای پاکسازی و اعتبارسنجی اطلاعات رایج در سیستم‌های ایرانی است. با این پکیج می‌توانید شماره موبایل و تاریخ را به فرمت استاندارد تبدیل کرده و اعتبار آن‌ها را بررسی کنید.

## ✨ ویژگی‌ها

### 📱 شماره موبایل
- ✅ پشتیبانی از اعداد فارسی، عربی و انگلیسی
- ✅ حذف خودکار فاصله، خط تیره و پرانتز
- ✅ پشتیبانی از پیشوندهای بین‌المللی (`+98`, `0098`, `098`)
- ✅ افزودن خودکار صفر ابتدایی
- ✅ تشخیص اپراتور (همراه اول، ایرانسل، رایتل، شاتل موبایل)
- ✅ حالت‌های مختلف اعتبارسنجی (Strict، Lenient)
- ✅ ماسک کردن شماره برای نمایش در UI

### 📅 تاریخ
- ✅ تبدیل خودکار تاریخ شمسی به میلادی
- ✅ پشتیبانی از فرمت‌های مختلف (`YYYY-MM-DD`, `YYYY/MM/DD`, `DD/MM/YYYY`)
- ✅ پشتیبانی از اعداد فارسی و عربی
- ✅ اعتبارسنجی تاریخ‌های شمسی و میلادی
- ✅ پشتیبانی از سال‌های کبیسه

## 📦 نصب

```bash
go get github.com/mrrashidpour/iransanitize@latest