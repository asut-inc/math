// Package geometry contains top-level domain value objects and interfaces;
// it does not depend on any package but other packages depends on it value objects
package geometry

import "math"

// Coordinate represents N-dimensional coordinate (reserved for future usage, or task improvement)
type Coordinate []float64

// Point2D represents point (ordered tuple) among with its x and y coordinates
type Point2D struct {
	X float64
	Y float64
}

// LineSegment is represented by an ordered tuple of points
type LineSegment struct {
	Start Point2D
	End   Point2D
}

// NewLineSegment is a constructor-like function
// mostly for sugar in this case, than for establishing invariant
func NewLineSegment(start, end Point2D) LineSegment {
	return LineSegment{Start: start, End: end}
}

// Points2D represents collection of 2d points
type Points2D []Point2D

const (
	// EPS establishes approximation details, 1*10^-7, used mostly for float comparison
	EPS = 1e-7
)

// Equals checks (quite roughly) whether one point has the same coordinates as another
func (p *Point2D) Equals(pt Point2D) bool {
	return math.Abs(p.X-pt.X) < EPS && math.Abs(p.Y-pt.Y) < EPS
}
