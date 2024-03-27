package address

import (
	"strings"
)

func Address_type(address string) string {
	approved_types := []string{"rua", "avenida", "estrada", "rodovia"}
	address_lower_case := strings.ToLower(address)
	first_word_addres := strings.Split(address_lower_case, " ")[0]
	result_check_address := false

	for _, addr := range approved_types {
		if first_word_addres == addr {
			result_check_address = true
		}
	}

	if result_check_address {
		return first_word_addres
	}

	return "Tipo Inválido"
}
