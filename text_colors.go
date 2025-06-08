package main

// helpers to make text stand out using colors
//

// Returns the string 'Error' in red
func FormattedErrorText() string {
	return "\x1b[31mError\x1b[0m"
}
