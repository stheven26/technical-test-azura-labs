package controllers

import (
	"net/http"
	"tech-test/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var Produk []models.Produk

func BuatProduk(c *gin.Context) {
	var data models.Produk
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	data.ID = uuid.New().String()
	Produk = append(Produk, data)

	c.JSON(http.StatusOK, gin.H{
		"time":   time.Now(),
		"status": "Berhasil Membuat Data Produk",
	})
}

func LihatDaftarProduk(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": Produk,
		"time": time.Now(),
	})
}

func LihatProdukByID(c *gin.Context) {
	id := c.Param("id")
	for i, e := range Produk {
		if e.ID == id {
			c.JSON(http.StatusOK, gin.H{
				"data": Produk[i],
				"time": time.Now(),
			})
			return
		}
	}

	c.JSON(http.StatusBadRequest, gin.H{
		"error": "Invalid Id",
	})

}

func EditProduk(c *gin.Context) {
	id := c.Param("id")
	var data models.Produk
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	for i, e := range Produk {
		if e.ID == id {
			Produk[i].Nama = data.Nama
			Produk[i].Detail = data.Detail
			Produk[i].Harga = data.Harga
			Produk[i].Rating = data.Rating
			Produk[i].Likes = data.Likes
			c.JSON(http.StatusOK, gin.H{
				"data":   Produk[i],
				"time":   time.Now(),
				"status": "Berhasil Edit Data Produk",
			})
			return
		}
	}
}

func HapusProduk(c *gin.Context) {
	id := c.Param("id")
	for i, e := range Produk {
		if e.ID == id {
			Produk = append(Produk[:i], Produk[i+1:]...)

			c.JSON(http.StatusOK, gin.H{
				"time":   time.Now(),
				"status": "Berhasil Delete Data Produk",
			})
			return
		}
	}

	c.JSON(http.StatusBadRequest, gin.H{
		"error": "Invalid Id",
	})
}
