package confidence

import "math"

type Confidentable struct {
  ups, downs int
}

func (c Confidentable) score() float64 {
  var n, z, phat float64

  n = float64(c.ups + c.downs)

  if n == 0 {
    return float64(0)
  }

  z = 1.0 //1.0 = 85%, 1.6 = 95%

  phat = float64(c.ups)/n
  return math.Sqrt(phat+z*z/(2*n)-z*((phat*(1-phat)+z*z/(4*n))/n))/(1+z*z/n)
}

func (c Confidentable) Score() float64 {
  if c.ups + c.downs == 0 {
    return 0
  } else {
    return c.score()
  }
}
