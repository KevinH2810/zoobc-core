package query

import (
	"bytes"
	"fmt"
	"strings"
)

type (
	// CaseQueryInterface is interface includes `func` that handle Operand
	// CaseQueryInterface for complex query builder
	CaseQueryInterface interface {
		Select(tableName string, columns ...string) *CaseQuery
		Where(query ...string) *CaseQuery
		And(expression ...string) *CaseQuery
		Or(expression ...string) *CaseQuery
		In(column string, value ...interface{}) string
		NotIn(column string, value ...interface{}) string
		Equal(column string, value interface{}) string
		NotEqual(column string, value interface{}) string
		Between(column string, start, end interface{}) string
		NotBetween(column string, start, end interface{}) string
		Paginate(limit, currentPage uint32) *CaseQuery
		Build() (query string, args []interface{})
	}
	// CaseQuery would be as swap `Query` and `Args` those can save the query and args values
	// until `Done` called
	CaseQuery struct {
		Query *bytes.Buffer
		Args  []interface{}
	}
)

// Select build buffer query string
func (fq *CaseQuery) Select(tableName string, columns ...string) *CaseQuery {
	fq.Query.WriteString("SELECT ")
	if columns != nil {
		fq.Query.WriteString(strings.Join(columns, ", "))
	} else {
		fq.Query.WriteString("*")
	}
	fq.Query.WriteString(fmt.Sprintf(" FROM %s ", tableName))
	return fq
}

// Where build buffer query string, can combine with `In(), NotIn() ...`
func (fq *CaseQuery) Where(query ...string) *CaseQuery {
	fq.Query.WriteString(fmt.Sprintf(
		"WHERE %s ",
		strings.Join(query, ""),
	))
	return fq
}

// And represents `expressionFoo AND expressionBar`
func (fq *CaseQuery) And(query ...string) *CaseQuery {
	fq.Query.WriteString(fmt.Sprintf("AND %s", strings.Join(query, "AND ")))
	return fq
}

// Or represents `expressionFoo OR expressionBar`
func (fq *CaseQuery) Or(expression ...string) *CaseQuery {
	fq.Query.WriteString(fmt.Sprintf("OR %s ", strings.Join(expression, "OR ")))
	return fq
}

// In represents `column` IN (value...)
func (fq *CaseQuery) In(column string, value ...interface{}) string {
	fq.Args = append(fq.Args, value...)
	return fmt.Sprintf("%s IN (?%s) ", column, strings.Repeat(", ?", len(value)-1))
}

// NotIn represents `column NOT IN (value...)`
func (fq *CaseQuery) NotIn(column string, value ...interface{}) string {
	fq.Args = append(fq.Args, value...)
	return fmt.Sprintf("%s NOT IN (?%s) ", column, strings.Repeat(", ?", len(value)-1))
}

// Equal represents `column` == `value`
func (fq *CaseQuery) Equal(column string, value interface{}) string {
	fq.Args = append(fq.Args, value)
	return fmt.Sprintf("%s = ? ", column)
}

// NotEqual represents `column` != `value`
func (fq *CaseQuery) NotEqual(column string, value interface{}) string {
	fq.Args = append(fq.Args, value)
	return fmt.Sprintf("%s <> ? ", column)
}

// Between represents `column BETWEEN foo AND bar`
func (fq *CaseQuery) Between(column string, start, end interface{}) string {
	fq.Args = append(fq.Args, start, end)
	return fmt.Sprintf("%s BETWEEN ? AND ? ", column)
}

// NotBetween represents `column NOT BETWEEN foo AND bar`
func (fq *CaseQuery) NotBetween(column string, start, end interface{}) string {
	fq.Args = append(fq.Args, start, end)
	return fmt.Sprintf("%s NOT BETWEEN ? AND ? ", column)
}

// Paginate represents `limit = ? offset = ?`
// default limit = 30, page start from 1
func (fq *CaseQuery) Paginate(limit, currentPage uint32) *CaseQuery {
	if limit == 0 {
		limit = 30
	}
	if currentPage == 0 {
		currentPage = 1
	}
	offset := (currentPage - 1) * limit
	fq.Query.WriteString("limit ? offset ?")
	fq.Args = append(fq.Args, limit, offset)
	return fq
}

// Build should be called in the end of `CaseQuery` circular.
// And build buffer query string into string
func (fq *CaseQuery) Build() (query string, args []interface{}) {
	return fq.Query.String(), fq.Args
}