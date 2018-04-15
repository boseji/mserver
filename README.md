# mserver
Golang Managed server wrapper for http.Server to handle process SIGINT, SIGKILL, Ctrl+C and provide graceful shutdown of server

## Example

```go
package main

import (
    "fmt"
    "net/http"
    "time"

    "github.com/boseji/mserver"
)

func home(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hello World!")
}

func main() {
    http.HandleFunc("/", home)
    // Configure Server on Localhost:8080 and set timeout for 10seconds before force termination
    server := mserver.NewMserver(":8080", 10*time.Second)

    // Start the Server and wait till we receive some events
    server.GracefulStop(true)

    fmt.Println(" Server is now stopped")
}
```