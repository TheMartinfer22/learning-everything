package exception

import "fmt"

type Exception struct {
	Code    int
	Message string
}

func (e Exception) Error() string {
	return fmt.Sprintf("Erro ao processar dados, code=%d, message=%s",
		e.Code,
		e.Message,
	)
}
