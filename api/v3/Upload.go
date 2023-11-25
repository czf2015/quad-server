package api_v3

import (
	"fmt"
	"goserver/libs/e"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func UploadFile(c *gin.Context) {
	code := e.SUCCESS
	file, err := c.FormFile("file")
	if err != nil {
		code = e.ERROR_FORM_FILE
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    code,
			"message": e.GetMsg(code),
			"err":     err.Error(),
		})
		return
	}

	// 可选：限制文件类型
	allowedExtensions := []string{".jpg", ".jpeg", ".bmp", ".gif", ".tiff", ".png", ".svg"}
	if !isAllowedExtension(file.Filename, allowedExtensions) {
		code = e.ERROR_FILE_TYPE
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    code,
			"message": e.GetMsg(code),
			"err":     "Invalid file type",
		})
		return
	}

	// 可选：限制文件大小
	maxFileSize := int64(8 << 20) // 8MB
	if file.Size > maxFileSize {
		code = e.ERROR_FILE_SIZE
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    code,
			"message": e.GetMsg(code),
			"err":     "File size exceeds the limit",
		})
		return
	}

	// 将文件保存到指定路径
	now := time.Now()
	filename := now.Format("2006-01-02 15:04:05") + "_" + file.Filename
	filename = strings.Replace(filename, " ", "_", -1)
	filename = strings.Replace(filename, ":", "-", -1)
	fmt.Println(filename)
	err = c.SaveUploadedFile(file, "static/uploads/"+filename)
	if err != nil {
		code = e.ERROR_FILE_SAVE
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    code,
			"message": e.GetMsg(code),
			"err":     "Failed to save file",
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
