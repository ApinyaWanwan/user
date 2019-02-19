// +buid integration

package user_test

import (
	"fmt"
	"testing"

	"github.com/ApinyaWanwan/user"
)

func TestIntegrationGetUser(t *testing.T) {
	u, err := user.Get()

	if err != nil {
		t.Error(err)
	}

	fmt.Print(u)
}
