package main

func createMobile(brand BrandType) iMobile {
	switch brand {
	case Samsung:
		return &samsung{}
	case Xiaomi:
		return &xiaomi{}
	}
	return nil
}
