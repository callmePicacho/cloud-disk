package test

import (
	"cloud-disk/core/helper"
	"fmt"
	"testing"
)

func TestMd5(t *testing.T) {
	md5 := helper.Md5("123")
	fmt.Println(md5 == "a554ccb472c14bdc6550b104962f7753")
}
