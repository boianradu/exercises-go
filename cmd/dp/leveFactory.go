package dp

type Leve struct {
}

func (p *Leve) MakePill() IPill {
	return &LevePill{
		Pill: Pill{
			name:  "Levetiracetam",
			price: 88.43,
		},
	}
}
