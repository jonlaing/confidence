package confidence

import (
	"testing" 
)

func TestZeroUpsAndDowns(t *testing.T) {
	c := Confidentable { 0, 0 }

	if v := int(c.Score()); v != 0 {
		t.Error( "For 0 ups and downs, expected 0, got", v )
	}
}
