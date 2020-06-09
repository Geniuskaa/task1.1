package main

func main() {
	purchase := 3333_33
	purchaseInRubles := purchase / 100
	percent := 1
	limit := 100
	bonus := purchaseInRubles / 100 * percent
	if bonus > limit {
		bonus = limit
	}
	println(bonus)
}
