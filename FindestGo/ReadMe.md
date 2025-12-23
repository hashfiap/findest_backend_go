SETUP Database
1. Download dan Install XAMPP
2. Jalankan XAMPP dan Start bagian "MySQL" dan "Apache"
3. Check di browser dengan url "localhost/phpmyadmin"
4. Klik "New" pada bagian kiri untuk menambahkan database baru
5. Tambahkan tabel dengan query
    "CREATE TABLE `user` (
    `ID` INT(3) NOT NULL AUTO_INCREMENT,
    `Nama` VARCHAR(64) NOT NULL,
    PRIMARY KEY (`ID`));"

    serta

    "CREATE TABLE `transactions` (
    `ID` INT(3) NOT NULL AUTO_INCREMENT,
    `UserID` INT(3) NOT NULL,
    `Amount` FLOAT NOT NULL,
    `Status` VARCHAR(64) NOT NULL,
    `CreatedAt` DATETIME NOT NULL,
    PRIMARY KEY (`ID`),
    FOREIGN KEY (`UserID`) REFERENCES user(`ID`));"

SETUP API
1. Download dan Install Postman
2. New file bernama "main.go" dalam Visual Studio Code dengan programming language "Go" (Golang) 
3. Ketik "go mod init /(nama folder)" serta get github yang dibutuhkan dalam import "go get (link github)"
3. Berikan Code untuk connect kepada API dan Database yang sudah dibuat, berikut code nya
package main

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gin-gonic/gin"
)

func main() {
	dbUser := "root"
	dbPassword := ""
	dbHost := "localhost"
	dbPort := "3306"
	dbName := "findest_go"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		dbUser,
		dbPassword,
		dbHost,
		dbPort,
		dbName,
	)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Failed to open DB:", err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Fatal("Failed to connect to DB:", err)
	}
	fmt.Println("Connected to MySQL successfully")

r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "API + DB connected",
		})
	})
    r.Run(":8080")
}

4. Setelah nya, lakukan "go run main.go" dalam terminal bash di Visual Studio Code
5. Lalu buka UI Postman dan pilih GET serta ketik "http://localhost:8080/ping" lalu klik tombol Send
6. Sebuah pesan akan tertera sesuai dengan variabel "message" yang sudah diketikkan di dalam main.go
7. Setelah verifikasi koneksi, dapat menambahkan endpoint sesuai dengan kebutuhan

MENJALANKAN Project
1. Run MySQL, Postman, Visual Studio Code
2. Ketik "go run main.go" di terminal bash VS Code

3. Ketikkan di UI Postman sesuai dengan endpoint yang sudah tersedia (contoh GET "http://localhost:8080/dashboard/summary" atau POST "http://localhost:8080/transactions" dan ketik didalam Body -> raw (Language JSON) sesuai dengan column tabel) untuk CRUD kedalam database MySQL
