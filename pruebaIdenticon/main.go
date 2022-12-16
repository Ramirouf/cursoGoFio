package main

import (
    "os"

    "github.com/nullrocks/identicon"
)
func main(){
	// New Generator: Rehuse 
	// New returns a pointer to a Generator with the desired configuration.
	// New is located in generator.go
ig, err := identicon.New(
    "githubaaa", // Namespace
    5,        // Number of blocks (Size)
    3,        // Density
)

if err != nil {
    panic(err) // Invalid Size or Density
}

username := "Ramasdasdasdaasdaasad"      // Text - decides the resulting figure
// Draw returns a pointer to an IdentIcon with a generated figure and a color.

ii, err := ig.Draw(username) // Generate an IdentIcon

if err != nil {
    panic(err) // Text is empty
}

// File writer
img, _ := os.Create("icon.png")
defer img.Close()
// Takes the size in pixels and any io.Writer
ii.Png(300, img) // 300px * 300px
}