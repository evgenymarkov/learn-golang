package dictionary

type Dictionary map[string]string

func (dictionary Dictionary) Add(word string, definition string) error {
	_, err := dictionary.Search(word)

	switch err {
	case ErrNoWordFound:
		dictionary[word] = definition
	case nil:
		return ErrWordExist
	default:
		return err
	}

	return nil
}

func (dictionary Dictionary) Search(word string) (string, error) {
	value, found := dictionary[word]

	if !found {
		return "", ErrNoWordFound
	}

	return value, nil
}

func (dictionary Dictionary) Update(word string, definition string) error {
	_, err := dictionary.Search(word)

	switch err {
	case ErrNoWordFound:
		return ErrWordDoesNotExist
	case nil:
		dictionary[word] = definition
	default:
		return err
	}

	return nil
}

func (dictionary Dictionary) Delete(word string) {
	delete(dictionary, word)
}
