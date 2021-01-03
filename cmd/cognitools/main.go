package main

import (
	"fmt"
	"log"
	"mime"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/kibach/cognitools/pkg/restful"
	"github.com/markbates/pkger"
	"github.com/urfave/cli/v2"
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

func createCORSMiddleware() gin.HandlerFunc {
	return cors.Default()
}

func addWebUIRoutes(r *gin.Engine) {
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
	app := &cli.App{
		Name:  "cognitools",
		Usage: "app that adds missing features to native AWS Cognito UI",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "listen",
				Value: "127.0.0.1:36836",
				Usage: "address to bind a web server",
			},
			&cli.BoolFlag{
				Name:  "no-open-browser",
				Value: false,
				Usage: "don't open browser with the binded address after server startup",
			},
			&cli.BoolFlag{
				Name:  "no-ui",
				Value: false,
				Usage: "don't serve UI (API-only mode)",
			},
		},
		Action: func(c *cli.Context) error {
			listenOn := c.String("listen")

			if !c.Bool("no-open-browser") {
				openBrowser("http://" + listenOn)
			}

			r := restful.CreateBaseRestfulServer(createCORSMiddleware())

			if !c.Bool("no-ui") {
				addWebUIRoutes(r)
			}

			return r.Run(listenOn)
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
