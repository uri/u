package u

import (
	"encoding/json"
	"testing"
)

func TestJSONTuples(t *testing.T) {
	t.Run("good", func(t *testing.T) {
		got := JSONTuples("hello", "there")
		var m map[string]any
		Must(json.Unmarshal([]byte(got), &m))
		val := m["hello"]
		if str, ok := val.(string); !ok || str != "there" {
			t.Fatal("did not get the correct value")
		}
	})

	t.Run("odd args", func(t *testing.T) {
		got := JSONTuples("hello", "there", "invalid")
		var m map[string]any
		Must(json.Unmarshal([]byte(got), &m))
		val := m["hello"]
		if str, ok := val.(string); !ok || str != "there" {
			t.Fatal("did not get the correct value")
		}

		if _, ok := m["invalid"]; ok {
			t.Fatal("key should not be in map")
		}
	})
}
