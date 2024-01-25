package main

import (
	handlers "go-redis/handlers/product"
	repositores "go-redis/repositories/product"
	service "go-redis/services/product"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
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
	 _ = redisClient

	productRepo := repositores.NewProductRepositoryDB(db)
	productService := service.NewCatalogServiceRedis(productRepo,redisClient)
	productHandler := handlers.NewCatalogHandler(productService)

	app := fiber.New()

	app.Get("/products",  productHandler.GetProducts)

	app.Listen(":8888")
}

func initDatabase() *gorm.DB {
	dial := mysql.Open("root:P@ssw0rd@tcp(localhost:3307)/redisgo?charset=utf8&parseTime=True&loc=Local")

	db, err := gorm.Open(dial, &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

func initRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
}
