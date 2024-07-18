
package parser

import (
	"github.com/vektah/gqlparser/v2/lexer"

	//nolint:revive
	. "github.com/vektah/gqlparser/v2/ast"
)

func ParseQuery(source *Source) (*QueryDocument, error) {
	p := parser{
		lexer:         lexer.New(source),
		maxTokenLimit: 0, // 0 means unlimited
	}
	return p.parseQueryDocument(), p.err
}

func ParseQueryWithTokenLimit(source *Source, maxTokenLimit int) (*QueryDocument, error) {
	p := parser{
		lexer:         lexer.New(source),
		maxTokenLimit: maxTokenLimit,
	}
	return p.parseQueryDocument(), p.err
}

// ...other existing functions...
