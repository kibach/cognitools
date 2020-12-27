package main

import (
	"fmt"
	"log"
	"mime"
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/kibach/cognitools/pkg/restful"
	"github.com/markbates/pkger"
)

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

func createCorsMiddleware() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowMethods: []string{"GET", "POST"},
		AllowHeaders: []string{"Content-Type"},
		AllowOriginFunc: func(origin string) bool {
			return true
		},
	})
}

func addWebUiRoutes(r *gin.Engine) {
	r.NoRoute(func(c *gin.Context) {
		path := c.Request.RequestURI
		if path == "/" {
			path = "/index.html"
		}
		path = "/web/ui/dist" + path
		f, err := pkger.Open(path)
		if err != nil {
			f, err = pkger.Open("/web/ui/dist/index.html")
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
}

func main() {
	openBrowser("http://127.0.0.1:8819")
	r := restful.CreateBaseRestfulServer(createCorsMiddleware())
	addWebUiRoutes(r)
	err := r.Run("127.0.0.1:8819")
	if err != nil {
		fmt.Printf("%s", err.Error())
	}
}
