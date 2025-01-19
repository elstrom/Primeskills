package database

import (
	"log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// ConnectDatabase menghubungkan ke database
func ConnectDatabase() {
	// Ganti konfigurasi database sesuai kebutuhan Anda
	dsn := "host=localhost user=eka password=ramdan77 dbname=comments_db port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database: ", err)
	}
	log.Println("Database connected successfully!")
}

// MigrateDatabase untuk migrasi tabel
func MigrateDatabase(models ...interface{}) {
	err := DB.AutoMigrate(models...)
	if err != nil {
		log.Fatal("Failed to migrate database: ", err)
	}
	log.Println("Database migrated successfully!")
}
