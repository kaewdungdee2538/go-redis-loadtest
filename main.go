package main

import (
	"fmt"
	repositores "go-redis/repositories/product"

	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// "time"

// fiber "github.com/gofiber/fiber/v2"

func main() {
	// app := fiber.New()

	// app.Get("/hello",func(c *fiber.Ctx) error {
	// 	time.Sleep(time.Millisecond * 10)
	// 	return c.SendString("Hello word")
	// })
	// app.Listen(":8000")

	db := initDatabase()
	redisClient := initRedis()

	productRepo := repositores.NewProductRepositoryRedis(db,redisClient)
	products, err := productRepo.GetProducts()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(products)
}

func initDatabase() *gorm.DB {
	dial := mysql.Open("root:P@ssw0rd@tcp(localhost:3306)/redisgo?charset=utf8&parseTime=True&loc=Local")

	db, err := gorm.Open(dial, &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

func initRedis() *redis.Client{
	return redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
}
