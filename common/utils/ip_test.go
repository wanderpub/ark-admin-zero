package utils

import (
	"fmt"
	"testing"
)

func Test_IpLocation(t *testing.T) {
	location, _ := GetIPLocation("101.69.251.46")
	fmt.Println(location)
}
