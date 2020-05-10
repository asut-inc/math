// Package intersection provides functionality for intersection detection
package intersection

import (
	"github.com/asutd/geom/geometry"
)

// Type represents intersection type
type Type int

const (
	// NoIntersection indicates that line segments do not intersect
	NoIntersection Type = iota
	// PointIntersection indicates that line segments intersect at a point
	PointIntersection
	// CollinearIntersection indicates that line segments overlap each other
	CollinearIntersection
)

var labels = [3]string{"NoIntersection", "PointIntersection", "CollinearIntersection"}

// Intersector establishes intersection detection interface
type Intersector interface {
	// LineSegmentIntersection takes two line segments and returns intersection points
	// if no intersection points was found len() is eq to 0
	// if 1 intersection points it is equal to 1
	// otherwise, in case of collinear intersection, it includes both endpoints
	LineSegmentIntersection(lineSegment1, lineSegment2 geometry.LineSegment) geometry.Points2D
}

// IntersectorFunc is an adapter to allow the use of own implementations without package modification
type IntersectorFunc func(lineSegment1, lineSegment2 geometry.LineSegment) geometry.Points2D

// LineSegmentIntersection makes possible to use own functions for intersection detection
// strictly speaking, this a realization of robustness principle,
// "Be conservative in what you do, be liberal in what you accept from others"
func (i IntersectorFunc) LineSegmentIntersection(lineSegment1, lineSegment2 geometry.LineSegment) geometry.Points2D {
	return i(lineSegment1, lineSegment2)
}

// Result includes type of intersection and intersection point(s)
type Result struct {
	intersectionType Type
	intersection     geometry.Points2D
}

// NewResult creates a new result object
func NewResult(intersectionType Type, coords geometry.Points2D) *Result {
	return &Result{
		intersectionType: intersectionType,
		intersection:     coords,
	}
}

// HasIntersection returns true if line segments have an intersection
func (r *Result) HasIntersection() bool {
	return r.intersectionType != NoIntersection
}

// IntersectionType returns the type of intersection between the two lines
func (r *Result) IntersectionType() Type {
	return r.intersectionType
}

// String is to satisfy Stringer interface
func (t Type) String() string {
	return labels[t]
}

// Intersection returns an array of Coords which are the intersection points.
// If the type is PointIntersection then there will only be a single Coordinate (the first coord).
// If the type is CollinearIntersection then there will two Coordinates the start and end points of the line
// that represents the intersection
func (r *Result) Intersection() geometry.Points2D {
	return r.intersection
}

// LineSegmentIntersectsLineSegment takes two line segments, calls specified implementation,
// and returns intersection result among with its type
func LineSegmentIntersectsLineSegment(algorithm Intersector, lineSegment1 geometry.LineSegment, lineSegment2 geometry.LineSegment) *Result {

	points := algorithm.LineSegmentIntersection(lineSegment1, lineSegment2)
	n := len(points)

	var intersectionType Type
	switch n {
	case 0:
		intersectionType = NoIntersection
	case 1:
		intersectionType = PointIntersection
	default:
		intersectionType = CollinearIntersection
	}

	return NewResult(intersectionType, points)
}
