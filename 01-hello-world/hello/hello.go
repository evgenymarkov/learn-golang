package hello

const (
	frenchLanguage  = "French"
	russianLanguage = "Russian"

	englishHelloPrefix = "Hello, "
	frenchHelloPrefix  = "Bonjour, "
	russianHelloPrefix = "Привет, "

	englishNameFallback = "world"
	frenchNameFallback  = "monde"
	russianNameFallback = "мир"
)

func Hello(name string, language string) string {
	prefix := getGreetingPrefix(language)

	if name == "" {
		name = getGreetingFallback(language)
	}

	return prefix + name
}

func getGreetingPrefix(language string) (prefix string) {
	switch language {
	case russianLanguage:
		prefix = russianHelloPrefix
	case frenchLanguage:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return
}

func getGreetingFallback(language string) (name string) {
	switch language {
	case russianLanguage:
		name = russianNameFallback
	case frenchLanguage:
		name = frenchNameFallback
	default:
		name = englishNameFallback
	}

	return
}
