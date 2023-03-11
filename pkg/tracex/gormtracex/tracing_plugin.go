package gormtracex

import (
	"context"
	"encoding/json"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"gorm.io/gorm"
	"time"
)

const (
	StartTimeKey      = "StartTime"
	DefaultSpanKey    = "gorm-tracing"
	DefaultPluginName = "GormTracingPlugin"

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
	_ gorm.Plugin = &TracingPlugin{}
)

type (
	Option       = func(t *TracingPlugin)
	TracingPlugin struct {
		name     string
		spanKey  string
		spanKind trace.SpanKind
	}
)

func NewTracingPlugin(opts ...Option) *TracingPlugin {
	plugin := &TracingPlugin{
		name:     DefaultPluginName,
		spanKey:  DefaultSpanKey,
		spanKind: trace.SpanKindClient,
	}
	for _, opt := range opts {
		opt(plugin)
	}
	return plugin
}

func WithTracingPluginName(name string) Option {
	return func(t *TracingPlugin) {
		t.name = name
	}
}

func WithTracingPluginSpanKey(spanKey string) Option {
	return func(t *TracingPlugin) {
		t.spanKey = spanKey
	}
}

func WithTracingPluginSpanKind(spanKind trace.SpanKind) Option {
	return func(t *TracingPlugin) {
		t.spanKind = spanKind
	}
}

func (p *TracingPlugin) Name() string {
	return p.name
}

func (p *TracingPlugin) Initialize(db *gorm.DB) (err error) {
	// 在 gorm 中注册各种回调事件
	for _, e := range []error{
		db.Callback().Create().Before(_ActionCreate).Register(_eventBeforeCreate, p.beforeCreate),
		db.Callback().Create().After(_ActionCreate).Register(_eventAfterCreate, p.after),
		db.Callback().Update().Before(_ActionUpdate).Register(_eventBeforeUpdate, p.beforeUpdate),
		db.Callback().Update().After(_ActionUpdate).Register(_eventAfterUpdate, p.after),
		db.Callback().Query().Before(_ActionQuery).Register(_eventBeforeQuery, p.beforeQuery),
		db.Callback().Query().After(_ActionQuery).Register(_eventAfterQuery, p.after),
		db.Callback().Delete().Before(_ActionDelete).Register(_eventBeforeDelete, p.beforeDelete),
		db.Callback().Delete().After(_ActionDelete).Register(_eventAfterDelete, p.after),
		db.Callback().Row().Before(_ActionRow).Register(_eventBeforeRow, p.beforeRow),
		db.Callback().Row().After(_ActionRow).Register(_eventAfterRow, p.after),
		db.Callback().Raw().Before(_ActionRaw).Register(_eventBeforeRaw, p.beforeRaw),
		db.Callback().Raw().After(_ActionRaw).Register(_eventAfterRaw, p.after),
	} {
		if e != nil {
			return e
		}
	}
	return
}
func (p *TracingPlugin) before(db *gorm.DB, op string) {
	if db.Statement == nil || db.Statement.Context == nil {
		db.Logger.Error(context.TODO(), "db.Statement and db.Statement.Context is nil")
		return
	}
	tracer := otel.Tracer(p.spanKey)
	_, span := tracer.Start(db.Statement.Context, op, trace.WithSpanKind(p.spanKind))
	//_, span := tracer.Start(db.Statement.Context, op)
	db.InstanceSet(p.spanKey, span).InstanceSet(StartTimeKey, time.Now())
}
func (p *TracingPlugin) after(db *gorm.DB) {
	if db.Statement == nil || db.Statement.Context == nil {
		db.Logger.Error(context.TODO(), "db.Statement and db.Statement.Context is nil")
		return
	}
	_span, ok := db.InstanceGet(p.spanKey)
	if !ok || _span == nil {
		return
	}
	span, ok := _span.(trace.Span)
	if !ok || span == nil {
		return
	}
	// execution time
	var executionTime time.Duration = 0
	_start, ok := db.InstanceGet(StartTimeKey)
	if ok {
		start, ok := _start.(time.Time)
		if ok {
			executionTime = time.Since(start)
		}
	}

	defer span.End()

	attrs := []attribute.KeyValue{
		attribute.String("Table", db.Statement.Table),
		attribute.String("Sql", db.Dialector.Explain(db.Statement.SQL.String(), db.Statement.Vars...)),
		attribute.String("Query", db.Statement.SQL.String()),
	}
	binds, err := json.Marshal(db.Statement.Vars)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
	} else {
		attrs = append(attrs, attribute.String("Bindings", string(binds)))
		if db.Error != nil {
			span.RecordError(db.Error)
			span.SetStatus(codes.Error, db.Error.Error())
		} else {
			attrs = append(attrs, attribute.Int64("AffectedRows", db.Statement.RowsAffected),
				attribute.Int64("ExecutionTime", executionTime.Milliseconds()))
			span.SetStatus(codes.Ok, "OK")
		}
	}
	span.SetAttributes(attrs...)
}

func (p *TracingPlugin) beforeCreate(db *gorm.DB) {
	p.before(db, _ActionCreate)
}
func (p *TracingPlugin) beforeUpdate(db *gorm.DB) {
	p.before(db, _ActionUpdate)
}
func (p *TracingPlugin) beforeQuery(db *gorm.DB) {
	p.before(db, _ActionQuery)
}
func (p *TracingPlugin) beforeDelete(db *gorm.DB) {
	p.before(db, _ActionDelete)
}
func (p *TracingPlugin) beforeRow(db *gorm.DB) {
	p.before(db, _ActionRow)
}
func (p *TracingPlugin) beforeRaw(db *gorm.DB) {
	p.before(db, _ActionRaw)
}
