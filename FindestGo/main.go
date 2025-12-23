package main

import (
	"database/sql"
	"fmt"
	"time"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	
)

func main() {
	dbUser := "root"
	dbPassword := ""
	dbHost := "localhost"
	dbPort := "3306"
	dbName := "findest_go"

	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.InfoLevel)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		dbUser,
		dbPassword,
		dbHost,
		dbPort,
		dbName,
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.WithError(err).Fatal("Failed to open DB")
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.WithError(err).Fatal("Failed to open DB")
	}

	log.Info("Connected to MySQL successfully")


	r := gin.Default()

r.POST("/transactions", func(c *gin.Context) {
	var tx struct {
		UserID int64 `json:"user_id"`
		Amount float64 `json:"amount"`
		Status string `json:"status"`
	}

	if err := c.ShouldBindJSON(&tx); err != nil {
		log.WithError(err).Warn("Invalid request body")
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	query := `
		INSERT INTO transactions (UserID, Amount, Status, CreatedAt)
		VALUES (?, ?, ?, NOW())
	`

	result, err := db.Exec(query, tx.UserID, tx.Amount, tx.Status)
	if err != nil {
		log.WithError(err).Error("Database query failed")
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	id, _ := result.LastInsertId()

	c.JSON(201, gin.H{
		"message": "Transaction created",
		"id": id,
	})
})

r.GET("/transactions", func(c *gin.Context) {

	userID := c.Query("user_id")
	status := c.Query("status")

	query := `
		SELECT ID, UserID, Amount, Status, CreatedAt
		FROM transactions
		WHERE 1=1
	`
	var args []interface{}

	if userID != "" {
		query += " AND UserID = ?"
		args = append(args, userID)
	}

	if status != "" {
		query += " AND Status = ?"
		args = append(args, status)
	}

	query += " ORDER BY CreatedAt DESC"

	rows, err := db.Query(query, args...)
	if err != nil {
		log.WithError(err).Error("Database query failed")
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var results []gin.H

	for rows.Next() {
		var id, userID int64
		var amount float64
		var status string
		var createdAt time.Time

		err := rows.Scan(&id, &userID, &amount, &status, &createdAt)
		if err != nil {
			log.WithError(err).Error("Database query failed")
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		results = append(results, gin.H{
			"id":         id,
			"user_id":    userID,
			"amount":     amount,
			"status":     status,
			"created_at": createdAt,
		})
	}

	c.JSON(200, results)
})

r.GET("/transactions/:id", func(c *gin.Context) {

	idParam := c.Param("id")

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid transaction ID"})
		return
	}

	var transaction struct {
		ID int64
		UserID int64
		Amount float64
		Status string
		CreatedAt time.Time
	}

	err = db.QueryRow(`
		SELECT ID, UserID, Amount, Status, CreatedAt
		FROM transactions
		WHERE ID = ?
	`, id).Scan(
		&transaction.ID,
		&transaction.UserID,
		&transaction.Amount,
		&transaction.Status,
		&transaction.CreatedAt,
	)

	if err == sql.ErrNoRows {
		c.JSON(404, gin.H{"error": "Transaction not found"})
		return
	}

	if err != nil {
		log.WithError(err).Error("Database query failed")
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"id": transaction.ID,
		"user_id": transaction.UserID,
		"amount": transaction.Amount,
		"status": transaction.Status,
		"created_at": transaction.CreatedAt,
	})
})

r.PUT("/transactions/:id", func(c *gin.Context) {

	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid transaction ID"})
		return
	}

	var body struct {
		Status string `json:"status"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		log.WithError(err).Warn("Invalid request body")
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if body.Status == "" {
		c.JSON(400, gin.H{"error": "Status is required"})
		return
	}

	result, err := db.Exec(`
		UPDATE transactions
		SET Status = ?
		WHERE ID = ?
	`, body.Status, id)

	if err != nil {
		log.WithError(err).Error("Database query failed")
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(404, gin.H{"error": "Transaction not found"})
		return
	}

	c.JSON(200, gin.H{
		"message": "Transaction status updated",
		"id": id,
		"status": body.Status,
	})
})

r.DELETE("/transactions/:id", func(c *gin.Context) {

	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid transaction ID"})
		return
	}

	result, err := db.Exec(`
		DELETE FROM transactions
		WHERE ID = ?
	`, id)

	if err != nil {
		log.WithError(err).Error("Database query failed")
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(404, gin.H{"error": "Transaction not found"})
		return
	}

	c.JSON(200, gin.H{
		"message": "Transaction deleted successfully",
		"id": id,
	})
})

r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "API + DB connected",
		})
	})

r.GET("/dashboard/summary", func(c *gin.Context) {

	var totalSuccess_today int64
	db.QueryRow(
		"SELECT COUNT(*) FROM transactions WHERE Status = 'success' AND DATE(CreatedAt) = CURDATE();",
	).Scan(&totalSuccess_today)

	rowsAvg, err := db.Query(`
		SELECT UserID, AVG(Amount)
		FROM transactions
		GROUP BY UserID
	`)
	if err != nil {
		log.WithError(err).Error("Database query failed")
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	defer rowsAvg.Close()

	avgPerUser := make(map[int64]float64)

	for rowsAvg.Next() {
		var userID int64
		var avg float64
		rowsAvg.Scan(&userID, &avg)
		avgPerUser[userID] = avg
	}

	rows, err := db.Query(`
		SELECT ID, UserID, Amount, Status, CreatedAt
		FROM transactions
		ORDER BY CreatedAt DESC
		LIMIT 10
	`)
	if err != nil {
		log.WithError(err).Error("Database query failed")
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var latest []map[string]interface{}

	for rows.Next() {
		var id, userID int64
		var amount float64
		var status string
		var createdAt string

		rows.Scan(&id, &userID, &amount, &status, &createdAt)

		latest = append(latest, gin.H{
			"id": id,
			"user_id": userID,
			"amount": amount,
			"status": status,
			"created_at": createdAt,
		})
	}

	c.JSON(200, gin.H{
		"Total Success Today": totalSuccess_today,
		"Average Amount Per User": avgPerUser,
		"10 Latest Transactions": latest,
	})
})

r.Run(":8080")
}
