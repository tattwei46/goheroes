package hero

type Hero struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Heroes struct {
	Heroes []Hero
}

func New() *Heroes {
	return &Heroes{
		Heroes: []Hero{},
	}
}

func (r *Heroes) Add(hero Hero) {
	r.Heroes = append(r.Heroes, hero)
}

func (r *Heroes) GetAll() []Hero {
	return r.Heroes
}
