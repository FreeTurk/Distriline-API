package product

import (
	"distriline/models"

	"github.com/gin-gonic/gin"
)

func AddProductToBusiness(c *gin.Context) {
	type ExpectedReq struct {
		Product models.Product `json:"product"`
	}
}
