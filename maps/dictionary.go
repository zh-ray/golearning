package maps

// Errors
const (
	ErrNotFound      = DictionaryErr("could not find the word you were looking for")
	ErrWordExists    = DictionaryErr("could not add word because it already exists")
	ErrWordNotUpdate = DictionaryErr("cannot update word because it does not exist")
)

// DictionaryErr is the error string type
type DictionaryErr string

func (d DictionaryErr) Error() string {
	return string(d)
}

// Dictionary is a type
type Dictionary map[string]string

// Search return the value in map accound to the word
func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

// Add receive a word and the definition of the word, then add them into the dictionary
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

// Update the definition in dictionary according to the word
func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case ErrNotFound:
		return ErrWordNotUpdate
	case nil:
		d[word] = definition
	default:
		return err
	}

	return nil
}

// Delete the definition in dictionary according to the word
func (d Dictionary) Delete(word string) {
	delete(d, word)
}
