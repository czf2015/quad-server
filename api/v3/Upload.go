package api_v3

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// 可选：限制文件类型
	allowedExtensions := []string{".jpg", ".jpeg", ".bmp", ".gif", ".tiff", ".png", ".svg"}
	if !isAllowedExtension(file.Filename, allowedExtensions) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid file type",
		})
		return
	}

	// 可选：限制文件大小
	maxFileSize := int64(8 << 20) // 8MB
	if file.Size > maxFileSize {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "File size exceeds the limit",
		})
		return
	}

	// 将文件保存到指定路径
	now := time.Now()
	filename := now.Format("2006-01-02 15:04:05") + "_" + file.Filename
	err = c.SaveUploadedFile(file, "static/uploads/"+filename)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to save file",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"url":     "/static/uploads/" + filename,
		"message": "File uploaded successfully",
	})
}

func isAllowedExtension(filename string, allowedExtensions []string) bool {
	for _, ext := range allowedExtensions {
		if ext == getFileExtension(filename) {
			return true
		}
	}
	return false
}

func getFileExtension(filename string) string {
	return filename[len(filename)-4:]
}
