package model

import "testing"

func TestInitUserRankClassifyDiscount(t *testing.T) {
	InitDbForTest()
	err := InitUserRankClassifyDiscount()
	if err != nil {
		t.Error(err)
	}
	t.Log(CacheUserRankClassifyDiscount)
}
