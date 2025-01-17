package main

import (
    "net/http"
    "errors"

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

func bookById(c *gin.Context) {
    id := c.Param("id")
    book, err := getBookById(id)
    if err != nil {
        c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
        return
    }
    c.IndentedJSON(http.StatusOK, book)
}

func checkoutBook(c *gin.Context) {
    id, ok := c.GetQuery("id")
    if !ok {
        c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "ID parameter is required"})
        return
    }
    book, err := getBookById(id)
    if err != nil {
        c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
        return
    }
    if book.Quantity <= 0 {
        c.IndentedJSON(http.StatusConflict, gin.H{"message": "Book out of stock"})
        return
    }
    book.Quantity -= 1
    c.IndentedJSON(http.StatusOK, book)
}

func  returnBook(c *gin.Context){
	id, ok := c.GetQuery("id")
    if !ok {
        c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "ID parameter is required"})
        return
    }
    book, err := getBookById(id)
    if err != nil {
        c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
        return
    }
    if book.Quantity < 0 {
        c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Book quantity cannot be negative"})
        return
    }
    book.Quantity += 1
    c.IndentedJSON(http.StatusOK, book)
}

func getBookById(id string) (*book, error) {
    for i, b := range books {
        if b.ID == id {
            return &books[i], nil
        }
    }
    return nil, errors.New("book not found")
}

func createBook(c *gin.Context) {
    var newBook book
    if err := c.BindJSON(&newBook); err != nil {
        c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
        return
    }

    books = append(books, newBook)
    c.IndentedJSON(http.StatusCreated, newBook)
}

func main() {
    router := gin.Default()

    router.GET("/books", getBooks)
    router.POST("/books", createBook)
    router.PATCH("/checkout", checkoutBook)
    router.GET("/books/:id", bookById)

    router.Run("localhost:8080")
}