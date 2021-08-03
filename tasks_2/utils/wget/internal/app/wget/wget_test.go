package wget

import (
	"net/http"
	"testing"
)

type FakeManager struct {
	url      string
	filePath string
}

func (f *FakeManager) Build() (string, error) {
	return f.filePath, nil
}
func (f *FakeManager) WriteResponse(res *http.Response) (int64, error) {
	return 0, nil
}

func TestWget_Parse(t *testing.T) {}
