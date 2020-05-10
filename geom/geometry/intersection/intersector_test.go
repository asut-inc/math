package intersection

import (
	"github.com/asutd/geom/geometry"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntersector(t *testing.T) {

	tests := []struct {
		name              string
		firstLineSegment  geometry.LineSegment
		secondLineSegment geometry.LineSegment
		expectedResult    geometry.Points2D
		visualization     string
	}{
		{
			name: "line segments intersect in one point",
			firstLineSegment: geometry.LineSegment{
				Start: geometry.Point2D{X: -4, Y: -5},
				End:   geometry.Point2D{X: 7, Y: 28},
			},
			secondLineSegment: geometry.LineSegment{
				Start: geometry.Point2D{X: -3, Y: 3},
				End:   geometry.Point2D{X: 7, Y: 23},
			},
			expectedResult: geometry.Points2D{geometry.Point2D{X: 2, Y: 13}},
			visualization:  "https://www.desmos.com/calculator/jmxcp7alup",
		},

		{
			name: "infinite lines intersect, but line segments do not",
			firstLineSegment: geometry.LineSegment{
				Start: geometry.Point2D{X: -4, Y: -5},
				End:   geometry.Point2D{X: 7, Y: 28},
			},
			secondLineSegment: geometry.LineSegment{
				Start: geometry.Point2D{X: 5, Y: 19},
				End:   geometry.Point2D{X: 7, Y: 23},
			},
			expectedResult: geometry.Points2D(nil),
			visualization:  "https://www.desmos.com/calculator/wezrniehtc",
		},

		{
			name: "one segment is sub segment of another",
			firstLineSegment: geometry.LineSegment{
				Start: geometry.Point2D{X: -4, Y: -5},
				End:   geometry.Point2D{X: 7, Y: 28},
			},
			secondLineSegment: geometry.LineSegment{
				Start: geometry.Point2D{X: 5, Y: 22},
				End:   geometry.Point2D{X: 7, Y: 28},
			},
			expectedResult: geometry.Points2D{geometry.Point2D{X: 5, Y: 22}, geometry.Point2D{X: 7, Y: 28}},
			visualization:  "https://www.desmos.com/calculator/u28zwq1j5o",
		},

		{
			name: "intersects vertical line segment",
			firstLineSegment: geometry.LineSegment{
				Start: geometry.Point2D{X: 4, Y: 0},
				End:   geometry.Point2D{X: 4, Y: 8},
			},
			secondLineSegment: geometry.LineSegment{
				Start: geometry.Point2D{X: -4, Y: -8},
				End:   geometry.Point2D{X: 7, Y: 14},
			},
			expectedResult: geometry.Points2D{geometry.Point2D{X: 4, Y: 8}},
			visualization:  "https://www.desmos.com/calculator/fl0bfjbfe2?lang=ru",
		},

		{
			name: "does not intersect vertical line segment",
			firstLineSegment: geometry.LineSegment{
				Start: geometry.Point2D{X: 4, Y: 0},
				End:   geometry.Point2D{X: 4, Y: 8},
			},
			secondLineSegment: geometry.LineSegment{
				Start: geometry.Point2D{X: -4, Y: -8},
				End:   geometry.Point2D{X: 3, Y: 6},
			},
			expectedResult: geometry.Points2D(nil),
			visualization:  "https://www.desmos.com/calculator/dnrtvbalyh?lang=ru",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := LineIntersector{}
			points := i.LineSegmentIntersection(tt.firstLineSegment, tt.secondLineSegment)
			assert.Equal(t, tt.expectedResult, points)
		})
	}
}

func TestOrientation(t *testing.T) {
	tests := []struct {
		name           string
		lineSegment    geometry.LineSegment
		point          geometry.Point2D
		expectedResult Orientation
	}{
		{
			name: "point is left to line segment",
			lineSegment: geometry.LineSegment{
				Start: geometry.Point2D{
					X: -4,
					Y: -9.2,
				},
				End: geometry.Point2D{
					X: 3,
					Y: 6.9,
				},
			},
			point: geometry.Point2D{
				X: -9.1,
				Y: -11.04,
			},
			expectedResult: CounterClockwise,
		},
		{
			name: "point is right to line segment",
			lineSegment: geometry.LineSegment{
				Start: geometry.Point2D{
					X: -4,
					Y: -9.2,
				},
				End: geometry.Point2D{
					X: 3,
					Y: 6.9,
				},
			},

			point: geometry.Point2D{
				X: 2.7,
				Y: -11.55,
			},
			expectedResult: Clockwise,
		},

		{
			name: "point is collinear to line segment",
			lineSegment: geometry.LineSegment{
				Start: geometry.Point2D{
					X: -4,
					Y: -9.2,
				},
				End: geometry.Point2D{
					X: 3,
					Y: 6.9,
				},
			},

			point: geometry.Point2D{
				X: 0,
				Y: 0,
			},
			expectedResult: Collinear,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := orientation(tt.lineSegment.Start, tt.lineSegment.End, tt.point)
			assert.Equal(t, tt.expectedResult, o)
		})
	}
}

func TestPointOnLineSegment(t *testing.T) {
	tests := []struct {
		name           string
		lineSegment    geometry.LineSegment
		point          geometry.Point2D
		expectedResult bool
	}{
		{
			name: "point not on line segment",
			lineSegment: geometry.LineSegment{
				Start: geometry.Point2D{
					X: 3,
					Y: 6.9,
				},
				End: geometry.Point2D{
					X: -4,
					Y: -9.2,
				},
			},
			point: geometry.Point2D{
				X: -2.5,
				Y: 1.75,
			},
			expectedResult: false,
		},

		{
			name: "point on line segment",
			lineSegment: geometry.LineSegment{
				Start: geometry.Point2D{
					X: 3,
					Y: 6.9,
				},
				End: geometry.Point2D{
					X: -4,
					Y: -9.2,
				},
			},
			point: geometry.Point2D{
				X: 0,
				Y: 0,
			},
			expectedResult: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := pointOnLineSegment(tt.lineSegment.Start, tt.lineSegment.End, tt.point)
			assert.Equal(t, tt.expectedResult, res)
		})
	}
}

func TestSolve(t *testing.T) {

	segment1 := geometry.LineSegment{Start: geometry.Point2D{X: -4, Y: -5}, End: geometry.Point2D{X: 7, Y: 28}}
	segment2 := geometry.LineSegment{Start: geometry.Point2D{X: -3, Y: 3}, End: geometry.Point2D{X: 7, Y: 23}}
	// visualization of result is here https://www.desmos.com/calculator/bv8pkx1r3q
	res := geometry.Point2D{X: 2, Y: 13}

	x, y := solve(segment1.Start, segment1.End, segment2.Start, segment2.End)

	assert.Equal(t, x, res.X)
	assert.Equal(t, y, res.Y)
}
