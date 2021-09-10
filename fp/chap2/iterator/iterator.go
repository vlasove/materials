package iterator

const (
	InvalidIntVal   = -1
	InvaidStringVal = ""
)

type CarIterator interface {
	Next() (string, bool)
}

type Collection struct {
	Index int
	List  []string
}

func NewCollection(s []string) *Collection {
	return &Collection{
		Index: InvalidIntVal,
		List:  s,
	}
}

func (c *Collection) Next() (string, bool) {
	c.Index++
	if c.Index >= len(c.List) {
		return InvaidStringVal, false
	}
	return c.List[c.Index], true
}
