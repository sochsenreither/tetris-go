package game

import (
	"image/color"
	"reflect"
	"testing"
)

func TestClone(t *testing.T) {
	b := &Block{
		row: 0,
		col: 0,
		color: color.White,
		inactive: false,
	}

	if !reflect.DeepEqual(b, b.Clone()) {
		t.Errorf("Expected to get an exact copy")
	}
}
