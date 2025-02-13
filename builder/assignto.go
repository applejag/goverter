package builder

import (
	"github.com/dave/jennifer/jen"
	"github.com/jmattheis/goverter/xtype"
)

type AssignTo struct {
	Stmt   *jen.Statement
	Must   bool
	Update bool
}

func AssignOf(s *jen.Statement) *AssignTo {
	return &AssignTo{Stmt: s}
}

func (a *AssignTo) WithIndex(s *jen.Statement) *AssignTo {
	return &AssignTo{
		Stmt: a.Stmt.Clone().Index(s),
	}
}

func (a *AssignTo) MustAssign() *AssignTo {
	a.Must = true
	return a
}

func (a *AssignTo) IsUpdate() *AssignTo {
	a.Update = true
	return a
}

func ToAssignable(assignTo *AssignTo) func(stmt []jen.Code, nextID *xtype.JenID, err *Error) ([]jen.Code, *Error) {
	return func(stmt []jen.Code, nextID *xtype.JenID, err *Error) ([]jen.Code, *Error) {
		if err != nil {
			return nil, err
		}
		stmt = append(stmt, assignTo.Stmt.Clone().Op("=").Add(nextID.Code))
		return stmt, nil
	}
}

func AssignByBuild(b Builder, gen Generator, ctx *MethodContext, assignTo *AssignTo, sourceID *xtype.JenID, source, target *xtype.Type, errPath ErrorPath) ([]jen.Code, *Error) {
	return ToAssignable(assignTo)(b.Build(gen, ctx, sourceID, source, target, errPath))
}

func BuildByAssign(b Builder, gen Generator, ctx *MethodContext, sourceID *xtype.JenID, source, target *xtype.Type, path ErrorPath) ([]jen.Code, *xtype.JenID, *Error) {
	buildStmt, valueVar, err := buildTargetVar(gen, ctx, sourceID, source, target, path)
	if err != nil {
		return nil, nil, err
	}

	stmt, err := b.Assign(gen, ctx, AssignOf(valueVar), sourceID, source, target, path)
	if err != nil {
		return nil, nil, err
	}

	buildStmt = append(buildStmt, stmt...)
	return buildStmt, xtype.VariableID(valueVar), nil
}
