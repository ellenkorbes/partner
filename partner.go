package partner

// 1.1.0-sliced
// 1.2.0-bruschetta
// 2.0.0-croissant

// 1.0.0-still
// 1.1.0-sparkling
// 1.2.0-hot

// 1.1.0-percolated
// 1.2.0-cappuccino
// 1.3.0-espresso
// 1.4.0-soymilk
// 1.5.0-decaf
// 1.6.0-iced

import (
	"github.com/ellenkorbes/bread"
	"github.com/ellenkorbes/coffee"
	"github.com/ellenkorbes/water"
)

func try(what, served string, dislikes []string) string {
	for _, ew := range dislikes {
		if served == ew {
			return "This " + served + " " + what + " is terrible!"
		}
	}
	return "This " + served + " " + what + " is great."
}

// Water drinks water.
func Water() string {
	dislikes := []string{"hot"}
	return try("water", water.Glass(), dislikes)
}

// Coffee drinks coffee.
func Coffee() string {
	dislikes := []string{"percolated", "cappuccino", "iced"}
	return try("coffee", coffee.Cup(), dislikes)
}

// Bread eats bread.
func Bread() string {
	dislikes := []string{"croissant"}
	return try("bread", bread.Slice(), dislikes)
}
