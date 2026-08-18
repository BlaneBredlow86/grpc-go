package main

import (
	"fmt"
	"github.com/BlaneBredlow86/grpc-go/internal/grpcutil"
)

func main() {
	// Example usage of the fix
	rawPath := "//MyService/MyMethod"
	canonical := grpcutil.CanonicalMethodPath(rawPath)
	fmt.Printf("Original: %s, Canonical: %s\n", rawPath, canonical)
}