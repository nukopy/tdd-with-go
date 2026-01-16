package greeting

const commonHelloSuffix = "!"

// English
const englishHelloPrefix = "Hello, "

// Spanish
const languageSpanish = "Spanish"
const spanishHelloPrefix = "Hola, "

// French
const languageFrench = "French"
const frenchHelloPrefix = "Bonjour, "

func getGreetingPrefix(language string) string {
	prefix := ""
	switch language {
	case languageFrench:
		prefix = frenchHelloPrefix
	case languageSpanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return prefix
}

func Hello(name, language string) string {
	if name == "" {
		name = "world"
	}

	prefix := getGreetingPrefix(language)

	return prefix + name + commonHelloSuffix
}
