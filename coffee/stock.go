package coffee

import (
	"fmt"
	"strconv"

	"github.com/onemorer0ad/recipes"
)


func StockGet(ingredients map[string]int){
	for i, v := range ingredients {
		fmt.Println(i,"=",v)
	}
}
func StockAdd(ingredients map[string]int, parts []string){
	item := parts[3]
	count, err := strconv.Atoi(parts[4])
	if err != nil {
		fmt.Println("некорректные параметры")
		return
	}
	if _,ok := ingredients[item]; ok && count >= 1 {
		ingredients[item] += count
		fmt.Println("ok")
	} else {
		fmt.Println("некорректные параметры")
	}
	
}
func StockSet(ingredients map[string]int, parts []string){
	item := parts[3]
	count, err := strconv.Atoi(parts[4])
	if err != nil {
		fmt.Println("некорректные параметры")
		return
	}
	
	if _,ok := ingredients[item]; ok && count >= 0 {
		ingredients[item] = count
		fmt.Println("ok")
	} else {
		fmt.Println("некорректные параметры")
	}
}

func Brew(coffeeRecipes []recipes.Coffee, ingredients map[string]int, parts []string, stats *Stats){
	
	coffeeName := parts[0]
	
	var found bool
	var recipe recipes.Coffee
	for i, v := range coffeeRecipes {
		if coffeeRecipes[i].Name == coffeeName {
			found = true
			recipe = v
			break
		}
	}
	if !found{
		fmt.Println("напиток не найден")
		return
	}


	// command := parts[1]
	amount, err := strconv.Atoi(parts[2])
	if err != nil {
		fmt.Println("некорректные параметры")
		return
	}
	
		// напиток может быть приготовлен если рецепт известен и внесен платеж >= price и на складе хватает всех ингридиентов
	beans := ingredients["beans"]
	water := ingredients["water"]
	milk := ingredients["milk"]
	sugar := ingredients["sugar"]

	recipeBeans := recipe.Ingredients["beans"]
	recipeWater := recipe.Ingredients["water"]
	recipeMilk := recipe.Ingredients["milk"]
	recipeSugar := recipe.Ingredients["sugar"]

	if beans >= recipeBeans && water >= recipeWater && milk >= recipeMilk && sugar >= recipeSugar {
		if amount >= recipe.Price {
			for ing, needed := range recipe.Ingredients {
				ingredients[ing] -= needed
			}
			for _, v := range recipe.CookingSteps {
				fmt.Println(v)
			}
			stats.AddOrder(amount)
		} else {
			fmt.Println("недостаточно денег")
		}

	} else {
		fmt.Println("Недостаточно ингридиентов")
	}

}
