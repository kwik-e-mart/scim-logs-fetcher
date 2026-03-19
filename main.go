package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("Hello from scim-logs-fetcher! (OS: %s, Arch: %s)\n", runtime.GOOS, runtime.GOARCH)
}
