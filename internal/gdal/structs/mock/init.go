package gdal_mock

import "gogdal/internal/config"

type MockGdalWorker struct {
	iter uint
}

func NewMockGdalWorker(conf *config.Config) *MockGdalWorker {
	return &MockGdalWorker{iter: 0}
}

func (w *MockGdalWorker) IntersectPolygons(polys ...string) (float64, bool, error) {
	defer func() {
		w.iter++
	}()
	if w.iter%2 == 0 {
		return float64(w.iter), true, nil
	} else {
		return float64(w.iter), false, nil
	}
}
