package coffee

const HelpText = `Использование:
  coffee help            — вывести справку и завершиться с кодом 0.
  coffee menu            — вывести доступные напитки и цены.
  coffee stock get       — показать остатки ингредиентов.
  coffee stock add <ingredient> <qty> — прибавить к остатку (qty >= 1).
  coffee stock set <ingredient> <qty> — установить остаток (qty >= 0).
  coffee brew <drink> --pay <amount>   — приготовить напиток.
  coffee stats           — вывести статистику: orders и revenue.
`