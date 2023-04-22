package main

import (
	"fmt"
	generatePkg "github.com/skbhati199/generateuniqueid/pkg"
)

func main() {
	generateId, err := generatePkg.NewGeneratePkg()
	if err != nil {
		_ = fmt.Errorf("fail to generate id %v", err)
	}

	fmt.Printf("Transaction Id : %v", generateId)
}
