package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws/session"
	cognitoIDP "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/markbates/pkger"
	"log"
	"mime"
	"os/exec"
	"path/filepath"
	"runtime"
)

type NewPasswordRequest struct {
	NewPassword string `json:"NewPassword" required:"true"`
}

func openBrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}

	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	cognito := cognitoIDP.New(sess)

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowMethods: []string{"GET", "POST"},
		AllowHeaders: []string{"Content-Type"},
		AllowOriginFunc: func(origin string) bool {
			return true
		},
	}))

	api := r.Group("/api")
	{
		api.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
		api.GET("/pools", func(c *gin.Context) {
			opInput := &cognitoIDP.ListUserPoolsInput{
				MaxResults: nil,
			}
			opInput.SetMaxResults(50)
			poolList, err := cognito.ListUserPools(opInput)
			if err == nil {
				c.JSON(200, gin.H{
					"Pools": poolList.UserPools,
				})
			} else {
				c.JSON(400, gin.H{
					"Message": err.Error(),
				})
			}
		})
		api.GET("/pools/:poolId", func(c *gin.Context) {
			opInput := &cognitoIDP.DescribeUserPoolInput{}
			opInput.SetUserPoolId(c.Param("poolId"))
			poolData, err := cognito.DescribeUserPool(opInput)
			if err == nil {
				c.JSON(200, gin.H{
					"Pool": poolData.UserPool,
				})
			} else {
				c.JSON(400, gin.H{
					"Message": err.Error(),
				})
			}
		})
		api.GET("/pools/:poolId/users", func(c *gin.Context) {
			opInput := &cognitoIDP.ListUsersInput{}
			opInput.SetUserPoolId(c.Param("poolId"))
			opInput.SetLimit(60)
			paginationToken := c.Query("after")
			if paginationToken != "" {
				opInput.SetPaginationToken(paginationToken)
			}
			poolUsers, err := cognito.ListUsers(opInput)
			if err == nil {
				c.JSON(200, gin.H{
					"Users":           poolUsers.Users,
					"PaginationToken": poolUsers.PaginationToken,
				})
			} else {
				c.JSON(400, gin.H{
					"Message": err.Error(),
				})
			}
		})
		api.GET("/pools/:poolId/users/:username", func(c *gin.Context) {
			opInput := &cognitoIDP.AdminGetUserInput{}
			opInput.SetUserPoolId(c.Param("poolId"))
			opInput.SetUsername(c.Param("username"))
			user, err := cognito.AdminGetUser(opInput)
			if err == nil {
				c.JSON(200, gin.H{
					"User": user,
				})
			} else {
				c.JSON(400, gin.H{
					"Message": err.Error(),
				})
			}
		})
		api.POST("/pools/:poolId/users/:username/attributes", func(c *gin.Context) {
			var updatedAttributeList []*cognitoIDP.AttributeType
			if err := c.ShouldBind(&updatedAttributeList); err != nil {
				c.JSON(400, gin.H{
					"Message": err.Error(),
				})

				return
			}

			opInput := &cognitoIDP.AdminUpdateUserAttributesInput{}
			opInput.SetUserPoolId(c.Param("poolId"))
			opInput.SetUsername(c.Param("username"))
			opInput.SetUserAttributes(updatedAttributeList)
			_, err := cognito.AdminUpdateUserAttributes(opInput)
			if err == nil {
				c.JSON(204, gin.H{})
			} else {
				c.JSON(400, gin.H{
					"Message": err.Error(),
				})
			}
		})
		api.POST("/pools/:poolId/users/:username/change_password", func(c *gin.Context) {
			var newPasswordRequest NewPasswordRequest
			if err := c.ShouldBind(&newPasswordRequest); err != nil {
				c.JSON(400, gin.H{
					"Message": err.Error(),
				})

				return
			}

			opInput := &cognitoIDP.AdminSetUserPasswordInput{}
			opInput.SetUserPoolId(c.Param("poolId"))
			opInput.SetUsername(c.Param("username"))
			opInput.SetPassword(newPasswordRequest.NewPassword)
			opInput.SetPermanent(true)
			_, err := cognito.AdminSetUserPassword(opInput)
			if err == nil {
				c.JSON(204, gin.H{})
			} else {
				c.JSON(400, gin.H{
					"Message": err.Error(),
				})
			}
		})
		api.POST("/pools/:poolId/users/:username/confirm_signup", func(c *gin.Context) {
			opInput := &cognitoIDP.AdminConfirmSignUpInput{}
			opInput.SetUserPoolId(c.Param("poolId"))
			opInput.SetUsername(c.Param("username"))
			_, err := cognito.AdminConfirmSignUp(opInput)
			if err == nil {
				c.JSON(204, gin.H{})
			} else {
				c.JSON(400, gin.H{
					"Message": err.Error(),
				})
			}
		})
	}
	r.NoRoute(func(c *gin.Context) {
		path := c.Request.RequestURI
		if path == "/" {
			path = "/index.html"
		}
		path = "/webui/dist" + path
		f, err := pkger.Open(path)
		if err != nil {
			f, err = pkger.Open("/webui/dist/index.html")
			if err != nil {
				c.JSON(500, gin.H{
					"Message": err.Error(),
				})
				return
			}
		}
		info, err := f.Stat()
		if err != nil {
			c.JSON(500, gin.H{
				"Message": err.Error(),
			})
			return
		}

		extension := filepath.Ext(path)
		mimeType := mime.TypeByExtension(extension)

		c.DataFromReader(200, info.Size(), mimeType, f, map[string]string{})

		defer f.Close()
	})

	openBrowser("http://127.0.0.1:8819")
	err := r.Run("127.0.0.1:8819")
	if err != nil {
		fmt.Printf(err.Error())
	}
}
