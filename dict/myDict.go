package myDict

import "errors"

// Dictionary type
type Dictionary map[string]string

// var errorNotFound error = errors.New("not found error")
// var errorAlreadyContent error = errors.New("already content error")
// var errorCantUpdate error = errors.New("can't update no exists word")
var (
	errorNotFound       error = errors.New("not found error")
	errorAlreadyContent error = errors.New("already content error")
	errorCantUpdate     error = errors.New("can't update no exists word")
)

func (d Dictionary) Search(searchKeyWord string) (string, error) {
	value, exists := d[searchKeyWord]
	if exists {
		return value, nil
	}
	return "", errorNotFound
}

// Add new contents Add
func (d Dictionary) Add(key, value string) error {
	_, error := d.Search(key)
	switch error {
	case errorNotFound:
		d[key] = value
	case nil:
		return errorAlreadyContent
	}
	return nil
}

// Update if a word exists to the Dictionary
func (d Dictionary) Update(key, value string) error {
	_, error := d.Search(key)
	switch error {
	case nil:
		d[key] = value
	case errorNotFound:
		return errorCantUpdate
	}
	return nil
}

// Delete a word to the Dictionary
func (d Dictionary) Delete(key string) {
	delete(d, key)
}
