# WebShelf
A quick HTTP webserver for stashing and retrieving some data

Implemented via Go (Golang) and the <a href="https://github.com/gin-gonic/gin">Gin web framework</a>

Just make a PUT request to /<tableName>/<itemKey> to store a string, then make a GET request to the same to get it back.
