# Fizz Buzz

## Testing

Clone the repository.

Navigate to the project root.

To run the test suite:  
`go test`

To run fizz buzz:  
`go run fizz_buzz.go`

To generate an HTML coverage report:  
```
go test -cover -coverprofile=c.out
go tool cover -html=c.out -o coverage.html
```
