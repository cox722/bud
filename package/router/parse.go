package router

import "github.com/cox722/go-fullstack-cox/package/router/lex"

func Parse(route string) (tokens []lex.Token) {
	lexer := lex.New(route)
	for token := lexer.Next(); token.Type != lex.EndToken; {
		tokens = append(tokens, token)
	}
	return tokens
}
