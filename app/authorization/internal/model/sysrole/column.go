package sysrole

import (
	"kratos-admin/pkg/expr"
	"strings"
)

type (
	column struct {
		Asterisk    expr.String
		Id          expr.String
		Name        expr.String
		Description expr.String
		IsEnable    expr.String
		Operator    expr.String
		CreatedAt   expr.String
		UpdatedAt   expr.String
		DeletedAt   expr.String

		slice []string
	}
	Option func(c *column)
)

var Column *column

func New(opts ...Option) *column {
	c := &column{
		Asterisk:    "*",
		Id:          "id",
		Name:        "name",
		Description: "description",
		IsEnable:    "is_enable",
		Operator:    "operator",
		CreatedAt:   "created_at",
		UpdatedAt:   "updated_at",
		DeletedAt:   "deleted_at",
	}
	c.slice = []string{
		c.Id.String(), c.Name.String(), c.Description.String(), c.IsEnable.String(),
		c.Operator.String(), c.CreatedAt.String(), c.UpdatedAt.String(), c.DeletedAt.String(),
	}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

func (c *column) Slice() []string {
	return c.slice
}

func (c *column) All() string {
	newSlice := make([]string, len(c.slice))
	copy(newSlice, c.slice)
	for _, v := range newSlice {
		v = "`" + v + "`"
	}
	return strings.Join(newSlice, ",")
}
