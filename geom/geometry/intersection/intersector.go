package intersection

import (
	"github.com/asutd/geom/geometry"
	"math"
)

// LineIntersector is an implementation of 'strategy' pattern
type LineIntersector struct{}

// Orientation represents orientation for points triple
type Orientation int

const (
	// Collinear when three points are collinear
	Collinear Orientation = iota
	// Clockwise is when one point is clockwise to segment (a,b) that is right formed
	Clockwise
	// CounterClockwise is when one point is counter clockwise to segment (a,b), that is left of line
	CounterClockwise
)

// orientation finds the orientation of point 'p' relative to the line segment [a,b]
// picture representation is here https://media.geeksforgeeks.org/wp-content/uploads/point1.png
// for details refer https://www.geeksforgeeks.org/direction-point-line-segment/
func orientation(a, b, p geometry.Point2D) Orientation {

	// here, we are going to utilize cross-product from vector algebra
	// which has a neat property that helps to determine direction of a point from a line segment

	// If the sign of this cross product is negative, then  is counterclockwise with respect to [a,b],
	// A positive cross product indicates a clockwise orientation and a right turn.
	// A cross product of 0 means that points are collinear.

	// http://mathworld.wolfram.com/CrossProduct.html (8) and (9) for Tex notation
	// http://staff.ustc.edu.cn/~csli/graduate/algorithms/book6/888_a.gif

	det := (b.Y-a.Y)*(p.X-b.X) - (b.X-a.X)*(p.Y-b.Y)
	if math.Abs(det) < geometry.EPS {
		return Collinear
	}

	if det > 0 {
		return Clockwise
	}

	return CounterClockwise

}

// tests whether point 'p' is on the line segment [a,b]
// ensure first that point 'p' is collinear to segment [a,b] and
// then check whether 'p' is within
func pointOnLineSegment(a, b, p geometry.Point2D) bool {
	return orientation(a, b, p) == Collinear &&
		math.Min(a.X, b.X) <= p.X &&
		p.X <= math.Max(a.X, b.X) &&
		math.Min(a.Y, b.Y) <= p.Y &&
		p.Y <= math.Max(a.Y, b.Y)
}

// determines if two segments intersects
func segmentsIntersect(p1, p2, p3, p4 geometry.Point2D) bool {
	o1 := orientation(p1, p2, p3)
	o2 := orientation(p1, p2, p4)
	o3 := orientation(p3, p4, p1)
	o4 := orientation(p3, p4, p2)

	// if points {p1, p2} are on opposite side of line formed by [p3, p4] and conversely
	// {p3, p4} on the opposite side of line formed by (p1,p2) then there is an intersection
	if o1 != o2 && o3 != o4 {
		return true
	}

	// collinear special cases
	if o1 == Collinear && pointOnLineSegment(p1, p2, p3) {
		return true
	}
	if o2 == Collinear && pointOnLineSegment(p1, p2, p4) {
		return true
	}
	if o3 == Collinear && pointOnLineSegment(p3, p4, p1) {
		return true
	}
	if o4 == Collinear && pointOnLineSegment(p3, p4, p2) {
		return true
	}

	return false
}

// getCommonEndpoints performs point-by-point comparison for equality appending to final result if any of endpoints are the same
func getCommonEndpoints(p1, p2, p3, p4 geometry.Point2D) []geometry.Point2D {
	commonPoints := make([]geometry.Point2D, 0, 2)
	if p1.Equals(p3) {
		commonPoints = append(commonPoints, p1)
		if p2.Equals(p4) {
			commonPoints = append(commonPoints, p2)
		}
	} else if p1.Equals(p4) {
		commonPoints = append(commonPoints, p1)
		if p2.Equals(p3) {
			commonPoints = append(commonPoints, p2)
		}
	} else if p2.Equals(p3) {
		commonPoints = append(commonPoints, p2)
		if p1.Equals(p4) {
			commonPoints = append(commonPoints, p1)
		}
	} else if p2.Equals(p4) {
		commonPoints = append(commonPoints, p2)
		if p1.Equals(p3) {
			commonPoints = append(commonPoints, p1)
		}
	}

	return commonPoints
}

// solve takes line segments endpoints and derives x,y coordinates of intersection point from them
func solve(p1, p2, p3, p4 geometry.Point2D) (float64, float64) {

	m1 := (p2.Y - p1.Y) / (p2.X - p1.X)
	m2 := (p4.Y - p3.Y) / (p4.X - p3.X)
	b1 := p1.Y - m1*p1.X
	b2 := p3.Y - m2*p3.X
	x := (b2 - b1) / (m1 - m2)
	y := (m1*b2 - m2*b1) / (m1 - m2)

	return x, y
}

// LineSegmentIntersection is core implementation of line segment intersection detection
func (l *LineIntersector) LineSegmentIntersection(lineSegment1, lineSegment2 geometry.LineSegment) geometry.Points2D {

	p1, p2 := lineSegment1.Start, lineSegment1.End
	p3, p4 := lineSegment2.Start, lineSegment2.End

	if !segmentsIntersect(p1, p2, p3, p4) {
		return nil
	}

	// case, when both line segments are a single point
	if p1.Equals(p2) && p2.Equals(p3) && p3.Equals(p4) {
		return geometry.Points2D{p1}
	}

	commonEndpoints := getCommonEndpoints(p1, p2, p3, p4)
	commonEndpointsNumber := len(commonEndpoints) // how much endpoints are common for both segments

	// one of the line segments is an intersecting single point
	// checking only commonEndpointsNumber == 1 is insufficient to return because the solution might be a sub segment
	singlePoint := p1.Equals(p2) || p3.Equals(p4)
	if commonEndpointsNumber == 1 && singlePoint {
		return geometry.Points2D{commonEndpoints[0]}
	}

	// segments are equal
	if commonEndpointsNumber == 2 {
		return geometry.Points2D{commonEndpoints[0], commonEndpoints[1]}
	}

	collinearSegments := orientation(p1, p2, p3) == Collinear && orientation(p1, p2, p4) == Collinear

	// the intersection will be a sub-segment of the two segment since they overlap each other

	if collinearSegments {

		// segment #2 is enclosed in segment #1
		if pointOnLineSegment(p1, p2, p3) && pointOnLineSegment(p1, p2, p4) {
			return geometry.Points2D{p3, p4}
		}

		// segment #1 is enclosed in segment #2
		if pointOnLineSegment(p3, p4, p1) && pointOnLineSegment(p3, p4, p2) {
			return geometry.Points2D{p1, p2}
		}

		// the sub segment is part of segment #1 and part of segment #2
		// find the middle points which corresponds to this segment
		var (
			midpoint1 geometry.Point2D
			midpoint2 geometry.Point2D
		)
		if pointOnLineSegment(p1, p2, p3) {
			midpoint1 = p3
		} else {
			midpoint1 = p4
		}
		if pointOnLineSegment(p3, p4, p1) {
			midpoint2 = p1
		} else {
			midpoint2 = p2
		}

		// there is only one middle point
		if midpoint1.Equals(midpoint2) {
			return geometry.Points2D{midpoint1}
		}

		return geometry.Points2D{midpoint1, midpoint2}

	}

	// here we have a unique intersection point

	// even seventh grade school formulas converted to code sometimes confuses
	// but here we have an extremely neat part
	// use the points coordinates to build lines equations in form y = mx + b
	// with making some checks for vertical lines and you are done

	// segment #1 is a vertical line
	if math.Abs(p1.X-p2.X) < geometry.EPS {
		m := (p4.Y - p3.Y) / (p4.X - p3.X)
		b := p3.Y - m*p3.X
		return geometry.Points2D{geometry.Point2D{X: p1.X, Y: m*p1.X + b}}
	}

	// segment #2 is a vertical line
	if math.Abs(p3.X-p4.X) < geometry.EPS {
		m := (p2.Y - p1.Y) / (p2.X - p1.X)
		b := p1.Y - m*p1.X
		return geometry.Points2D{geometry.Point2D{X: p3.X, Y: m*p3.X + b}}
	}

	x, y := solve(p1, p2, p3, p4)

	return geometry.Points2D{geometry.Point2D{X: x, Y: y}}

}
