package address

import (
	"testing"
)

type test_scenario struct {
	entered_address string
	expected_return string
}

func TestAddress_type(t *testing.T) {
	t.Parallel() // Roda os testes em paralelo

	test_scenarios := []test_scenario{
		{"Rua ABC", "rua"},
		{"Rua Capitão Martion Machado", "rua"},
		{"Avenida Santas Cruz", "avenida"},
		{"Avenida Duque de Caxias", "avenida"},
		{"Estrada Lagoa Queimada", "estrada"},
		{"Estrada Lopes", "estrada"},
		{"Rodovia BR-304", "rodovia"},
		{"Rodovia BR-116", "rodovia"},
		{"Zona Rural Medeiros", "Tipo Inválido"},
	}

	for _, scenario := range test_scenarios {
		address_type_gift := Address_type(scenario.entered_address)

		if address_type_gift != scenario.expected_return {
			t.Error("O tipo recebido é diferente do esperado")
		}
	}
}
