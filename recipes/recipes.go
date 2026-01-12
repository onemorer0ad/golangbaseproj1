package recipes

type Coffee struct{
	Name string
	Price int
	Ingredients map[string]int
	CookingSteps []string
}



func NewCoffee() []Coffee {
	espresso := Coffee{
		Name: "espresso",
		Price: 150,
		Ingredients: map[string]int{
			"beans":8,
			"water":30,
		},
		CookingSteps: []string{"grind","tamp","brew 25s"},
	}

	americano := Coffee{
		Name: "americano",
		Price: 180,
		Ingredients: map[string]int{
			"beans":8,
			"water":120,
		},
		CookingSteps: []string{"grind","brew 25s","add water"},
	}

	latte := Coffee{
		Name: "latte",
		Price: 220,
		Ingredients: map[string]int{
			"beans":8,
			"water":30,
			"milk": 150,
		},
		CookingSteps: []string{"grind","brew 25s","steam milk","mix"},
	}

	return []Coffee{espresso,americano,latte}
}