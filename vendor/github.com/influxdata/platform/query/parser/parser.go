// +build !parser_debug

package parser

//go:generate pigeon -optimize-parser -optimize-grammar -o ifql.go ifql.peg

import (
	"github.com/influxdata/platform/query/ast"
)

// NewAST parses ifql query and produces an ast.Program
func NewAST(ifql string, opts ...Option) (*ast.Program, error) {
	f, err := Parse("", []byte(ifql), opts...)
	if err != nil {
		return nil, err
	}
	return f.(*ast.Program), nil
}
