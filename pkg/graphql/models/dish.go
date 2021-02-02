package models

func NewDish() []*Dish {
	result := make([]*Dish, 0)
	result = append(result, &Dish{
		ID:       0,
		Category: nil,
		Name:     "Peito de Frango",
		Price:    15.50,
		CookTime: 15,
	})

	return result
}
