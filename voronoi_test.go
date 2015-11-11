package main

import (
	"image/color"
	"testing"
)

func TestNumCentroids(t *testing.T) {
	centroids := createCentroids(10, 10, 10)
	if len(centroids) != 10 {
		t.Errorf("number of centroids should be %d but was %d", 10, len(centroids))
	}
}

func TestCentroidMustBeInsideImage(t *testing.T) {
	width := 100
	height := 100
	centroid := createCentroids(1, width, height)[0]
	if centroid.x < 0 || centroid.x > width {
		t.Errorf("centroid.x was %d", centroid.x)
	}
	if centroid.y < 0 || centroid.y > height {
		t.Errorf("centroid.y was %d", centroid.y)
	}
}

func TestCentroidMustHaveColor(t *testing.T) {
	c := createCentroids(1, 100, 100)[0]
	if c.color == nil {
		t.Error("centroid must have color")
	}
}

func TestPointMustHaveTheColorOfNearestCentroid(t *testing.T) {
	centroids := []*point{
		&point{1, 1, color.Black},
		&point{100, 100, color.White},
	}

	point := point{x: 5, y: 5}
	point.color = nearestCentroidColor(point, centroids)
	if point.color != color.Black {
		t.Errorf("color of point should be %v but was %v", color.Black, point.color)
	}
}
