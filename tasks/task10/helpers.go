package task10

// min ...
func min(s []float64) (float64, error) {
	if len(s) == 0 {
		return 0, errZeroLength
	}
	m := s[0]
	for _, v := range s {
		if v < m {
			m = v
		}
	}
	return m, nil
}

// max ...
func max(s []float64) (float64, error) {
	if len(s) == 0 {
		return 0, errZeroLength
	}
	m := s[0]
	for _, v := range s {
		if v > m {
			m = v
		}
	}
	return m, nil
}

// borders ...
func borders(min, max float64) (int, int) {
	minBord := int(min) / step
	maxBord := int(max) / step

	minBord--
	maxBord++

	return minBord * step, maxBord * step

}

// buildChunk ...
func buildChunck(minB, maxB int) map[int][]float64 {
	m := make(map[int][]float64)
	for i := minB + step; i <= maxB; i += step {
		m[i] = []float64{}
	}
	return m
}

// temperatireChecked ...
type temperatureChecked struct {
	t         float64
	isChecked bool
}

// fillChunk ...
func fillChunk(temps []float64, m map[int][]float64) map[int][]float64 {
	tempsChecked := []*temperatureChecked{}
	for _, t := range temperatures {
		tempsChecked = append(tempsChecked, &temperatureChecked{t, false})
	}

	for k := range m {
		for _, t := range tempsChecked {
			if t.t <= float64(k) && !t.isChecked {
				m[k] = append(m[k], t.t)
				t.isChecked = true
			}
		}
	}
	return sanitize(m)
}

// sanitize ...
func sanitize(m map[int][]float64) map[int][]float64 {
	keysToDelete := []int{}
	for k, v := range m {
		if len(v) == 0 {
			keysToDelete = append(keysToDelete, k)
		}
	}

	for _, k := range keysToDelete {
		delete(m, k)
	}
	return m
}
