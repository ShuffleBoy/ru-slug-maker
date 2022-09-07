package ruslug

import "strings"

var (
	transliterateRules = []string{
		"а", "a",
		"б", "b",
		"в", "v",
		"г", "g",
		"д", "d",
		"е", "e",
		"ё", "e",
		"ж", "zh",
		"з", "z",
		"и", "i",
		"й", "i",
		"к", "k",
		"л", "l",
		"м", "m",
		"н", "n",
		"о", "o",
		"п", "p",
		"р", "r",
		"с", "s",
		"т", "t",
		"у", "u",
		"ф", "f",
		"х", "kh",
		"ц", "ts",
		"х", "kh",
		"ч", "ch",
		"ш", "sh",
		"щ", "shch",
		"ъ", "ie",
		"ы", "y",
		"ь", "",
		"э", "e",
		"ю", "iu",
		"я", "ia",
	}

	transliterateReplacer = strings.NewReplacer(transliterateRules...)
	dashReplacer          = strings.NewReplacer(" ", "-")
)
