package main

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

type book struct {
    ID       string `json:"id"`
    Title    string `json:"title"`
    Author   string `json:"author"`
    Quantity int    `json:"quantity"`
}

var books = []book{
    {
        ID:       "1",
        Title:    "In Search of Lost Time",
        Author:   "Marcel Proust",
        Quantity: 2,
    },
    {
        ID:       "2",
        Title:    "To Kill a Mockingbird",
        Author:   "Harper Lee",
        Quantity: 3,
    },
    {
        ID:       "3",
        Title:    "1984",
        Author:   "George Orwell",
        Quantity: 1,
    },
}

func getBooks(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, books)
}

func main() {
    router := gin.Default()

    router.GET("/books", getBooks)

    router.Run("localhost:8080")
}