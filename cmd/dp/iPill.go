package dp

type IPill interface {
	setName(name string)
	setPrice(price float32)
	GetName() string
	GetPrice() float32
}

type Pill struct {
	name  string
	price float32
}

func (p *Pill) setName(name string) {
	p.name = name
}

func (p *Pill) GetName() string {
	return p.name
}

func (p *Pill) setPrice(price float32) {
	p.price = price
}
func (p *Pill) GetPrice() float32 {
	return p.price
}
