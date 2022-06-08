package routes

import (
	"tech-test/controllers"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func mapUrls() {
	router.POST("/api/produk", controllers.BuatProduk)
	router.GET("/api/produk", controllers.LihatDaftarProduk)
	router.GET("/api/produk/:id", controllers.LihatProdukByID)
	router.PUT("/api/produk/:id", controllers.EditProduk)
	router.DELETE("/api/produk/:id", controllers.HapusProduk)
}

func StartApplication() {
	mapUrls()
	router.Run(":8081")
}
