package parser

func Parse(raw []byte) []interface{} {
	return tengoParse(raw)
}
