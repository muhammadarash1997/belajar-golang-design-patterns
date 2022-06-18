/// AbstractFactory ///

package main

import (
	"fmt"
)

type iSportsApparelFactory interface {
	makeShoe() iShoe
	makeShort() iShort
}

func getSportsApparelFactory(brand string) (iSportsApparelFactory, error) {
	if brand == "brandA" {
		return &brandA{}, nil
	}
	if brand == "brandB" {
		return &brandB{}, nil
	}
	return nil, fmt.Errorf("Wrong brand type passed")
}