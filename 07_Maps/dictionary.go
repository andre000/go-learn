package maps

type Dictionary map[string]string

func (d Dictionary) Search(key string) (string, error) {
	definition, ok := d[key]
	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

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

func (d Dictionary) Update(key string, value string) {
	d[key] = value
}

var (
	ErrNotFound   = DictionaryErr("Word not found!")
	ErrWordExists = DictionaryErr("Word already exists!")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}
