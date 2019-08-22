package maps

// Dictionary type map[string]string
type Dictionary map[string]string

// Search return the value of a key in a Dictionary
func (d Dictionary) Search(key string) (string, error) {
	definition, ok := d[key]
	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

// Add add a new word on the Dictionary
func (d Dictionary) Add(key string, value string) error {
	_, err := d.Search(key)
	switch err {
	case ErrNotFound:
		d[key] = value
	case nil:
		return ErrWordExists
	default:
		return err
	}
	return nil
}

// Update changes the value of a key
func (d Dictionary) Update(key string, value string) error {
	_, err := d.Search(key)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[key] = value
	default:
		return err
	}

	return nil
}

// Delete remove an item from the Dictionary
func (d Dictionary) Delete(key string) {
	delete(d, key)
}

var (
	// ErrNotFound returned when word can't be found
	ErrNotFound = DictionaryErr("Word not found!")
	// ErrWordExists returned when word exists
	ErrWordExists = DictionaryErr("Word already exists!")
	// ErrWordDoesNotExist returned when word doesn't exist
	ErrWordDoesNotExist = DictionaryErr("Cannot update. Word doesn't exist.")
)

// DictionaryErr type for errors on Dictionary
type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}
