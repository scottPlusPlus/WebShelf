package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/braintree/manners"
	"github.com/gin-gonic/gin"
)

const _key = "id"
const _table = "route"

var Data TableKeyValueRepo
var Tables []string

func main() {
	var err error

	port := getPort()
	fmt.Println("Serving on port " + port)

	router := gin.New()

	addRoutes(router)
	router.LoadHTMLGlob("templates/*.tmpl.html")

	Data = NewTableKeyValueRepo()
	Tables = make([]string, 0)

	err = manners.ListenAndServe(":"+port, router)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

func getPort() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		portPtr := flag.Int("port", 3001, "sets the port")
		flag.Parse()
		port = fmt.Sprintf("%v", *portPtr)
	}
	return port
}

func addRoutes(router *gin.Engine) {
	t := "/:" + _table
	k := "/:" + _key
	router.GET("/", indexHandler)

	router.GET(t+k, getTableKeyHandler)
	router.GET(t, getTableHandler)

	router.POST(t+k, postTableKeyHandler)
	router.POST(t, postTableHandler)

	router.PUT(t+k, postTableKeyHandler)
	router.PUT(t, postTableHandler)
}
