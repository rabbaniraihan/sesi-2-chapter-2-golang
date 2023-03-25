package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Book struct {
	Id     int
	Title  string
	Author string
	Desc   string
}

var mapBooks = make(map[int]Book, 0)
var counter int

func main() {
	g := gin.Default()

	g.GET("/book", getAllBook)
	g.POST("/book", addBook)
	g.DELETE("/book/:id", deleteBook)
	g.GET("/book/:id", getBookById)
	g.PUT("/book/:id", updateBook)

	g.Run(":8080")
}

func getAllBook(ctx *gin.Context) {
	books := make([]Book, 0)

	for _, v := range mapBooks {
		books = append(books, v)
	}

	ctx.JSON(http.StatusOK, books)
}

func addBook(ctx *gin.Context) {
	var newBook Book

	err := ctx.ShouldBindJSON(&newBook)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	newBook.Id = counter
	mapBooks[counter] = newBook
	counter++

	ctx.JSON(http.StatusOK, newBook)
}

func deleteBook(ctx *gin.Context) {
	//Ambil id dari param
	stringId := ctx.Param("id")

	//Convert string -> int
	id, err := strconv.Atoi(stringId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	//Cek apakah id yang dicari ada atau ngga
	v, found := mapBooks[id]
	if !found {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": "not found",
		})
	}

	delete(mapBooks, id)
	ctx.JSON(http.StatusOK, v)
}

func getBookById(ctx *gin.Context) {
	stringId := ctx.Param("id")

	id, err := strconv.Atoi(stringId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	v, found := mapBooks[id]
	if !found {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": "not found",
		})
	}

	ctx.JSON(http.StatusOK, v)
}

func updateBook(ctx *gin.Context) {
	stringId := ctx.Param("id")

	id, err := strconv.Atoi(stringId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	v, found := mapBooks[id]
	if !found {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": "not found",
		})
		return
	}

	var updatedBook Book

	err = ctx.ShouldBindJSON(&updatedBook)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	updatedBook.Id = v.Id

	mapBooks[id] = updatedBook

	ctx.JSON(http.StatusOK, updatedBook)
}
