package documents_test

import (
	"testing"

	"github.com/mauriciofsnts/vulcano/internal/providers/documents"
)

func TestGenerateCPF(t *testing.T) {
	cpfWithMask, cpfWithoutMask := documents.GenerateCPF()

	if len(cpfWithMask) != 14 {
		t.Errorf("Expected CPF mask to have length 14, but got %d", len(cpfWithMask))
	}

	if len(cpfWithoutMask) != 11 {
		t.Errorf("Expected CPF without mask to have length 11, but got %d", len(cpfWithoutMask))
	}
}
