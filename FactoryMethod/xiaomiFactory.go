package main

type xiaomiFactory struct {

}

func (this *xiaomiFactory) createMobile(model ModelType) iMobile {
	switch model {
	case Xiaomi10:
		return &xiaomi10{}
	case Xiaomi11:
		return &xiaomi11{}
	}
	return nil
}