package restful

import (
	"github.com/gin-gonic/gin"
	"github.com/kibach/cognitools/pkg/cognito_client"
)

func CreateBaseRestfulServer() *gin.Engine {
	client := cognito_client.NewClient()

	r := gin.Default()

	api := r.Group("/api")
	{
		api.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
		api.GET("/pools", func(c *gin.Context) {
			poolList, err := client.ListPools()
			jsonOrError(c, gin.H{
				"Pools": poolList,
			}, err)
		})
		api.GET("/pools/:poolId", func(c *gin.Context) {
			poolInfo, err := client.GetPool(c.Param("poolId")).GetInfo()
			jsonOrError(c, gin.H{
				"Pool": poolInfo,
			}, err)
		})
		api.GET("/pools/:poolId/users", func(c *gin.Context) {
			paginationToken := c.Query("after")
			poolUsers, newPaginationToken, err := client.
				GetPool(c.Param("poolId")).
				GetUsers(paginationToken)

			jsonOrError(c, gin.H{
				"Users":           poolUsers,
				"PaginationToken": newPaginationToken,
			}, err)
		})
		api.POST("/pools/:poolId/users", func(c *gin.Context) {
			var createUserRequest *cognito_client.CreateUserRequest
			if err := mustBindOrJSONError(c, &createUserRequest); err != nil {
				return
			}

			createdUser, err := client.
				GetPool(c.Param("poolId")).
				CreateUser(createUserRequest)

			jsonOrError(c, gin.H{"User": createdUser}, err)
		})
		api.GET("/pools/:poolId/users/:username", func(c *gin.Context) {
			user, mfaPreferences, err := client.
				GetPool(c.Param("poolId")).
				GetUser(c.Param("username")).
				GetInfo()

			jsonOrError(c, gin.H{
				"User":           user,
				"MFAPreferences": mfaPreferences,
			}, err)
		})
		api.POST("/pools/:poolId/users/:username/attributes", func(c *gin.Context) {
			var updatedAttributeList cognito_client.UserAttributeList
			if err := mustBindOrJSONError(c, &updatedAttributeList); err != nil {
				return
			}

			err := client.
				GetPool(c.Param("poolId")).
				GetUser(c.Param("username")).
				UpdateAttributes(updatedAttributeList)

			jsonOrError(c, gin.H{}, err)
		})
		api.POST("/pools/:poolId/users/:username/change_password", func(c *gin.Context) {
			var newPasswordRequest cognito_client.NewPasswordRequest
			if err := mustBindOrJSONError(c, &newPasswordRequest); err != nil {
				return
			}

			err := client.
				GetPool(c.Param("poolId")).
				GetUser(c.Param("username")).
				ChangePassword(newPasswordRequest.NewPassword)

			jsonOrError(c, gin.H{}, err)
		})
		api.POST("/pools/:poolId/users/:username/confirm_signup", func(c *gin.Context) {
			err := client.
				GetPool(c.Param("poolId")).
				GetUser(c.Param("username")).
				ConfirmSignup()

			jsonOrError(c, gin.H{}, err)
		})
	}

	return r
}