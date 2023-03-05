package sysuser

import (
	"kratos-admin/pkg/expr"
	"strings"
)

type (
	column struct {
		Asterisk expr.String
		Id       expr.String
		Uid      expr.String
		Username expr.String
		Password expr.String
		Realname expr.String
		Mobile   expr.String
		AreaCode expr.String
		Email    expr.String
		Weixin   expr.String
		Operator expr.String
		CreatedAt expr.String
		UpdatedAt expr.String
		DeletedAt expr.String

		slice []string
	}
	Option func(c *column)
)

var Column *column

func New(opts ...Option) *column {
	c := &column{
		Asterisk: "*",
		Id:       "id",
		Uid:      "uid",
		Username: "username",
		Password: "password",
		Realname: "realname",
		Mobile:   "mobile",
		AreaCode: "area_code",
		Email:    "email",
		Weixin:   "weixin",
		Operator: "operator",
		CreatedAt: "created_at",
		UpdatedAt: "updated_at",
		DeletedAt: "deleted_at",
	}
	c.slice = []string{
		c.Id.String(), c.Uid.String(), c.Username.String(), c.Password.String(), c.Realname.String(),
		c.Mobile.String(), c.AreaCode.String(), c.Email.String(), c.Weixin.String(),
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
	newSlice := make([]string,len(c.slice))
	copy( newSlice, c.slice)
	for _, v := range newSlice {
		v = "`"+v+"`"
	}
	return strings.Join(newSlice,",")
}


