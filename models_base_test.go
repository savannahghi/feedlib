package go_utils_test

import (
	"testing"

	base "github.com/savannahghi/go_utils"
)

func TestModelsIsEntity(t *testing.T) {

	t11 := base.EmailOptIn{}
	t11.IsEntity()

	t12 := base.USSDSessionLog{}
	t12.IsEntity()

	t13 := base.PhoneOptIn{}
	t13.IsEntity()
}
