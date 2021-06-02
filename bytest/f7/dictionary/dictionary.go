package dictionary

const (
	// ErrNotFound ...
	ErrNotFound = DictionaryErr("could not find the word you were looking for")
	// ErrWordExists ...
	ErrWordExists = DictionaryErr("cannot add word because it already exists")
	// ErrWordNotExists ...
	ErrWordNotExists = DictionaryErr("can not update/delete word because not exists")
)

// DictionaryErr ...
type DictionaryErr string

// Error ...
func (d DictionaryErr) Error() string {
	return string(d)
}

// Dictionary ...
type Dictionary map[string]string

// Search ...
func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

// Add ...
func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}
	return nil
}

// Update ...
func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case ErrNotFound:
		return ErrWordNotExists
	case nil:
		d[word] = definition
	default:
		return err
	}
	return nil
}

// Delete ...
func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)
	switch err {
	case ErrNotFound:
		return ErrWordNotExists
	case nil:
		delete(d, word)
	default:
		return err
	}
	return nil
}
