package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"optumlabs.com/myapi/internal/entities"
	"optumlabs.com/myapi/internal/product"
	"optumlabs.com/myapi/pkg/dbcontext"
)

func main() {
	db, err := gorm.Open(sqlite.Open("app.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&entities.Product{})

	app := gin.New()

	app.GET("/healthz", func(c *gin.Context) {
		c.String(200, "ok")
	})

	dbctx := dbcontext.New(db)
	product.RegisterHandler(app.Group("/products"), product.NewService(product.NewRepository(dbctx)))

	app.Run()
}
