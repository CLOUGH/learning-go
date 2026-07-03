// Package shapes exists purely to demonstrate Go's visibility rule:
// identifiers starting with a capital letter are exported (visible to
// other packages that import this one); lowercase ones are not, and the
// compiler enforces the boundary - it's not just a convention.
package shapes

// piApprox is unexported: no package outside `shapes` can name it.
// Try referencing shapes.piApprox from main.go - it won't compile.
const piApprox = 3.14159

// Circle is exported: usable as shapes.Circle from any importing package.
type Circle struct {
	Radius float64
}

// NewCircle is exported. Constructor functions like this are the usual
// way to build a value when a type has unexported fields to initialize,
// or invariants to check - callers outside the package can't reach in
// and construct the zero value incorrectly.
func NewCircle(radius float64) Circle {
	return Circle{Radius: radius}
}

// Area is an exported method, usable from outside the package. It's free
// to use piApprox internally even though piApprox itself isn't exported -
// visibility applies to whether OTHER packages can name the identifier,
// not to whether code within this package can use it.
func (c Circle) Area() float64 {
	return piApprox * c.Radius * c.Radius
}
