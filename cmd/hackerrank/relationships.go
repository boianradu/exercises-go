package hackerrank

import "log"

/*
We are given 10 individuals say, a, b, c, d, e, f, g, h, i, j


Following are relationships to be added:
a <-> b
b <-> d
c <-> f
c <-> i
j <-> e
g <-> j
*/

type Relation struct {
	a string
	b string
}

type Group struct {
	individuals []string
}

type IGroup interface {
	FindRelation(string, string)
	ListMembers()
	AddMember(string)
	AddMembers(string)
}

func (g *Group) FindRelation(a string, b string) bool {
	foundA := false
	foundB := false
	for _, x := range g.individuals {
		if x == a {
			foundA = true
		} else if x == b {
			foundB = true
		}
	}
	if foundA && foundB {
		return true
	}
	return false
}

func (g *Group) AddMember(m string) {
	g.individuals = append(g.individuals, m)
}
func (g *Group) AddMembers(m []string) {
	g.individuals = append(g.individuals, m...)
}

func (g *Group) ListMembers() {
	log.Print("Individuals are: ", g.individuals)
}
