package postgres

import (
	"context"
	"database/sql"
	"github.com/jmoiron/sqlx"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

// TracedDB оборачивает sqlx.DB для добавления трассировки
type TracedDB struct {
	*sqlx.DB
	tracer trace.Tracer
}

// NewTracedDB создает новый TracedDB
func NewTracedDB(db *sqlx.DB, tracer trace.Tracer) *TracedDB {
	return &TracedDB{
		DB:     db,
		tracer: tracer,
	}
}

// QueryContext оборачивает sqlx.QueryContext с трассировкой
func (t *TracedDB) QueryContext(ctx context.Context, query string, args ...interface{}) (*sqlx.Rows, error) {
	// Начинаем новый span для SQL-запроса
	ctx, span := t.tracer.Start(ctx, "QueryContext")
	defer span.End()

	// Добавляем атрибуты в span
	span.SetAttributes(
		attribute.String("db.system", "postgresql"),
		attribute.String("db.statement", query),
	)

	// Выполняем запрос
	return t.DB.QueryxContext(ctx, query, args...)
}

// ExecContext оборачивает sqlx.ExecContext с трассировкой
func (t *TracedDB) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	// Начинаем новый span для SQL-запроса
	ctx, span := t.tracer.Start(ctx, "ExecContext")
	defer span.End()

	// Добавляем атрибуты в span
	span.SetAttributes(
		attribute.String("db.system", "postgresql"),
		attribute.String("db.statement", query),
	)

	// Выполняем запрос
	return t.DB.ExecContext(ctx, query, args...)
}

// GetContext оборачивает sqlx.GetContext с трассировкой
func (t *TracedDB) GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	// Начинаем новый span для SQL-запроса
	ctx, span := t.tracer.Start(ctx, "GetContext")
	defer span.End()

	// Добавляем атрибуты в span
	span.SetAttributes(
		attribute.String("db.system", "postgresql"),
		attribute.String("db.statement", query),
	)

	// Выполняем запрос
	return t.DB.GetContext(ctx, dest, query, args...)
}

// Scan оборачивает sqlx.Scan с трассировкой
func (t *TracedDB) Scan(ctx context.Context, rows *sqlx.Rows, dest ...interface{}) error {
	// Начинаем новый span для Scan
	ctx, span := t.tracer.Start(ctx, "Scan")
	defer span.End()

	// Добавляем атрибуты в span
	span.SetAttributes(
		attribute.String("db.system", "postgresql"),
	)

	// Выполняем сканирование
	if err := rows.Scan(dest...); err != nil {
		span.RecordError(err)
		return err
	}
	return nil
}

// QueryRowContext оборачивает sqlx.QueryRowContext с трассировкой
func (t *TracedDB) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sqlx.Row {
	// Начинаем новый span для SQL-запроса
	ctx, span := t.tracer.Start(ctx, "QueryRowContext")
	defer span.End()

	// Добавляем атрибуты в span
	span.SetAttributes(
		attribute.String("db.system", "postgresql"),
		attribute.String("db.statement", query),
	)

	// Выполняем запрос
	return t.DB.QueryRowxContext(ctx, query, args...)
}
