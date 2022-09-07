package ruslug

import (
	"regexp"
	"strings"
)

var (
	compatibleChars = regexp.MustCompile("[а-яёa-z0-9 -]")
	wideSpace       = regexp.MustCompile("( +)")
)

func toLower(source string) string {
	return strings.ToLower(source)
}

func trim(source string) string {
	return strings.TrimSpace(source)
}

func escapeUnCompatibleChars(source string) string {
	var result strings.Builder
	for _, char := range source {
		if compatibleChars.MatchString(string(char)) {
			result.WriteRune(char)
		}
	}
	return result.String()
}

func removeWideSpaces(source string) string {
	return wideSpace.ReplaceAllString(source, " ")
}

func transliterate(source string) string {
	return transliterateReplacer.Replace(source)
}

func fillWithDashes(source string) string {
	return dashReplacer.Replace(source)
}

func Generate(source string) string {
	content := toLower(source)
	content = trim(content)
	content = escapeUnCompatibleChars(content)
	content = removeWideSpaces(content)
	content = transliterate(content)
	content = fillWithDashes(content)
	return content
}
