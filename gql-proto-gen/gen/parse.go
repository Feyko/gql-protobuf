package gen

import (
	"github.com/graphql-go/graphql/language/ast"
	"github.com/graphql-go/graphql/language/parser"
	"github.com/pkg/errors"
)

func Parse(content string) (*ast.Document, error) {
	params := parser.ParseParams{
		Source: content,
		Options: parser.ParseOptions{
			NoLocation: false,
			NoSource:   false,
		},
	}

	doc, err := parser.Parse(params)
	return doc, errors.Wrap(err, "could not parse the document")
}
