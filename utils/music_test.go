package utils

import (
	"fmt"
	"testing"
)

func TestMusicGet(t *testing.T) {

	s := MusicGet(2175288)
	fmt.Println(s.Url)
}
