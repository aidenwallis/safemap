package safemap 

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMapCreation(t *testing.T) {
	m := New()
	if m == nil {
		t.Error("Map is null")
	}
	if m.Count() != 0 {
		t.Error("Map is not empty")
	}
}

func TestSet(t *testing.T) {
	m := New()

	m.Set("one", "one")
	m.Set("two", "two")
	m.Set("three", "three")

	if m.Count() != 3 {
		t.Error("map should contain exactly three elements.")
	}

	v, ok := m.Get("one")
	if !ok {
		t.Error("Values didn't set in map")
	}
	assert.Equal(t, v, "one")

	v, ok = m.Get("two")
	if !ok {
		t.Error("Values didn't set in map")
	}
	if v != "two" {
		t.Error("Incorrect value set in map")
	}

	v, ok = m.Get("three")
	if !ok {
		t.Error("Values didn't set in map")
	}
	if v != "three" {
		t.Error("Incorrect value set in map")
	}
}

func TestGet(t *testing.T) {
	m := New()
	m.Set("Pog", "Champ")

	value, ok := m.Get("Pog")
	if !ok {
		t.Error("Failed to get value from map")
	}
	if value != "Champ" {
		t.Error("Failed to get correct value from map")
	}

	_, ok = m.Get("doesntexist")
	if ok {
		t.Error("Managed to get value a sok when it doesn't exist")
	}
}

func TestHas(t *testing.T) {
	m := New()
	m.Set("Pepe", "Ga")

	has := m.Has("PogChamp")
	if has {
		t.Error("Managed to get value which doesn't exist")
	}

	has = m.Has("Pepe")
	if !has {
		t.Error("Didn't manage to get a value set in the map")
	}
}

func TestCount(t *testing.T) {
	m := New()

	m.Set("Pepe", "Ga")
	m.Set("monka", "S")
	m.Set("Pog", "Champ")

	if m.Count() != 3 {
		t.Error("Incorrect count returned")
	}
}

func TestDelete(t *testing.T) {
	m := New()

	m.Set("Pepe", "Ga")
	m.Set("monka", "S")
	m.Set("Pog", "Champ")

	if m.Count() != 3 {
		t.Error("Incorrect count returned")
	}

	m.Delete("Pepe")

	if m.Count() != 2 {
		t.Error("Incorrect count returned after delete")
	}

	if m.Has("Pepe") {
		t.Error("Key still exists after delete")
	}
}
