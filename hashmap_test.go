package github.com/slimaneakalia/goHashmap/hashmap_test

import (
	"testing"
	"github.com/slimaneakalia/goHashmap/hashmap"
)

func TestSetAndGet(t *testing.T) {
	h := NewHashMap()
	added, err := h.Set("key-1", "Value 1")
	if !added {
		t.Errorf("Failed to set in the hashmap, error: %v", err)
	} else {
		t.Log("Setting a value in a hashmap is working")
	}

	// Set a second value
	added, err = h.Set("key-2", "Value 2")

	if !added {
		t.Errorf("Failed to set a second value in the hashmap, error: %v", err)
	}

	// Get
	value := h.Get("key-1")
	if value != "Value 1" {
		t.Errorf("Failed to get from the hashmap, value: %v", value)
	} else {
		t.Log("Getting a value from hashmap is working")
	}

	// Set the same key
	added, err = h.Set("key-1", "Value 3")

	if !added {
		t.Errorf("Failed to set a value for an already used key in the hashmap, error: %v", err)
	}

	value = h.Get("key-1")
	if value != "Value 3" {
		t.Errorf("Failed to replace the value of a key in the hashmap, value: %v", value)
	} else {
		t.Log("Replacing a value in the hashmap is working")
	}

}
