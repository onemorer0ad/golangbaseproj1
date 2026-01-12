package coffee

import (
	"fmt"

	"github.com/onemorer0ad/recipes"
)

func GetMenu(coffee []recipes.Cofee) {
	for _, v := range coffee {
		fmt.Println(v.Name, v.Price)
	}
}