package gofmtkata

import "fmt"

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func Describe(r Rectangle) string {
	if r.Width == r.Height {
		return fmt.Sprintf("a %vx%v square", r.Width, r.Height)
	} else {
		return fmt.Sprintf("a %vx%v rectangle", r.Width, r.Height)
	}
}
