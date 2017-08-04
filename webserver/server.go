package webserver

import (
	"html/template"
	"path"

	"github.com/gin-gonic/gin"
)

func isStringInSlice(s string, l []string) bool {
	for _, s2 := range l {
		if s == s2 {
			return true
		}
	}
	return false
}

func makeApiUrl(s string) string {
	return "http://dev-videoso:8001" + path.Join("/", s)
}

const fileDir = "/opt/videoso/webclient/"

var staticExt = []string{".css", ".js"}
var htmlExt = []string{".html"}
var templateExt = []string{".gotmpl"}

func RunServer(addr string) {
	router := gin.Default()

	router.SetFuncMap(template.FuncMap{
		"makeApiUrl": makeApiUrl,
	})
	for _, ext := range templateExt {
		router.LoadHTMLGlob(path.Join(fileDir, "*"+ext))
	}

	router.GET("/", func(c *gin.Context) {
		serveTemplate("index"+templateExt[0], c)
	})
	router.GET("/login", func(c *gin.Context) {
		serveTemplate("login"+templateExt[0], c)
	})
	router.GET("/register", func(c *gin.Context) {
		serveTemplate("login"+templateExt[0], c)
	})
	router.GET("/forgot-password", func(c *gin.Context) {
		serveTemplate("login"+templateExt[0], c)
	})
	router.GET("/reset-password", func(c *gin.Context) {
		serveTemplate("reset-password"+templateExt[0], c)
	})

	router.GET("/f/:filename", func(c *gin.Context) {
		filename := c.Param("filename")
		if filename == "" {
			c.JSON(400, gin.H{
				"error": "filename is required",
			})
			return
		}
		filenameBase := path.Base(filename)
		if filename != filenameBase {
			c.JSON(400, gin.H{
				"error": "filename must be flat",
			})
			return
		}
		filenameExt := path.Ext(filename)
		filenameRoot := filename[0 : len(filename)-len(filenameExt)]

		if isStringInSlice(filenameExt, htmlExt) {
			serveTemplate(filenameRoot+templateExt[0], c)
		} else if isStringInSlice(filenameExt, staticExt) {
			serveStaticFile(filename, c)
		}
	})

	router.Run(addr)
}

func serveTemplate(filename string, c *gin.Context) {
	c.HTML(200, filename, nil)
}

func serveStaticFile(filename string, c *gin.Context) {
	c.Status(200)
	c.File(path.Join(fileDir, filename))
}
