package test

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"testing"
)

func TestUUID(t *testing.T) {
	fmt.Println(uuid.NewV4().String())
}
