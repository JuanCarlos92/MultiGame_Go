package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"juego/models"
	"log"
)

var DB *gorm.DB

func Connect() {
	username := "compa"
	password := "1234"
	host := "192.168.1.11" // tu IP local
	port := "3306"
	dbname := "juego"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		username, password, host, port, dbname)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Error al conectar con la base de datos:", err)
	}

	err = DB.AutoMigrate(&models.Jugador{})
	if err != nil {
		log.Fatal("❌ Error al migrar modelos:", err)
	}

	log.Println("✅ Conexión exitosa a la base de datos.")
}
func Close() error {
	if DB != nil {
		sqlDB, err := DB.DB()
		if err != nil {
			return err
		}
		return sqlDB.Close()
	}
	return nil
}
