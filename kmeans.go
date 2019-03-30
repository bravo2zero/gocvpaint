package main

import (
	"gocv.io/x/gocv"
)

type KmeansConfig struct {
	ClusterCount  int
	MaxIterations int
	Delta         float64
}

// Calculate the clusters centroids using k-means clustering method
func KMeans(source gocv.Mat, config *KmeansConfig) (labels []int, centroids []Color) {

	if config == nil {
		config = &KmeansConfig{ClusterCount: 12, MaxIterations: 500, Delta: .001}
	}

	// TODO

	return
}
