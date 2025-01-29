// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/safedep/vet/ent/codesourcefile"
	"github.com/safedep/vet/ent/predicate"
)

// CodeSourceFileUpdate is the builder for updating CodeSourceFile entities.
type CodeSourceFileUpdate struct {
	config
	hooks    []Hook
	mutation *CodeSourceFileMutation
}

// Where appends a list predicates to the CodeSourceFileUpdate builder.
func (csfu *CodeSourceFileUpdate) Where(ps ...predicate.CodeSourceFile) *CodeSourceFileUpdate {
	csfu.mutation.Where(ps...)
	return csfu
}

// SetPath sets the "path" field.
func (csfu *CodeSourceFileUpdate) SetPath(s string) *CodeSourceFileUpdate {
	csfu.mutation.SetPath(s)
	return csfu
}

// SetNillablePath sets the "path" field if the given value is not nil.
func (csfu *CodeSourceFileUpdate) SetNillablePath(s *string) *CodeSourceFileUpdate {
	if s != nil {
		csfu.SetPath(*s)
	}
	return csfu
}

// Mutation returns the CodeSourceFileMutation object of the builder.
func (csfu *CodeSourceFileUpdate) Mutation() *CodeSourceFileMutation {
	return csfu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (csfu *CodeSourceFileUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, csfu.sqlSave, csfu.mutation, csfu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (csfu *CodeSourceFileUpdate) SaveX(ctx context.Context) int {
	affected, err := csfu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (csfu *CodeSourceFileUpdate) Exec(ctx context.Context) error {
	_, err := csfu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (csfu *CodeSourceFileUpdate) ExecX(ctx context.Context) {
	if err := csfu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (csfu *CodeSourceFileUpdate) check() error {
	if v, ok := csfu.mutation.Path(); ok {
		if err := codesourcefile.PathValidator(v); err != nil {
			return &ValidationError{Name: "path", err: fmt.Errorf(`ent: validator failed for field "CodeSourceFile.path": %w`, err)}
		}
	}
	return nil
}

func (csfu *CodeSourceFileUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := csfu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(codesourcefile.Table, codesourcefile.Columns, sqlgraph.NewFieldSpec(codesourcefile.FieldID, field.TypeInt))
	if ps := csfu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := csfu.mutation.Path(); ok {
		_spec.SetField(codesourcefile.FieldPath, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, csfu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{codesourcefile.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	csfu.mutation.done = true
	return n, nil
}

// CodeSourceFileUpdateOne is the builder for updating a single CodeSourceFile entity.
type CodeSourceFileUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CodeSourceFileMutation
}

// SetPath sets the "path" field.
func (csfuo *CodeSourceFileUpdateOne) SetPath(s string) *CodeSourceFileUpdateOne {
	csfuo.mutation.SetPath(s)
	return csfuo
}

// SetNillablePath sets the "path" field if the given value is not nil.
func (csfuo *CodeSourceFileUpdateOne) SetNillablePath(s *string) *CodeSourceFileUpdateOne {
	if s != nil {
		csfuo.SetPath(*s)
	}
	return csfuo
}

// Mutation returns the CodeSourceFileMutation object of the builder.
func (csfuo *CodeSourceFileUpdateOne) Mutation() *CodeSourceFileMutation {
	return csfuo.mutation
}

// Where appends a list predicates to the CodeSourceFileUpdate builder.
func (csfuo *CodeSourceFileUpdateOne) Where(ps ...predicate.CodeSourceFile) *CodeSourceFileUpdateOne {
	csfuo.mutation.Where(ps...)
	return csfuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (csfuo *CodeSourceFileUpdateOne) Select(field string, fields ...string) *CodeSourceFileUpdateOne {
	csfuo.fields = append([]string{field}, fields...)
	return csfuo
}

// Save executes the query and returns the updated CodeSourceFile entity.
func (csfuo *CodeSourceFileUpdateOne) Save(ctx context.Context) (*CodeSourceFile, error) {
	return withHooks(ctx, csfuo.sqlSave, csfuo.mutation, csfuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (csfuo *CodeSourceFileUpdateOne) SaveX(ctx context.Context) *CodeSourceFile {
	node, err := csfuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (csfuo *CodeSourceFileUpdateOne) Exec(ctx context.Context) error {
	_, err := csfuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (csfuo *CodeSourceFileUpdateOne) ExecX(ctx context.Context) {
	if err := csfuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (csfuo *CodeSourceFileUpdateOne) check() error {
	if v, ok := csfuo.mutation.Path(); ok {
		if err := codesourcefile.PathValidator(v); err != nil {
			return &ValidationError{Name: "path", err: fmt.Errorf(`ent: validator failed for field "CodeSourceFile.path": %w`, err)}
		}
	}
	return nil
}

func (csfuo *CodeSourceFileUpdateOne) sqlSave(ctx context.Context) (_node *CodeSourceFile, err error) {
	if err := csfuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(codesourcefile.Table, codesourcefile.Columns, sqlgraph.NewFieldSpec(codesourcefile.FieldID, field.TypeInt))
	id, ok := csfuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "CodeSourceFile.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := csfuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, codesourcefile.FieldID)
		for _, f := range fields {
			if !codesourcefile.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != codesourcefile.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := csfuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := csfuo.mutation.Path(); ok {
		_spec.SetField(codesourcefile.FieldPath, field.TypeString, value)
	}
	_node = &CodeSourceFile{config: csfuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, csfuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{codesourcefile.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	csfuo.mutation.done = true
	return _node, nil
}
