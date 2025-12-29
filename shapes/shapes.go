package shapes

import "math"

type Shape interface{
	Area() float64
}
type Rectangle struct {
	Width float64
	Height float64
}
type Circle struct{
	Radius float64
}
type Triangle struct{
	Height float64
	Base float64
}

func(c Circle) Area() float64{
	return  math.Pow(c.Radius,2) * math.Pi
}
func(r Rectangle) Area() float64{
	return  r.Height * r.Width
}
func(r Triangle) Area() float64{
	return  (r.Height * r.Base) / 2
}