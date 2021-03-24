package entity

type Dish struct {
	Id         int     `db:"id"`
	CategoryID *int    `db:"id_categoria"`
	Name       string  `db:"nome"`
	Price      float64 `db:"preco"`
	CookTime   int     `db:"tempo_de_preparo"`
}

func (d *Dish) IsEmpty() bool {
	return d == nil || (d.Id == 0 && d.Name == "")
}
