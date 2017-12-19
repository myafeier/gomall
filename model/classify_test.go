package model

import (
	"testing"
)

func TestFindAllClassifyChildIdSlice(t *testing.T) {
	InitDbForTest()
	result := FindAllClassifyChildIdSlice(8)
	t.Log(result)

}
