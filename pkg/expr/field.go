package expr

var (
	Symbol String = "?"
)

func (f String) Tab(val string) string {
	return f.expr(func(f String) String {
		f = String(val) + ". `" + f + "`"
		return f
	}).toString()
}

func (f String) As(val string) string {
	return f.expr(func(f String) String {
		f = "`" + f + "` as " + String(val)
		return f
	}).toString()
}
func (f String) Eq() string {
	return f.expr(func(f String) String {
		f = "`" + f + "` = " + Symbol
		return f
	}).toString()
}
func (f String) Neq() string {
	return f.expr(func(f String) String {
		f = "`" + f + "` != " + Symbol
		return f
	}).toString()
}
func (f String) Gt() string {
	return f.expr(func(f String) String {
		f = "`" + f + "` > " + Symbol
		return f
	}).toString()
}
func (f String) Gte() string {
	return f.expr(func(f String) String {
		f = "`" + f + "` >= " + Symbol
		return f
	}).toString()
}
func (f String) Lt() string {
	return f.expr(func(f String) String {
		f = "`" + f + "` < " + Symbol
		return f
	}).toString()
}
func (f String) Lte() string {
	return f.expr(func(f String) String {
		f = "`" + f + "` <= " + Symbol
		return f
	}).toString()
}
func (f String) In() string {
	return f.expr(func(f String) String {
		//str := String(strings.TrimRight(strings.Repeat(string(Symbol+","), n), ","))
		f = "`" + f + "` IN (?)"
		return f
	}).toString()
}
func (f String) NotIn() string {
	return f.expr(func(f String) String {
		//str := String(strings.TrimRight(strings.Repeat(string(Symbol+","), n), ","))
		f = "`" + f + "` NOT IN (?)"
		return f
	}).toString()
}
func (f String) Between() string {
	return f.expr(func(f String) String {
		f = "`" + f + "` BETWEEN " + Symbol + " AND " + Symbol
		return f
	}).toString()
}
func (f String) NotBetween() string {
	return f.expr(func(f String) String {
		f = "`" + f + "` NOT BETWEEN " + Symbol + " AND " + Symbol
		return f
	}).toString()
}
func (f String) IsNULL() string {
	return f.expr(func(f String) String {
		f = "`" + f + "` IS NULL"
		return f
	}).toString()
}
func (f String) NotNULL() string {
	return f.expr(func(f String) String {
		f = "`" + f + "` NOT IS NULL"
		return f
	}).toString()
}
func (f String) Like() string {
	return f.expr(func(f String) String {
		f = "`" + f + "` LIKE ?"
		return f
	}).toString()
}
func (f String) NotLike() string {
	return f.expr(func(f String) String {
		f = "`" + f + "` NOT LIKE ?"
		return f
	}).toString()
}
func (f String) IfNULL() string {
	return f.expr(func(f String) String {
		f = "IFNULL(`" + f + "`, ?)"
		return f
	}).toString()
}

func (f String) order(asc bool) string {
	return f.expr(func(f String) String {
		if asc {
			f = "`" + f + "` ASC"
		} else {
			f = "`" + f + "` DESC"
		}
		return f
	}).toString()
}

func (f String) OrderAsc() string {
	return f.order(true)
}

func (f String) OrderDesc() string {
	return f.order(false)
}
