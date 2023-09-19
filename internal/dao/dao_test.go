package dao

import (
	"context"
	"fmt"
	"testing"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
)

func TestUserDao_GetUser(t *testing.T) {
	u, err := User.GetUser(context.Background(), 1)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("%+v\n", u)
}
