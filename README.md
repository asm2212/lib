# Bookstore API
 Bookstore API built with Go and the Gin framework. The API allows you to manage a collection of books, including viewing all books, creating a new book, checking out a book, returning a book, and retrieving a book by its ID.

## Endpoints

### Get All Books

- **URL**: `/books`
- **Method**: `GET`
- **Description**: Retrieves a list of all books in the bookstore.

#### Example Request

```sh
curl -X GET "http://localhost:8080/books"
