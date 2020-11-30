package main

import "errors"

var ErrNotFound = errors.New("word not found")

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	def, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return def, nil
}

func (d Dictionary) Add(word, def string) {
	d[word] = def
}
