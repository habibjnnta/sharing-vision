package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Posts holds the schema definition for the Posts entity.
type Posts struct {
	ent.Schema
}

// Fields of the Posts.
func (Posts) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").
			StorageKey("id").
			DefaultFunc(func() int { return 0 }),
		field.String("title").MaxLen(200),
		field.Text("content"),
		field.String("category").MaxLen(100),
		field.Enum("status").Values("publish", "draft", "thrash"),
	}
}

// Edges of the Posts.
func (Posts) Edges() []ent.Edge {
	return nil
}

// Add created_at, updated_at, and deleted_at
func (Posts) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
