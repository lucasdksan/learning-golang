package forms

import (
	"math"
	"testing"
)

func TestArea(t *testing.T) {
	t.Run("Rectangle", func(t *testing.T) {
		rect := Rectangle{10, 12}
		area_expected := float64(120)
		area_result := rect.Area()

		if area_expected != area_result {
			t.Fatalf("Área recebida %f é diferente da esperada %f", area_expected, area_result)
		}
	})

	t.Run("Circle", func(t *testing.T) {
		circ := Circle{10}
		area_expected := float64(math.Pi * 100)
		area_result := circ.Area()

		if area_expected != area_result {
			t.Fatalf("Área recebida %f é diferente da esperada %f", area_expected, area_result)
		}
	})
}
