package template

import "testing"

// Tabella di test: ogni caso è una struct con input e output atteso.
// Aggiungi righe a testCases per coprire più scenari.
func TestAdd(t *testing.T) {
	testCases := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"positivi", 2, 3, 5},
		{"zero", 0, 0, 0},
		{"negativi", -1, 1, 0},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := Add(tc.a, tc.b)
			if actual != tc.expected {
				t.Errorf("Add(%d, %d) = %d; expected %d", tc.a, tc.b, actual, tc.expected)
			}
		})
	}
}
