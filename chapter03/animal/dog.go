package animal

type Dog struct {
	a *Animal
	p *Pet
}

func NewDog(animal *Animal, p *Pet) Dog {
	return Dog{a: animal, p: p}
}

func (d Dog) Call() string {
	return "汪汪..."
}

func (d Dog) FavorFood() string {
	return "爱骨头"
}

func (d Dog) GetPetName() string {
	return d.p.GetName()
}
