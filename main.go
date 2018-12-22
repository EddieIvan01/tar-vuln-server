package main

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	"github.com/gin-gonic/gin"
)

func index_get(ctx *gin.Context) {
	cmd := exec.Command("ls")
	result, _ := cmd.Output()
	ctx.HTML(200, "index.html", string(result))
}

func index_post(ctx *gin.Context) {
	file, _ := ctx.FormFile("file")
	file_content, _ := file.Open()
	defer file_content.Close()
	filename := file.Filename
	fp, _ := os.Create(filename)
	defer fp.Close()
	_, err := io.Copy(fp, file_content)
	if err != nil {
		log.Println(err.Error())
	}
}

func read_file(ctx *gin.Context) {
	filename, _ := ctx.GetQuery("file")
	log.Println(filename)
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Println(err.Error())
	}
	ctx.String(200, "%s", content)
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("/", index_get)
	router.POST("/", index_post)
	router.GET("/readfile", read_file)
	router.Run(":2333")
}
