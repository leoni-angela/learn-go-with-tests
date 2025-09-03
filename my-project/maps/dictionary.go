package maps



type Dictionary map[string]string
// maps you do not need to send reference of it
// map value is a pointer to a runtime.hmap structure.
// when passing maps, you are still copy, but just pointer not underlying
// maps can be of nil value, and when attempt to write to a nil map, it will panic
// therefore, NEVER initialise a nil map variable

const (
	ErrNotFound = DictionaryErr("could not find the word you were looking for")
	ErrWordExists = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("cannot perform operation on word because it does not exist")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}
// makes errors more reusable and immutable

func (d Dictionary) Search(word string) (string, error) {

	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

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
	// using switch as a guard incase Search sends errors we have not considered
	return nil
}

func (d Dictionary) Update(word, definition string) error{

	_, err := d.Search(word)
	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = definition
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(word string) error {

	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		delete(d, word)
	default:
		return err
	}
	return nil
}