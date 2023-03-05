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
		Unionid  expr.String
		CreateAt expr.String
		UpdateAt expr.String
		Operator expr.String
		IsDelete expr.String
		DeleteAt expr.String

		slice []string
		//useQuote bool // 是否使用反引号
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
		Unionid:  "unionid",
		CreateAt: "create_at",
		UpdateAt: "update_at",
		Operator: "operator",
		IsDelete: "is_delete",
		DeleteAt: "delete_at",
	}
	c.slice = []string{
		c.Id.String(), c.Uid.String(), c.Username.String(), c.Password.String(), c.Realname.String(),
		c.Mobile.String(), c.AreaCode.String(), c.Email.String(), c.Weixin.String(), c.Unionid.String(),
		c.CreateAt.String(), c.UpdateAt.String(), c.Operator.String(), c.IsDelete.String(), c.DeleteAt.String(),
	}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

//func WithUseQuote() Option  {
//	return func(c *column) {
//		c.useQuote = true
//	}
//}

func (c *column) Slice() []string {
	return c.slice
}

func (c *column) All() string {
	//if c.useQuote {
	//	return strings.Join(c.slice,",")
	//}
	newSlice := make([]string,len(c.slice))
	copy( newSlice, c.slice)
	for _, v := range newSlice {
		v = "`"+v+"`"
	}
	return strings.Join(newSlice,",")
}


