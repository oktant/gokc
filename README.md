# Go Keycloak Client

A Go client library for interacting with Keycloak, an open-source identity and access management solution.

## Features

- Easy authentication and token management
- User and group management
- Client and role administration
- Realm configuration
- Support for Keycloak REST API endpoints

## Installation

```bash
go get github.com/oktant/gokc
```

## Usage

```go
package main

import (
    "context"
    "fmt"
    "log"

    "github.com/yourusername/gokc"
)

func main() {
    // Initialize Keycloak client
    client := gokc.NewClient("http://localhost:8080/auth", "master", "admin", "admin")

    // Get token
    token, err := client.Login(context.Background())
    if err != nil {
        log.Fatalf("Failed to login: %v", err)
    }
    fmt.Printf("Token: %s\n", token.AccessToken)

    // Example: Get realm info
    realm, err := client.GetRealm(context.Background(), "master")
    if err != nil {
        log.Fatalf("Failed to get realm: %v", err)
    }
    fmt.Printf("Realm: %s\n", realm.Realm)
}
```

## API Reference

For detailed API documentation, please refer to [godoc](https://pkg.go.dev/github.com/yourusername/gokc).

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.