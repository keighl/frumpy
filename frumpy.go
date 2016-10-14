package frumpy

import (
	"encoding/json"
	"strings"
)

func FilterJSON(in []byte, badKeys ...string) ([]byte, error) {
	targ := map[string]interface{}{}
	err := json.Unmarshal(in, &targ)
	if err != nil {
		return in, err
	}

	for _, bk := range badKeys {
		// Sniff for JSON dot notation
		splitz := strings.Split(bk, ".")
		if len(splitz) == 1 {
			delete(targ, bk)
			continue
		}

		huntDownKeys(splitz, 0, huntMap{targ})
	}

	b, err := json.Marshal(targ)
	if err != nil {
		return in, err
	}

	return b, nil
}

type huntMap struct {
	object map[string]interface{}
}

// huntDownKeys recursively walks down a JSON map to kill keys it finds. It will step into
// arrays in search of more objects with keys to slay
func huntDownKeys(dotKeys []string, keyIndex int, hm huntMap) {
	// K, we're doing some dot-notation
	// First, let's remember how deep we need to go
	depth := len(dotKeys)
	currKey := dotKeys[keyIndex]

	// If we're at the bottom of the keys, just delete it!
	if keyIndex == (depth - 1) {
		delete(hm.object, currKey)
		return
	}
	// If we've made it this far, it means we're dealing with a map,
	// and we're not at the bottom of the key nest yet

	// Next, see if we can go deeper into a Map
	nextMap, ok := hm.object[currKey].(map[string]interface{})
	if ok {
		// Ok, it's an object. Deeper we go!
		nextIndex := keyIndex + 1
		huntDownKeys(dotKeys, nextIndex, huntMap{nextMap})
		return
	}

	// So it's not a map, maybe a slice?
	nextSlice, ok := hm.object[currKey].([]interface{})
	if ok {
		// Interesting, a nested slice
		// What we'll do is go through the slice, looking for more maps
		for _, val := range nextSlice {
			nextMap, ok := val.(map[string]interface{})
			if ok {
				// It's a map! So recurse it again
				nextIndex := keyIndex + 1
				huntDownKeys(dotKeys, nextIndex, huntMap{nextMap})
			}
		}
	}

	// Ok, nothing we can deal with. Give up!
}
