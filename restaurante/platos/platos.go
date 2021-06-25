package platos

type Platos struct {
	Id    int64
	Name  string
	Price int
}

type crearPlatos struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
}
