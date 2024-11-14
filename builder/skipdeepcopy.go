package builder

import (
	"go/types"

	"github.com/dave/jennifer/jen"
	"github.com/jmattheis/goverter/xtype"
)

// SkipDeepCopy handles direct assignment instead of deep copying.
type SkipDeepCopy struct{}

// Matches returns true, if the builder can create handle the given types.
func (*SkipDeepCopy) Matches(ctx *MethodContext, source, target *xtype.Type) bool {
	return !ctx.Conf.DeepCopySameType && types.Identical(source.T, target.T)
}

// Build creates conversion source code for the given source and target type.
func (*SkipDeepCopy) Build(_ Generator, _ *MethodContext, sourceID *xtype.JenID, _, _ *xtype.Type, _ ErrorPath) ([]jen.Code, *xtype.JenID, *Error) {
	return nil, sourceID, nil
}

func (*SkipDeepCopy) Assign(_ Generator, _ *MethodContext, assignTo *AssignTo, sourceID *xtype.JenID, _, _ *xtype.Type, _ ErrorPath) ([]jen.Code, *Error) {
	return []jen.Code{assignTo.Stmt.Clone().Op("=").Add(sourceID.Code)}, nil
}
