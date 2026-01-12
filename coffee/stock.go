package coffee

import (
	"fmt"
	"strconv"

	"github.com/onemorer0ad/recipes"
)


func StockGet(ingridients map[string]int){
	for i, v := range ingridients {
		fmt.Println(i,"=",v)
	}
}
func StockAdd(ingridients map[string]int, parts []string){
	item := parts[3]
	count, err := strconv.Atoi(parts[4])
	if err != nil {
		fmt.Println("error when converting to integer")
		return
	}
	if _,ok := ingridients[item]; ok && count >= 1 {
		ingridients[item] += count
		fmt.Println("ok")
	} else {
		fmt.Println("некорректные параметры")
	}
	
}
func StockSet(ingridients map[string]int, parts []string){
	item := parts[3]
	count, err := strconv.Atoi(parts[4])
	if err != nil {
		fmt.Println("некорректные параметры")
		return
	}
	
	if _,ok := ingridients[item]; ok && count >= 0 {
		ingridients[item] = count
		fmt.Println("ok")
	} else {
		fmt.Println("некорректные параметры")
	}
}

func Brew(recipes []recipes.Cofee, ingridients map[string]int, parts []string){
	fmt.Println(parts)
	coffeeName := parts[0]
	// command := parts[1]
	amount, err := strconv.Atoi(parts[2])
	if err != nil {
		fmt.Println("некорректные параметры")
		return
	}
	
	for _, value := range recipes {
		// напиток может быть приготовлен если рецепт известен и внесен платеж >= price и на складе хватает всех ингридиентов
		if coffeeName == value.Name {
			beans := ingridients["beans"]
			water := ingridients["water"]
			milk := ingridients["milk"]
			sugar := ingridients["sugar"]

			recipeBeans := value.Ingredients["beans"]
			recipeWater := value.Ingredients["water"]
			recipeMilk := value.Ingredients["milk"]
			recipeSugar := value.Ingredients["sugar"]

			if beans >= recipeBeans && water >= recipeWater && milk >= recipeMilk && sugar >= recipeSugar && amount >= value.Price {
				for _, v := range value.CookingSteps {
					fmt.Println(v)
				}
			} else {
				fmt.Println("Недостаточно ингридиентов либо денег")
			}
		}
	}
}
