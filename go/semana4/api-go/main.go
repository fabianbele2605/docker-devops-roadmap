package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/redis/go-redis/v9"
)

type User struct {
	ID         int    `json:"id"`
	Name       string  `json:"name"`
	Email      string  `json:"email"`
}

func main() {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("Error conectando a DB:", err)
	}
	defer conn.Close(context.Background())

	opt, err := redis.ParseURL(os.Getenv("REDIS_URL"))
	if err != nil {
		log.Fatal("Error parseando REDIS_URL:", err)
	}
	rdb := redis.NewClient(opt)

	createTableSQL := `
	CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			name TEXT NOT NULL,
			email TEXT UNIQUE NOT NULL
	);`

	if err := conn.Ping(context.Background()); err != nil {
		log.Fatal("DB no responde:", err)
	}
	log.Println("DB conectada")

	if err := rdb.Ping(context.Background()).Err(); err != nil {
		log.Fatal("Redis no responde:", err)
	}
	log.Println("Redis conectado")

	_, err = conn.Exec(context.Background(), createTableSQL)

	if err != nil {
		log.Fatal("Error creando tabla:", err)
	}

	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	r.POST("/users", func(c *gin.Context) {
		var user User
		
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		query := "INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id"
		err := conn.QueryRow(context.Background(), query, user.Name, user.Email).Scan(&user.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo insertar el usuario"})
			return
		}

		rdb.Del(context.Background(), "users")

		c.JSON(http.StatusCreated, user)
	})

	r.GET("/users", func(c *gin.Context) {
		cached, err := rdb.Get(context.Background(), "users").Result()
		if err == nil {
			c.Data(http.StatusOK, "application/json", []byte(cached))
			return
		}
		query := "SELECT id, name, email FROM users"
		rows, err := conn.Query(context.Background(), query)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo encontrar el usuario"})
			return
		}

		defer rows.Close()

		users := []User{}

		for rows.Next() {
			var u User
			if err := rows.Scan(&u.ID, &u.Name, &u.Email); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al escanear los datos"})
				return
			}
			users = append(users, u)
		}

		if err = rows.Err(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error durante la lectura de la fila"})
			return
		}

		data, _ := json.Marshal(users)
		rdb.Set(context.Background(), "users", data, 30*time.Second)

		c.JSON(http.StatusOK, users)
	})

	r.Run(":8080")
}