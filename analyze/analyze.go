package analyze

import (
	"os"
)

type Results struct {
	Words       int `json:"words"`
	Digits      int `json:"digits"`
	SpecialChar int `json:"special_char"`
	Lines       int `json:"lines"`
	Spaces      int `json:"spaces"`
	Sentences   int `json:"sentences"`
	Punctuation int `json:"punctuation"`
	Consonants  int `json:"consonants"`
	Vowels      int `json:"vowels"`
}

func Analyze(a string) Results {
	results := Results{}

	for i := 0; i < len(a); i++ {
		if a[i] == ' ' {
			results.Words++
			results.Spaces++
		} else if a[i] == '\n' {
			results.Words++
			results.Lines++
			results.Sentences++
		} else if a[i] == '\t' {
			results.Words++
		} else if a[i] == '.' || a[i] == ',' || a[i] == '!' || a[i] == '?' || a[i] == ';' || a[i] == ':' || a[i] == ')' || a[i] == '(' || a[i] == '-' || a[i] == '_' || a[i] == '"' || a[i] == '\'' {
			results.Punctuation++
		} else if a[i] == 'a' || a[i] == 'e' || a[i] == 'i' || a[i] == 'o' || a[i] == 'u' || a[i] == 'A' || a[i] == 'E' || a[i] == 'I' || a[i] == 'O' || a[i] == 'U' {
			results.Vowels++
		} else if a[i] >= '0' && a[i] <= '9' {
			results.Digits++
		} else if a[i] == '@' || a[i] == '#' || a[i] == '$' || a[i] == '%' || a[i] == '^' || a[i] == '&' || a[i] == '*' || a[i] == '+' || a[i] == '=' || a[i] == '{' || a[i] == '}' || a[i] == '[' || a[i] == ']' {
			results.SpecialChar++
		} else if (a[i] >= 'a' && a[i] <= 'z' || a[i] >= 'A' && a[i] <= 'Z') && !(a[i] == 'a' || a[i] == 'e' || a[i] == 'i' || a[i] == 'o' || a[i] == 'u' || a[i] == 'A' || a[i] == 'E' || a[i] == 'I' || a[i] == 'O' || a[i] == 'U') {
			results.Consonants++
		}
	}

	// Adjust for the last line and paragraph
	results.Lines++
	results.Words++

	return results
}

func AnalyzeFile(filepath string) (Results, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return Results{}, err
	}
	return Analyze(string(data)), nil
}
