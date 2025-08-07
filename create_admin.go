package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	// 数据库连接配置
	// 请根据您的实际数据库配置进行修改
	// 格式： "username:password@tcp(host:port)/database_name?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := "username:password@tcp(host:port)/database_name?charset=utf8mb4&parseTime=True&loc=Local"

	// 连接到数据库
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// 检查数据库连接
	err = db.Ping()
	if err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}
	fmt.Println("Successfully connected to the database!")

	// 创建管理员账号
	email := "admin@example.com"
	password := "admin123" // 请使用一个强密码，并妥善保管
	name := "Super Admin"

	// 检查用户是否已存在
	var existingUser struct {
		ID int
	}
	err = db.QueryRow("SELECT id FROM users WHERE email = ?", email).Scan(&existingUser.ID)
	if err == nil {
		fmt.Printf("Admin user with email %s already exists. Skipping creation.\n", email)
		return
	} else if err != sql.ErrNoRows {
		log.Fatalf("Error checking for existing user: %v", err)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("Failed to hash password: %v", err)
	}

	// 插入管理员账号
	_, err = db.Exec("INSERT INTO users (name, email, password, role, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)",
		name, email, string(hashedPassword), "admin", time.Now(), time.Now())
	if err != nil {
		log.Fatalf("Failed to create admin user: %v", err)
	}

	fmt.Printf("Admin user '%s' with email '%s' has been successfully added.\n", name, email)
}
