// The Models for my Database
package model

type LoginRequest struct {
    Email string 
    Password string 
}

type User struct {
    ID    int
    Name  string
    Email string
    Password string
    Token string
}