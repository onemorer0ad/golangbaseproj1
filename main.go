package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/onemorer0ad/coffee"
	"github.com/onemorer0ad/recipes"
)


func main(){
	coffees := recipes.NewCoffee()
	ingredients := coffee.NewIngredients()
	stats := coffee.Stats{
		Orders: 0,
		Revenue: 0,
	}
	scanner := bufio.NewScanner(os.Stdin)
	for {
		if !scanner.Scan() { break }
		parts := strings.Fields(scanner.Text())
		if len(parts) == 0 { continue }

		switch {
		case len(parts) >= 2 && parts[0] == "coffee" && parts[1] == "help":
			fmt.Print(coffee.HelpText)
			return
		case len(parts) >= 2 && parts[0] == "coffee" && parts[1] == "menu":
			coffee.GetMenu(coffees)
		case len(parts) >= 2 && parts[0] == "coffee" && parts[1] == "quit":
			return
		case len(parts) >= 4 && parts[0] == "coffee" && parts[1] == "stock" && parts[2] == "add":
			coffee.StockAdd(ingredients, parts)
		case len(parts) >= 4 && parts[0] == "coffee" && parts[1] == "stock" && parts[2] == "set":
			coffee.StockSet(ingredients, parts)
		case len(parts) >= 3 && parts[0] == "coffee" && parts[1] == "stock" && parts[2] == "get":
			coffee.StockGet(ingredients)
		case len(parts) >= 3 && parts[0] == "coffee" && parts[1] == "brew":
			coffee.Brew(coffees, ingredients, parts[2:], &stats)
		case len(parts) >= 2 && parts[0] == "coffee" && parts[1] == "stats":
			stats.GetStats()

		default:
			fmt.Println("Неизвестная команда")
		}
	}
}