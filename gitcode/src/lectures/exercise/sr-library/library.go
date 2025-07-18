//--Summary:
//  Create a program to manage lending of library books.
//
//--Requirements:
//* The library must have books and members, and must include:
//  - Which books have been checked out
//  - What time the books were checked out
//  - What time the books were returned
//* Perform the following:
//  - Add at least 4 books and at least 3 members to the library
//  - Check out a book
//  - Check in a book
//  - Print out initial library information, and after each change
//* There must only ever be one copy of the library in memory at any time
//
//--Notes:
//* Use the `time` package from the standard library for check in/out times
//* Liberal use of type aliases, structs, and maps will help organize this project

package main

import (
	"time"
)

type Title string
type UserName string

type Book struct {
	Title        Title
	CheckedOutAt time.Time
	CheckInAt    time.Time
}
type Books struct {
	Items []Book
}

type User struct {
	Name  UserName
	Books []Book
}
type Users struct {
	Items []User
}

type library struct {
	books Books
	users Users
}

func CheckOutBook(title Title, user User) {

}

func CheckInBook(title Title, user User) {

}

func RegisterUser() {

}

func main() {

}
