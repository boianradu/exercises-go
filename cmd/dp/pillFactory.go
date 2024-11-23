package dp

import "fmt"

type IPillFactory interface {
	MakePill() IPill
}

func GetPillFactory(brand string) (IPillFactory, error) {
	if brand == "Leve" {
		return &Leve{}, nil
	}
	return nil, fmt.Errorf("Wrong brand type passed")
}
