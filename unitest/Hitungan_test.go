package unitest

import "testing"

func TestHitunganPerkalian(t *testing.T) {
	result := HitunganPerkalian(3, 4)
	expected := 12

	if result != expected {
		panic("Result is wrong")
	}
}
