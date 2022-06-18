package main

type samsungFactory struct {

}

func (this *samsungFactory) createMobile(model ModelType) iMobile {
	switch model {
	case SamsungS10:
		return &samsungS10{}
	case SamsungS11:
		return &samsungS11{}
	}
	return nil
}