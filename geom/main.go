package main

import (
	"fmt"
	"github.com/asutd/geom/geometry"
	"github.com/asutd/geom/geometry/intersection"
)

func main() {

	// first example with standard library algorithm
	firstLineSegment := geometry.NewLineSegment(geometry.Point2D{X: 10, Y: 20}, geometry.Point2D{X: 30, Y: 40})
	secondLineSegment := geometry.NewLineSegment(geometry.Point2D{X: 30, Y: 40}, geometry.Point2D{X: 40, Y: 80})
	intersectionDetectionAlg := &intersection.LineIntersector{}
	res := intersection.LineSegmentIntersectsLineSegment(intersectionDetectionAlg, firstLineSegment, secondLineSegment)
	if res.HasIntersection() {
		fmt.Printf("specified line segments intersected: intersection is of a type %s, actual points: %#v \n", res.IntersectionType(), res.Intersection())
	} else {
		fmt.Printf("specified line segments does not intersect")
	}

	// snippet just to show that you can provide your own implementation without touching library(so-called) internals
	// only relying on top-level domain language
	// I usually return function with closure context to be able to extend functionality making it less rigid
	ownImplementation := func() intersection.IntersectorFunc {
		return func(lineSegment1, lineSegment2 geometry.LineSegment) geometry.Points2D {
			// implement me!
			return geometry.Points2D{}
		}
	}
	_ = intersection.LineSegmentIntersectsLineSegment(ownImplementation(), firstLineSegment, secondLineSegment)

}
