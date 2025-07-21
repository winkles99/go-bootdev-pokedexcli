package pokecache

import (
	"fmt"
	"testing"
	"time"

	"github.com/winkles99/go-bootdev-pokedexcli/internal/pokeapi"
)

func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second
	prev20 := "bla?offset=0"
	next20 := "bla?offset=40"
	prev40 := "bla?offset=20"
	next40 := "bla?offset=60"
	cases := []struct {
		key string
		val pokeapi.LocationAreasResp
	}{
		{
			key: "http://test.com/?ofset=20",
			val: pokeapi.LocationAreasResp{Count: 200, Previous: &prev20, Next: &next20, Results: nil},
		},
		{
			key: "http://test.com/?ofset=40",
			val: pokeapi.LocationAreasResp{Count: 200, Previous: &prev40, Next: &next40, Results: nil},
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add(c.key, c.val)
			v, ok := cache.Get(c.key)
			val := v.(pokeapi.LocationAreasResp)
			if !ok {
				t.Errorf("expected to find key")
				return
			}
			if val.Next != c.val.Next || val.Previous != c.val.Previous || val.Count != c.val.Count {
				t.Errorf("expected to find value")
				return
			}
		})
	}
}

func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + 5*time.Millisecond
	cache := NewCache(baseTime)
	prev := "bla?offset=20"
	next := "bla?offset=60"
	cache.Add("http://test.com/?ofset=40", pokeapi.LocationAreasResp{Count: 200, Previous: &prev, Next: &next, Results: nil})

	_, ok := cache.Get("http://test.com/?ofset=40")
	if !ok {
		t.Errorf("expected to find key")
		return
	}

	time.Sleep(waitTime)

	_, ok = cache.Get("http://test.com/?ofset=40")
	if ok {
		t.Errorf("expected to not find key")
		return
	}
}