package main

func main() {
	var mobileFactory1 iMobileFactory = &samsungFactory{}
	var mobile1 iMobile = mobileFactory1.createMobile(SamsungS10)
	var mobile2 iMobile = mobileFactory1.createMobile(SamsungS11)
	mobile1.getMobile()
	mobile2.getMobile()

	var mobileFactory2 iMobileFactory = &xiaomiFactory{}
	var mobile3 iMobile = mobileFactory2.createMobile(Xiaomi10)
	var mobile4 iMobile = mobileFactory2.createMobile(Xiaomi11)
	mobile3.getMobile()
	mobile4.getMobile()
}