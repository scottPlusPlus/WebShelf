package main

import (
	"bytes"
	"net/http"

	"github.com/gin-gonic/gin"
)

type viewModel struct {
	Items []string
}

type TableKeyValue struct {
	Table string
	Key   string
	Value string
}

const _admin = "admin"

func getTableHandler(c *gin.Context) {
	table := c.Param(_table)
	if table == _admin {
		indexHandler(c)
		return
	} else {
		handlerGetAll(c)
		return
	}
	handler404(c)
}

func getTableKeyHandler(c *gin.Context) {
	table := c.Param(_table)
	if table == _admin {
		key := c.Param(_key)
		if key == "all" {
			c.JSON(200, gin.H{"data": Data})
			return
		}

	} else {
		handlerGet(c)
		return
	}
	handler404(c)
}

func postTableKeyHandler(c *gin.Context) {
	table := c.Param(_table)
	if table == _admin {
		handler404(c)
		return
	} else {
		handlerSet(c)
		return
	}
	handler404(c)
}

func postTableHandler(c *gin.Context) {
	table := c.Param(_table)
	if table == _admin {
		handler404(c)
		return
	} else {
		handlerAppend(c)
		return
	}
	handler404(c)
}

func clearHandler(c *gin.Context) {
	Data = NewTableKeyValueRepo()
	Tables = make([]string, 0)
	c.JSON(200, "success")
}

func indexHandler(c *gin.Context) {
	vm := viewModel{
		Items: Data.tables(),
	}
	c.HTML(http.StatusOK, "home.tmpl.html", vm)
}

func handlerGet(c *gin.Context) {
	id := c.Param(_key)
	table := c.Param(_table)

	item, err := Data.get(table, id)
	if err != nil {
		returnError(c, err)
		return
	}
	c.String(200, item)
}

func handlerGetAll(c *gin.Context) {
	table := c.Param(_table)

	items, err := Data.getAll(table)
	if err != nil {
		returnError(c, err)
		return
	}
	c.JSON(200, items)
}

func handlerSet(c *gin.Context) {
	table := c.Param(_table)
	id := c.Param(_key)
	var b bytes.Buffer
	_, err := b.ReadFrom(c.Request.Body)
	if err != nil {
		returnError(c, err)
		return
	}
	str := b.String()
	Data.set(table, id, str)
	c.JSON(200, "success")
}

func handlerAppend(c *gin.Context) {
	table := c.Param(_table)
	var b bytes.Buffer
	_, err := b.ReadFrom(c.Request.Body)
	if err != nil {
		returnError(c, err)
		return
	}
	str := b.String()
	Data.append(table, str)
	c.JSON(200, "success")
}

func handler404(c *gin.Context) {
	c.JSON(404, "404, yo")
}

func returnError(c *gin.Context, err error) {
	c.JSON(500, gin.H{"error": err.Error()})
}
