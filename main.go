package main

import (
	"fmt"
	"soc-tool/internal/api"
	"soc-tool/internal/database"
	"soc-tool/internal/ingestion"
)

func main() {
	fmt.Println("🚀 Starting SOC Tool...")

	// تهيئة قاعدة البيانات
	db := database.InitDB()

	// تشغيل ingestion (إضافة حدث تجريبي أو قراءة logs)
	ingestion.Start(db)

	// تشغيل API
	api.StartServer(db)
}
