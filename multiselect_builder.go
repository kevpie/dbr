package dbr

import "context"

// MultiSelectBuilder builds `SELECT ...; SELECT ...;`
type MultiSelectBuilder struct {
	runner
	EventReceiver
	Dialect Dialect

	queries []Builder
}

func (m *MultiSelectBuilder) Load(values ...interface{}) (int, error) {
	return m.LoadContext(context.Background(), values...)
}

func (m *MultiSelectBuilder) LoadContext(ctx context.Context, values ...interface{}) (int, error) {
	return query(ctx, m.runner, m.EventReceiver, m.queries, m.Dialect, values...)
}

func (sess *Session) SelectMultiple(builders ...Builder) *MultiSelectBuilder {
	return &MultiSelectBuilder{
		runner:        sess,
		EventReceiver: sess,
		Dialect:       sess.Dialect,
		queries:       builders,
	}
}
