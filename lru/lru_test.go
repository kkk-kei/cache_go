package lru

import (
	"reflect"
	"testing"
)

type String string

func (d String) Len() int {
	return len(d)
}
func TestGet(t *testing.T) {
	lru := New(int64(0), nil)
	lru.Add("k1", String("1234"))
	if v, ok := lru.Get("k1"); !ok || string(v.(String)) != "1234" {
		t.Fatalf("cache hit k1=1234 faild")
	}
	if _, ok := lru.Get("k2"); ok {
		t.Fatalf("cache miss k2 faild")
	}
}

func TestRemoveOldest(t *testing.T) {
	k1, k2, k3 := "k1", "k2", "k3"
	v1, v2, v3 := "v1", "v2", "v3"
	cap := len(k1 + k2 + v1 + v2)
	lru := New(int64(cap), nil)
	lru.Add(k1, String(v1))
	lru.Add(k2, String(v2))
	lru.Add(k3, String(v3))
	if _, ok := lru.Get(k1); ok || lru.Len() != 2 {
		t.Fatalf("RemoveOldest k1 faild")
	}
}

func TestOnEvicted(t *testing.T) {
	keys := make([]string, 0)
	callback := func(key string, value Value) {
		keys = append(keys, key)
	}
	k1, k2, k3, k4 := "k1", "k2", "k3", "k4"
	v1, v2, v3, v4 := "v1", "v2", "v3", "v4"
	cap := len(k1 + k2 + v1 + v2)
	lru := New(int64(cap), callback)
	lru.Add(k1, String(v1))
	lru.Add(k2, String(v2))
	lru.Add(k3, String(v3))
	lru.Add(k4, String(v4))
	expect := []string{"k1", "k2"}
	if !reflect.DeepEqual(expect, keys) {
		t.Fatalf("Call OnEvicted faild,expect keys equal to %s", expect)
	}
}
