package gen

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateControllerName(t *testing.T) {
	TestNames := []struct {
		Case        string
		Expectation bool
	}{
		{"Users", true},
		{"Us3rs", true},
		{"U_sers", false},
		{"U!sers", false},
		{"U@sers", false},
		{"U#sers", false},
		{"U$sers", false},
		{"U%sers", false},
		{"U^sers", false},
		{"U&sers", false},
		{"U*sers", false},
		{"U(sers", false},
		{"U)ers", false},
		{"U>ers", false},
		{"U<ers", false},
		{"U?ers", false},
		{"U[ers", false},
		{"U]ers", false},
	}

	for _, tn := range TestNames {
		t.Run(fmt.Sprintf("validating %s is %t", tn.Case, tn.Expectation), func(t *testing.T) {
			err := ValidateControllerName(tn.Case)
			if tn.Expectation {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}
