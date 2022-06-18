package main

type iMobileFactory interface {
	createMobile(model ModelType) iMobile
}