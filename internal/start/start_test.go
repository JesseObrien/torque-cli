package start

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAppNameValidation(t *testing.T) {
	testnames := []struct {
		name    string
		isValid bool
	}{
		{"testapp", true},
		{"TESTAPP", true},
		{"test-app", true},
		{"test_app", true},
		{"TEST_APP", true},
		{"my-test-app", true},
		{"testapp!", false},
		{"testapp.", false},
		{"testapp*", false},
		{"testapp(", false},
		{"testapp)", false},
	}

	for _, data := range testnames {
		t.Run(fmt.Sprintf("Expecting %s is %t", data.name, data.isValid), func(t *testing.T) {
			assert.Equal(t, data.isValid, isValidAppName(data.name))
		})
	}
}
