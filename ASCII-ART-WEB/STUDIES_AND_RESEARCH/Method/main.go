package main

import (
	"fmt"
)

// 1. Define Price as a custom float64 type
type Price float64

// 2. Attach your logic as a method to the Price type
func (p Price) CanCashPr() string {
	// Add 0.5 to prevent Go from rounding 1.15 down to 114
	totalCents := int(p*100 + 0.5)

	remainder := totalCents % 5
	quotiant := totalCents / 5

	if remainder < 3 {
		pr := float64(quotiant*5) / 100
		s := fmt.Sprintf("%.2f", pr)
		return s
	}

	pr := (float64(quotiant*5) + 5) / 100
	s := fmt.Sprintf("%.2f", pr)
	return s
}

func main() {
	// 3. Declare a variable using your custom Price type
	var itemPrice Price = 1.02

	// 4. Call the method directly on the variable using dot notation
	cashPrice := itemPrice.CanCashPr()

	fmt.Printf("Original Price: $%.2f\n", itemPrice)
	fmt.Printf("Rounded Cash Price: $%s\n", cashPrice)
}
