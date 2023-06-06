# Uniques

![Go](https://img.shields.io/badge/GO-00ADD8?logo=go&logoColor=white&style=for-the-badge)

## Introduction

Uniques is a library that facilitate you to work with unique ids in different standards.

## Getting Started

### Prerequisites

- **[Go](https://go.dev/)**

### Get the Uniques

Simply run the following command to get the uniques package.

```bash
go get github.com/rashintha/uniques
```

### Basic Usage

```go
package main

import "github.com/rashintha/uniques/standards/gs1"

func main() {
	upc, err = gs1.EPCtoUPC("3034257BF400B78000007AE3")
    if err != nil {
        fmt.Println(err)
    }

    fmt.Printf("UPC: %v\n", upc)
}

```

## License

Uniques is released under Apache-2.0 license. Please read [LICENSE](LICENSE) for more information.

## Contributing

I always appreciate help from the community to develop. Please send me an email to mail@rashintha.com to get started!