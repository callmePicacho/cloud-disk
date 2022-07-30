package test

import (
	"cloud-disk/core/helper"
	"fmt"
	"testing"
)

func TestRandCode(t *testing.T) {
	fmt.Println(helper.GenerateVerifyCode())
}
