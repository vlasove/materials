package mansort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type finisher struct {
}

func (f *finisher) execute(m *ManSort) {
}

func (f *finisher) setNext(next handler) {
}

func TestMansort_KeyColumnSorterHandler(t *testing.T) {

	ms := &ManSort{
		options: Options{
			keyColumnNum: -1,
		},
		data: []string{"bob fisher", "alex messar", "gordon friman"},
	}
	mockFinish := &finisher{}
	columnSort := &keyColumnSorter{}
	columnSort.setNext(mockFinish)

	testCases := []struct {
		name     string
		colNum   int
		expected []string
	}{
		{
			name:     "sort by 0 column",
			colNum:   0,
			expected: []string{"alex messar", "bob fisher", "gordon friman"},
		},
		{
			name:     "test deafult",
			colNum:   -1,
			expected: []string{"alex messar", "bob fisher", "gordon friman"},
		},
		{
			name:     "sort by 1 column",
			colNum:   1,
			expected: []string{"bob fisher", "gordon friman", "alex messar"},
		},
		{
			name:     "sort by not existing column",
			colNum:   10,
			expected: []string{"alex messar", "bob fisher", "gordon friman"},
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			ms.keyColumnSortDone = false
			ms.options.keyColumnNum = test.colNum
			columnSort.execute(ms)
			assert.Equal(t, ms.data, test.expected)
		})
	}

}
func TestMansort_ReverseDataHandler(t *testing.T) {
	ms := &ManSort{
		options: Options{
			reverseNeeded: true,
		},
		data: []string{"bob fisher", "alex messar", "gordon friman"},
	}
	mockFinish := &finisher{}
	reverseSorter := &reverseDataSorter{}
	reverseSorter.setNext(mockFinish)

	testCases := []struct {
		name     string
		expected []string
	}{
		{
			name:     "regular reversing",
			expected: []string{"gordon friman", "alex messar", "bob fisher"},
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			reverseSorter.execute(ms)
			assert.Equal(t, ms.data, test.expected)
		})
	}

}
func TestMansort_UniqueSanitizerHandler(t *testing.T) {
	ms := &ManSort{
		options: Options{
			onlyUnique: true,
		},
		data: []string{"bob fisher", "alex messar", "gordon friman", "bob fisher", "alex messar"},
	}
	mockFinish := &finisher{}
	uniqueHandler := &uniqueSanitizer{}
	uniqueHandler.setNext(mockFinish)

	testCases := []struct {
		name     string
		expected []string
	}{
		{
			name:     "regular sanitizing",
			expected: []string{"bob fisher", "alex messar", "gordon friman"},
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			uniqueHandler.execute(ms)
			assert.Equal(t, ms.data, test.expected)
		})
	}
}
func TestMansort_AlreadySortedCheckerHandler(t *testing.T) {

}
func TestMansort_TailsCheckerHandler(t *testing.T)   {}
func TestMansort_NumColSorterHandler(t *testing.T)   {}
func TestMansort_MonthColSorterHandler(t *testing.T) {}
