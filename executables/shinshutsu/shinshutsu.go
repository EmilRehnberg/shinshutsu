package main

import "github.com/emilrehnberg/shinshutsu/brewci"

func main() {
	brewci.Execute(brewci.BuildTeasMap())
}
