package documents_test

import (
	"testing"

	"github.com/mauriciofsnts/vulcano/internal/providers/documents"
)

func TestGenerateCPF(t *testing.T) {
	cpf := documents.GenerateCPF()

	if len(cpf) != 11 {
		t.Errorf("CPF should have 11 characters, got %d", len(cpf))
	}
}
