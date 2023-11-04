package dictionary

type DictionaryErr string

const (
	ErrWordExist        = DictionaryErr("cannot add word because it already exists")
	ErrNoWordFound      = DictionaryErr("could not find the word you were looking for")
	ErrWordDoesNotExist = DictionaryErr("cannot update word because it does not exist")
)

func (err DictionaryErr) Error() string {
	return string(err)
}
