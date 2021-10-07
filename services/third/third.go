package third

import (
	"net/http"
	
	"github.com/gin-gonic/gin"
)

func GetThirdService(c *gin.Context) {
	response, err := http.Get("http://p1.music.126.net/Cd4x1A2MLpkrv-knaghxmg==/109951166491544670.jpg?imageView&quality=89")
	if err != nil || response.StatusCode != http.StatusOK {
		c.Status(http.StatusServiceUnavailable)
		return
	}

	reader := response.Body
	contentLength := response.ContentLength
	contentType := response.Header.Get("Content-Type")

	extraHeaders := map[string]string{
		"Content-Disposition": `attachment; filename="gopher.png"`,
	}

	c.DataFromReader(http.StatusOK, contentLength, contentType, reader, extraHeaders)
}