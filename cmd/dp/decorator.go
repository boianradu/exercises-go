package dp

type ICarr interface {
	GetPrice() int
}

type Skoda struct {
}

func (s *Skoda) GetPrice() int {
	return 15
}

type ChoseColor struct {
	car ICarr
}

func (cc *ChoseColor) GetPrice() int {
	return cc.car.GetPrice() + 200
}

type ChoseMirrors struct {
	Car ICarr
}

func (cm *ChoseMirrors) GetPrice() int {
	return cm.Car.GetPrice() + 175
}
