package unitest

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTablePerkalian(t *testing.T) {
	testiIter := []struct {
		x        int
		y        int
		expected int
	}{
		{4, 3, 12},
		{2, 4, 8},
		{3, 6, 18},
	}

	for _, test := range testiIter {
		res := HitunganPerkalian(test.x, test.y)
		if res != test.expected {
			t.Error("Tidak Sesuai")
		}
	}
}

func BenchmarkPerkalian(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HitunganPerkalian(3, 4)
	}
}

func TestTablePenjumlahan(t *testing.T) {
	test := []struct {
		name     string
		x        int
		y        int
		expected int
	}{
		{name: "Penjumlahan1",
			x:        2,
			y:        3,
			expected: 5},
		{name: "Penjumlahan2",
			x:        4,
			y:        6,
			expected: 10},
		{name: "Penjumlahan3",
			x:        1,
			y:        1,
			expected: 2},
	}

	for _, test := range test {
		t.Run(test.name, func(t *testing.T) {
			result := Penjumlahan(test.x, test.y)
			assert.Equal(t, test.expected, result)
		})
	}
}

func TestSubTest(t *testing.T) {
	t.Run("Perkalian", func(t *testing.T) {
		x, y := 3, 4
		result := x * y
		expected := 12
		if result != expected {
			require.Equal(t, expected, result, "Hasil Harus Sama")
		}
	})

	t.Run("Pembagian", func(t *testing.T) {
		x, y := 2, 8
		result := y / x
		expect := 4
		if result != expect {
			require.Equal(t, expect, result, "Hasil Bagi Sama")
		}
	})
}

func TestMain(m *testing.M) {
	fmt.Println("Sebelum Unit Test")

	m.Run() //eksekusi semua unitest

	fmt.Println("Setelah Unit Test")
}

func TestHitunganPerkalian(t *testing.T) {
	result := HitunganPerkalian(3, 4)
	expected := 12

	if result != expected {
		assert.Equal(t, expected, result, "Hasil Harus 12")
	}

	fmt.Println("TestHitunganPerkalian Done")
}

func TestHitunganPembagian(t *testing.T) {
	result := HitunganPerkalian(4, 4)
	expected := 15

	if result != expected {
		assert.NotEqual(t, expected, result, "Hasil Harus Tidak Sesuai")
	}

	fmt.Println("TestHitunganPembagian Done")
}
