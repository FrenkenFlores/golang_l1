package l1_1

type Human struct {
	Name string
}

type Action struct {
	Human
}

func (a Action) GetName() string {
	return a.Name
}
