package constraintchecker

import (
	"testing"
)

func testConstraint(t *testing.T, isValid bool, symbol string, value string, params []string) {
	c, _ := ConstrainFuncMap[symbol]
	valid, err := c(value, params...)
	if err != nil {
		t.Errorf("Constraint failed with an error %s on constraint %s", err, symbol)
	}
	if valid != isValid {
		if isValid {
			t.Errorf("False negative on constraint %s", symbol)
		} else {
			t.Errorf("False positive on constraint %s", symbol)
		}
	}
}

// Test each bound: (, ), [, ], [), [], (], ()
func TestBounds(t *testing.T) {
	// For )
	testConstraint(t, true, ")", "5", []string{"10"})
	testConstraint(t, false, ")", "11", []string{"10"})
	testConstraint(t, false, ")", "10", []string{"10"})
	// For (
	testConstraint(t, true, "(", "11", []string{"10"})
	testConstraint(t, false, "(", "5", []string{"10"})
	testConstraint(t, false, "(", "10", []string{"10"})
	// For [
	testConstraint(t, true, "[", "10", []string{"10"})
	testConstraint(t, true, "[", "12", []string{"10"})
	testConstraint(t, false, "[", "5", []string{"10"})
	// For ]
	testConstraint(t, true, "]", "10", []string{"10"})
	testConstraint(t, true, "]", "8", []string{"10"})
	testConstraint(t, false, "]", "11", []string{"10"})
	// For [)
	testConstraint(t, true, "[)", "10", []string{"10","11"})
	testConstraint(t, false, "[)", "11", []string{"10","11"})
	// For []
	testConstraint(t, true, "[]", "10", []string{"10","11"})
	testConstraint(t, false, "[]", "12", []string{"10","11"})
	// For (]
	testConstraint(t, true, "(]", "11", []string{"10","11"})
	testConstraint(t, false, "(]", "10", []string{"10","11"})
	// For ()
	testConstraint(t, true, "()", "11", []string{"10","12"})
	testConstraint(t, false, "()", "11", []string{"10","11"})
}

func TestMember(t *testing.T) {
	// For =
	testConstraint(t, true, "=", "foo", []string{"bar","foo"})
	testConstraint(t, false, "=", "foo", []string{"bar","baz"})
}
