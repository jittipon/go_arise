package thaiid

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ValidateHandler(ctx *gin.Context) {
	var thaiId ThaiId

	if err := ctx.Bind(&thaiId); err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
		return
	}

	if err := ValidateThaiId(thaiId.Id); err != nil {
		ctx.JSON(http.StatusOK, map[string]string{
			"valid": "false",
		})
		return
	}

	ctx.JSON(http.StatusOK, map[string]string{
		"valid": "true",
	})

}
