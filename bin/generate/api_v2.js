// #!/bin/bash
const fs = require('fs')

const generate = (Model, description) => {
	const filepath = `${__dirname}/../../api/v2/${Model}.go`
	const tmpl = `// ${description}
package api_v2

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"goserver/libs/gorm"
	"goserver/middlewares"
	models "goserver/models/v2"
)

type ${Model}List []models.${Model}

func GetAll${Model}Api(c *gin.Context) {
	middlewares.GetAll(c, &${Model}List{})
}

func Get${Model}Api(c *gin.Context) {
	middlewares.GetOne(c, &models.${Model}{}, &models.${Model}{})
}

func Get${Model}ListApi(c *gin.Context) {
	middlewares.GetList(c, &models.${Model}{}, &${Model}List{})
}

func Create${Model}Api(c *gin.Context) {
	middlewares.CreateOne(c, &models.${Model}{})
}

func Create${Model}ListApi(c *gin.Context) {
	middlewares.CreateList(c, &${Model}List{})
}

func Update${Model}Api(c *gin.Context) {
	middlewares.UpdateOne(c, &models.${Model}{})
}

func Update${Model}ListApi(c *gin.Context) {
	var params ${Model}List
	if middlewares.BindJSON(c, &params) == nil {
		for _, v := range params {
			gorm.Updates(&v)
		}
		c.JSON(http.StatusOK, gin.H{"message": "更新成功！"})
	}
}

func Delete${Model}Api(c *gin.Context) {
	middlewares.DeleteOne(c, &models.${Model}{})
}

func Delete${Model}ListApi(c *gin.Context) {
	middlewares.DeleteList(c, &models.${Model}{})
}
`
	fs.writeFileSync(filepath, tmpl)
}

const [Model, description]= process.argv.splice(2);
console.log(Model, description)
generate(Model, description)


