# ShelfWebServer
A quick webserver you can spin up to dump and retreive some data.

Just make a PUT request to /<tableName>/<itemKey> to store a string, then make a GET request to the same to get it back.

Implemented via Go (Golang) and the <a href="https://github.com/gin-gonic/gin">Gin web framework</a>
