package confidence

import (
  "math"
  "time"
)

type Confidentable struct {
  ups, downs int
  updated_at time.Time
}

///////////////////////////////////////

func (c Confidentable) score() int {
  return c.ups - c.downs
}


func (c Confidentable) Hot() float64 {
  var sign int

  score   := float64(c.score())
  order   := math.Log10( math.Max( math.Abs(score), 1) )
  seconds := float64(c.updated_at.Second() - 1134028003)

  if score > 0 {
    sign = 1
  } else if score < 0 {
    sign = -1
  } else {
    sign = 0
  }
  
  buf := int(order + float64(sign) * seconds / 45000.0) * int(math.Pow10(7))
  return float64( buf ) / math.Pow10(7)
}


func (c Confidentable) confidence() float64 {
  var n, z, phat float64

  n = float64(c.ups + c.downs)

  if n == 0 {
    return float64(0)
  }

  z = 1.0                  //1.0 = 85%, 1.6 = 95%
  phat = float64(c.ups)/n

  return math.Sqrt(phat+z*z/(2*n)-z*((phat*(1-phat)+z*z/(4*n))/n))/(1+z*z/n)
}


func (c Confidentable) Confidence() float64 {
  if c.ups + c.downs == 0 {
    return 0
  } else {
    return c.confidence()
  }
}
