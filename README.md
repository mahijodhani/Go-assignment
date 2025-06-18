# Go-assignment
# Instructions and example

## Objective
Create an API in Go using the Gin framework that maintains a list of integers.
- If the sign of the input number matches the sign of the existing list elements, append it.
- If the sign is opposite, remove the input value's quantity from the list using FIFO logic.

## How to Run
1. Install Go: https://go.dev/doc/install
2. In terminal or PowerShell:

go mod init go-app
go get github.com/gin-gonic/gin
go get github.com/sirupsen/logrus
go run main.go

Server starts at: http://localhost:8080

## API Endpoints
# POST /add
To add a number to the list:
Invoke-RestMethod -Uri http://localhost:8080/add -Method Post -Body (@{ number = 10 } | ConvertTo-Json) -ContentType "application/json"

# GET /list
To view the list:
http://localhost:8080/list

## Exaample
Input: 10
Output: [10]

Input: 10
Output: [10, 10]

Input: -8
Output: [2, 10]

[PS C:\Users\HP\OneDrive\Documents\go-app>  Invoke-RestMethod -Uri http://localhost:8080/add -Method Post -Body (@{ number = 10 } | ConvertTo-Json) -ContentType "application/json"
list    
----
{10, 10}
PS C:\Users\HP\OneDrive\Documents\go-app>  Invoke-RestMethod -Uri http://localhost:8080/add -Method Post -Body (@{ number = -8 } | ConvertTo-Json) -ContentType "application/json"
list
----
{2, 10}
]

## Unit Test
To test the getSign() function:
In terminal or powershell:
go test

{It checks:
Positive input → returns 1
Negative input → returns -1
Zero input → returns 0}

## List Behavior Summary
Starts empty: []
If same sign: append to list
If opposite sign: subtract value from leftmost (FIFO)
List is updated and returned after each /add
