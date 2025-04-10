package calculadora

import "errors"

// Nome da função deve ser maiúsculo para ser exportada igual ao Java
func Dividir (a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("não é possível dividir por zero")
	}
	return (a / b), nil
}