package compiler

type Location struct {
	Line   uint
	Column uint
}

type TokenType uint

type Token struct {
	Type     TokenType
	Lexeme   string
	Location Location
}

type Error struct {
	Message  string
	Location Location
}
