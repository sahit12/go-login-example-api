# Go Login Example API
This is an example API to showcase a simple register/login in Golang.

## Directory Structure
    <GO-GITHUB> - Package name
        |
        ├───config
        │    └───db
        ├───model
        └───view

## How to run
    cd <package-name>
    go get go.mongodb.org/mongo-driver/mongo
    go get -u github.com/gorilla/mux
    
    // Create the above directory structure
    
    go mod init <package-name>
    go run roll.go
    
## Output
    curl -X POST -H "Content-Type: application/json" -d "{\"username\":\"abc12\", \"firstname\":\"abc\", \"lastname\":\"cde\", \"email\":\"abc@example.com\", \"password\":\"testing123\"}" "localhost:8080/register"
