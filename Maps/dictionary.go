package maps

import "errors"

type Dictionary map[string]string

var ErrNotFound = errors.New("word not found")

func (d Dictionary) Search (word string) (string, error) {
	value, ok := d[word]

	if !ok {
		return "", ErrNotFound
	}

	return value, nil
}