package simulator

import (
	"testing"
)

func TestRun(t *testing.T) {

	NewRealm(30).Run(1000)
}
