package main

// import "errors"

type Dictionary map[string]string

const (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrWordExists       = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("The word doesnt exist")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(key string) (string, error) {
	value, found := d[key] //maps can return two values, the second one if whether or not the key was found.

	if !found {
		return "", ErrNotFound
	}

	return value, nil
}

// An interesting property of maps is that you can modify them without passing them as a pointer.
// This is because map is a reference type.
// Meaning it holds a reference to the underlying data structure, much like a pointer.
// The underlying data structure is a hash table, or hash map, and you can read more about hash tables
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

func (d Dictionary) Update(word, newDefinition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = newDefinition
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
