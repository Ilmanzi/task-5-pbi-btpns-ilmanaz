package controllers

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/Ilmanzi/task-5-pbi-btpns-ilmanaz.git/app"
	db "github.com/Ilmanzi/task-5-pbi-btpns-ilmanaz.git/database"
	"github.com/Ilmanzi/task-5-pbi-btpns-ilmanaz.git/models"
	"github.com/gin-gonic/gin"
)

func UploadPicture(ctx *gin.Context) {
	var reqBody app.PictureRequestBody
	userID, _ := ctx.Get("userID")

	file, err := ctx.FormFile("picture_url")

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Missing picture file"})
		return
	}

	timestamp := time.Now().UnixNano()
	filePicture := fmt.Sprintf("%d-%s", timestamp, strings.ReplaceAll(file.Filename, " ", ""))
	filename := "http://localhost:8080/api/v1/pictures/" + filePicture

	reqBody.Title = ctx.PostForm("title")
	reqBody.Caption = ctx.PostForm("caption")
	reqBody.PictureUrl = filename

	picture := models.Picture{
		Title:      reqBody.Title,
		Caption:    reqBody.Caption,
		PictureUrl: reqBody.PictureUrl,
		UserID:     userID.(string),
	}

	db := db.Init()

	err = ctx.SaveUploadedFile(file, "./uploads/"+filePicture)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}

	if err := db.Create(&picture).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Picture uploaded successfully",
		"success": true,
	})
}

func GetPicture(ctx *gin.Context) {
	userID, _ := ctx.Get("userID")

	db := db.Init()

	var pictures []models.Picture
	result := db.Where("user_id = ?", userID).Preload("User").Find(&pictures)

	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to fetch pictures",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Successfully retrieve data",
		"data":    pictures,
	})
}

func UpdatePicture(ctx *gin.Context) {
	pictureID := ctx.Param("id")
	userID, _ := ctx.Get("userID")

	file, err := ctx.FormFile("picture_url")

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Missing picture file"})
		return
	}

	timestamp := time.Now().UnixNano()
	filePicture := fmt.Sprintf("%d-%s", timestamp, strings.ReplaceAll(file.Filename, " ", ""))
	filename := "http://localhost:8080/api/v1/pictures/" + filePicture

	updateData := app.PictureRequestBody{
		Title:      ctx.PostForm("title"),
		Caption:    ctx.PostForm("caption"),
		PictureUrl: filename,
	}

	picture := models.Picture{
		Title:      updateData.Title,
		Caption:    updateData.Caption,
		PictureUrl: updateData.PictureUrl,
	}

	db := db.Init()
	result := db.Where("id = ?", pictureID).Where("user_id", userID).Preload("User").First(&picture)

	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "Picture not found",
			"success": false,
		})
		return
	}

	if err := db.Save(&picture).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Picture updated successfully",
		"data":    picture,
	})
}

func DeletePicture(ctx *gin.Context) {
	pictureID := ctx.Param("id")
	userID, _ := ctx.Get("userID")

	var picture models.Picture

	db := db.Init()

	if err := db.Where("id = ?", pictureID).Where("user_id = ?", userID).First(&picture).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "Picture not found",
		})
		return
	}

	if err := db.Where("id = ?", pictureID).Where("user_id = ?", userID).Unscoped().Delete(&picture).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to delete picture",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Picture successfully deleted",
	})
}
