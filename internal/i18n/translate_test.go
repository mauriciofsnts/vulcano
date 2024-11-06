package i18n_test

import (
	"testing"

	"github.com/mauriciofsnts/bot/internal/i18n"
	"github.com/stretchr/testify/assert"
)

func TestReplace(t *testing.T) {
	text := "${0} está ${1}"
	replaced := i18n.Replace(text, "Vulcano", "online")
	assert.Equal(t, "${0} está ${1}", replaced)
}
