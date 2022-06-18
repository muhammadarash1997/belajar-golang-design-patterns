/// Main ///

package main

import "fmt"

func main() {
	brandAFactory, _ := getSportsApparelFactory("brandA")
	brandAShoe := brandAFactory.makeShoe()
	brandAShort := brandAFactory.makeShort()
	
	brandBFactory, _ := getSportsApparelFactory("brandB")
	brandBShoe := brandBFactory.makeShoe()
	brandBShort := brandBFactory.makeShort()

	printShoeDetails(brandBShoe)
	printShortDetails(brandBShort)
	printShoeDetails(brandAShoe)
	printShortDetails(brandAShort)
}

func printShoeDetails(s iShoe) {
	fmt.Printf("Logo: %s", s.getLogo())
	fmt.Println()
	fmt.Printf("Size: %d", s.getSize())
	fmt.Println()
}

func printShortDetails(s iShort) {
	fmt.Printf("Logo: %s", s.getLogo())
	fmt.Println()
	fmt.Printf("Size: %d", s.getSize())
	fmt.Println()
}
