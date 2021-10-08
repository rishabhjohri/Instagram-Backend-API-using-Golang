// author : Rishabh Johri
// description : Test case 0 (init file)
package main

import (
	"fmt"
	"github.com/rishabhjohri/Instagram-Backend-API-using-Golang/instagram"
	"os"
)

func main() {

	client := instagram.NewClient(nil)
	client.ClientID = "8f2c0ad697ea4094beb2b1753b7cde9c"

	media, next, err := client.Media.Popular()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
	}
	for _, m := range media {
		fmt.Printf("ID: %v, Type: %v\n", m.ID, m.Type)
	}
	if next.NextURL != "" {
		fmt.Println("Next URL", next.NextURL)
	}
}

//Rishabh Johri 19BDS0021
