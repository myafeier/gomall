package model

import (
	"testing"
)

func TestInitCacheSafeGoodsSpecAttr(t *testing.T) {
	InitDbForTest()

	err := InitCacheSafeGoodsSpecAttr()
	if err != nil {
		t.Error(err)
	}
	t.Log("attr:", CacheSafeGoodsSpecAttr)

}

func TestInitCacheSafeGoodsSpecValue(t *testing.T) {
	InitDbForTest()
	err := InitCacheSafeGoodsSpecValue()
	if err != nil {
		t.Error(err)
	}
	t.Log("value:", CacheSafeGoodsSpecValue)

}
