package restful

import (
	"github.com/gin-gonic/gin"
)

func jsonOrError(c *gin.Context, data gin.H, err error) {
	if err != nil {
		c.JSON(400, gin.H{
			"Message": err.Error(),
		})

		return
	}

	if len(data) == 0 {
		c.JSON(204, data)
	} else {
		c.JSON(200, data)
	}
}

func mustBindOrJSONError(c *gin.Context, obj interface{}) error {
	if err := c.ShouldBind(obj); err != nil {
		c.JSON(400, gin.H{
			"Message": err.Error(),
		})

		return err
	}

	return nil
}
