package database

import "fmt"

// InitDB وظيفتها تهيئة الاتصال بقاعدة البيانات
func InitDB() string {
	fmt.Println("✅ [Database] Initialized and connected successfully...")

	// بنرجع نص تجريبي حالياً لتمثيل الاتصال
	return "Active_DB_Connection"
}
