package utiles

import (
	"os"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

var existingEntryNames = []string{".archivo-resultados-aqui", "resultados_test.go"}

func TestResultados(t *testing.T) {
	entries, err := os.ReadDir(".")
	if err != nil {
		assert.Error(t, err)
	}

	var filteredEntries []os.DirEntry
	for _, entry := range entries {
		if !slices.Contains(existingEntryNames, entry.Name()) {
			filteredEntries = append(filteredEntries, entry)
		}
	}

	if len(filteredEntries) == 0 {
		assert.Fail(t, "Archivo no encontrado")
	}

	for _, entry := range filteredEntries {
		info, err := entry.Info()
		if err != nil {
			assert.Error(t, err)
		}

		assert.NotEqual(t, int64(0), info.Size(), "El archivo "+entry.Name()+" está vacío")
	}
}
