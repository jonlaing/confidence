package confidence

import (
  "testing" 
  "time"
)

func TestZeroUpsAndDowns(t *testing.T) {
  c := Confidentable { 0, 0, time.Now() }

  if v := int(c.Confidence()); v != 0 {
    t.Error( "For 0 ups and downs, expected 0, got", v )
  }
}

func TestNonZeroUp(t *testing.T) {
  c := Confidentable { 3, 0, time.Now() }

  if c.ups == 0 {
    t.Error("Expected 1 ups, got", c.ups)
  }

  if c.downs != 0 {
    t.Error("Expected 0 downs, got", c.downs)
  }

  if v := c.Confidence(); v <= float64(0) {
    t.Error("For 1 up and 0 downs, expected not 0, got", v )
  }
}
