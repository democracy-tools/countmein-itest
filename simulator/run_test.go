package simulator

import (
	"testing"
)

func TestRun(t *testing.T) {

	NewRealm(50).Run(1000)
}
