package gormplugin

import (
	"context"
	"encoding/json"
	"fmt"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"gorm.io/gorm"
)

const (
	spanKey = "gorm-tracing"

	// 自定义事件名称
	_eventBeforeCreate = "gorm-tracing-event:before_create"
	_eventAfterCreate  = "gorm-tracing-event:after_create"
	_eventBeforeUpdate = "gorm-tracing-event:before_update"
	_eventAfterUpdate  = "gorm-tracing-event:after_update"
	_eventBeforeQuery  = "gorm-tracing-event:before_query"
	_eventAfterQuery   = "gorm-tracing-event:after_query"
	_eventBeforeDelete = "gorm-tracing-event:before_delete"
	_eventAfterDelete  = "gorm-tracing-event:after_delete"
	_eventBeforeRow    = "gorm-tracing-event:before_row"
	_eventAfterRow     = "gorm-tracing-event:after_row"
	_eventBeforeRaw    = "gorm-tracing-event:before_raw"
	_eventAfterRaw     = "gorm-tracing-event:after_raw"

	// 自定义span的操作名称
	_ActionCreate = "gorm:create"
	_ActionUpdate = "gorm:update"
	_ActionQuery  = "gorm:query"
	_ActionDelete = "gorm:delete"
	_ActionRow    = "gorm:row"
	_ActionRaw    = "gorm:raw"
)

var (
	_ gorm.Plugin = &GormPlugin{}
)

type GormPlugin struct {
}

func (gp *GormPlugin) Name() string {
	return "GormTracingPlugin"
}

func (gp *GormPlugin) Initialize(db *gorm.DB) (err error) {
	// 在 gorm 中注册各种回调事件
	for _, e := range []error{
		db.Callback().Create().Before(_ActionCreate).Register(_eventBeforeCreate, beforeCreate),
		db.Callback().Create().After(_ActionCreate).Register(_eventAfterCreate, after),
		db.Callback().Update().Before(_ActionUpdate).Register(_eventBeforeUpdate, beforeUpdate),
		db.Callback().Update().After(_ActionUpdate).Register(_eventAfterUpdate, after),
		db.Callback().Query().Before(_ActionQuery).Register(_eventBeforeQuery, beforeQuery),
		db.Callback().Query().After(_ActionQuery).Register(_eventAfterQuery, after),
		db.Callback().Delete().Before(_ActionDelete).Register(_eventBeforeDelete, beforeDelete),
		db.Callback().Delete().After(_ActionDelete).Register(_eventAfterDelete, after),
		db.Callback().Row().Before(_ActionRow).Register(_eventBeforeRow, beforeRow),
		db.Callback().Row().After(_ActionRow).Register(_eventAfterRow, after),
		db.Callback().Raw().Before(_ActionRaw).Register(_eventBeforeRaw, beforeRaw),
		db.Callback().Raw().After(_ActionRaw).Register(_eventAfterRaw, after),
	} {
		if e != nil {
			return e
		}
	}
	return
}
func before(db *gorm.DB, op string) {
	if db.Statement == nil || db.Statement.Context == nil {
		db.Logger.Error(context.TODO(), "db.Statement and db.Statement.Context is nil")
		return
	}
	tracer := otel.Tracer(spanKey)
	//ctx, span := tracer.Start(db.Statement.Context, op, trace.WithSpanKind(trace.SpanKindServer))
	_, span := tracer.Start(db.Statement.Context, op)
	//db.WithContext(ctx).InstanceSet(spanKey, span)
	db.InstanceSet(spanKey, span)
}
func after(db *gorm.DB) {
	if db.Statement == nil || db.Statement.Context == nil {
		db.Logger.Error(context.TODO(), "db.Statement and db.Statement.Context is nil")
		return
	}
	_span, ok := db.InstanceGet(spanKey)
	if !ok || _span == nil {
		return
	}
	span, ok := _span.(trace.Span)
	if !ok || span == nil {
		return
	}
	//tracer := otel.Tracer(spanKey)
	//_, span := tracer.Start(db.Statement.Context, _ActionCreate, trace.WithSpanKind(trace.SpanKindServer))
	defer span.End()
	attrs := []attribute.KeyValue{
		attribute.String("table", db.Statement.Table),
		attribute.String("sql", db.Dialector.Explain(db.Statement.SQL.String(), db.Statement.Vars...)),
		attribute.String("query", db.Statement.SQL.String()),
	}
	binds, err := json.Marshal(db.Statement.Vars)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
	} else {
		span.SetStatus(codes.Ok, "OK")
		attrs = append(attrs, attribute.String("bindings", string(binds)))
	}
	fmt.Println(attrs)
	span.SetAttributes(attrs...)

}

func beforeCreate(db *gorm.DB) {
	before(db, _ActionCreate)
}
func beforeUpdate(db *gorm.DB) {
	before(db, _ActionUpdate)
}
func beforeQuery(db *gorm.DB) {
	before(db, _ActionQuery)
}
func beforeDelete(db *gorm.DB) {
	before(db, _ActionDelete)
}
func beforeRow(db *gorm.DB) {
	before(db, _ActionRow)
}
func beforeRaw(db *gorm.DB) {
	before(db, _ActionRaw)
}
