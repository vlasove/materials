package wget

import (
	"errors"
	"net/http"

	"github.com/vlasove/materials/tasks_2/utils/wget/internal/app/managers"
)

var (
	errManagerBuild = errors.New("wget: errors when creating directory")
	errHTTPGetBad   = errors.New("wget: can not make get request")
)

// WGet ...
type WGet struct {
	manager *managers.DirectoryManager
}

// New ...
func New(manager *managers.DirectoryManager) *WGet {
	return &WGet{
		manager: manager,
	}
}

// Parse ...
func (w *WGet) Parse() (*http.Response, error) {
	_, err := w.manager.Build()
	if err != nil {
		return nil, errManagerBuild
	}
	client := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}

	res, err := client.Get(w.manager.BaseURL)
	if err != nil {
		return nil, errHTTPGetBad
	}

	return res, nil
}
