package authorize

import (
	"fmt"

	"github.com/google/cel-go/cel"
	"github.com/google/cel-go/checker/decls"
	"github.com/google/cel-go/interpreter"
	"github.com/google/cel-go/parser"
	v1alpha1 "google.golang.org/genproto/googleapis/api/expr/v1alpha1"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func BuildAuthzProgram(expr string, req interface{}, config *FileRule, libs ...cel.Library) (cel.Program, error) {
	var reqDesc protoreflect.MessageDescriptor
	var reqOpt cel.EnvOption
	if r, ok := req.(proto.Message); ok {
		reqOpt = cel.Types(r)
		reqDesc = r.ProtoReflect().Descriptor()
	} else if r, ok := req.(protoreflect.MessageDescriptor); ok {
		reqOpt = cel.TypeDescs(r.ParentFile())
		reqDesc = r
	} else {
		return nil, fmt.Errorf("invalid req")
	}
	authzCtx := &AuthorizationContext{}
	envOpts := []cel.EnvOption{
		reqOpt,
		cel.Types(authzCtx),
		cel.Declarations(
			decls.NewVar(
				"context",
				decls.NewObjectType(string(authzCtx.ProtoReflect().Descriptor().FullName()))),
			decls.NewVar(
				"request",
				decls.NewObjectType(string(reqDesc.FullName())),
			),
		),
		buildGlobalConstantsOption(config),
		buildOverloadFunctionsOption(config),
	}
	for i := 0; i < len(libs); i++ {
		envOpts = append(envOpts, cel.Lib(libs[i]))
	}
	macros := []parser.Macro{}
	if rawMacros, err := findMacros(config, envOpts, expr); err != nil {
		return nil, err
	} else {
		env, err := cel.NewEnv(envOpts...)
		if err != nil {
			return nil, fmt.Errorf("new env error: %w", err)
		}
		for _, macro := range rawMacros {
			ast, issues := env.Compile(config.Globals.Functions[macro])
			if issues != nil && issues.Err() != nil {
				return nil, issues.Err()
			}
			macros = append(macros, parser.NewGlobalMacro(macro, 0, BuildMacroExpander(ast)))
		}
	}
	envOpts = append(envOpts, cel.Macros(macros...))
	env, err := cel.NewEnv(envOpts...)
	if err != nil {
		return nil, fmt.Errorf("new env error: %w", err)
	}
	ast, issues := env.Compile(expr)
	if issues != nil && issues.Err() != nil {
		return nil, fmt.Errorf("compile error: %w", issues.Err())
	}
	switch ast.ResultType().TypeKind.(type) {
	case *v1alpha1.Type_Primitive:
		if v1alpha1.Type_PrimitiveType(ast.ResultType().GetPrimitive().Number()) != v1alpha1.Type_BOOL {
			return nil, fmt.Errorf("result type not bool")
		}
	default:
		return nil, fmt.Errorf("result type not bool")
	}
	pgr, err := env.Program(ast, cel.OptimizeRegex(interpreter.MatchesRegexOptimization))
	if err != nil {
		return nil, fmt.Errorf("program error: %w", err)
	}
	return pgr, nil
}

func buildGlobalConstantsOption(config *FileRule) cel.EnvOption {
	constantDecls := []*v1alpha1.Decl{}
	if config != nil && config.Globals != nil {
		for k, v := range config.Globals.Constants {
			constantDecls = append(constantDecls, decls.NewConst(
				k,
				decls.String,
				&v1alpha1.Constant{ConstantKind: &v1alpha1.Constant_StringValue{StringValue: v}},
			))
		}
	}
	return cel.Declarations(constantDecls...)
}

func buildOverloadFunctionsOption(config *FileRule) cel.EnvOption {
	functionDecls := []*v1alpha1.Decl{}
	if config != nil && config.Overloads != nil {
		for name, v := range config.Overloads.Functions {
			args := []*v1alpha1.Type{}
			overload := name
			for i := 0; i < len(v.Args); i++ {
				args = append(args, TypeFromFunctionType(v.Args[i]))
			}
			functionDecls = append(functionDecls, decls.NewFunction(
				name, decls.NewOverload(
					overload,
					args,
					TypeFromFunctionType(v.Result),
				),
			))
		}
	}
	return cel.Declarations(functionDecls...)
}

func findMacros(config *FileRule, opts []cel.EnvOption, expr string) ([]string, error) {
	if config == nil {
		return nil, nil
	}
	envOpts := opts
	for k := range config.Globals.Functions {
		envOpts = append(envOpts, cel.Declarations(decls.NewFunction(k, decls.NewOverload(k, []*v1alpha1.Type{}, &v1alpha1.Type{TypeKind: &v1alpha1.Type_Dyn{}}))))
	}
	env, err := cel.NewEnv(envOpts...)
	if err != nil {
		return nil, fmt.Errorf("new env error: %w", err)
	}
	ast, issues := env.Compile(expr)
	if issues != nil && issues.Err() != nil {
		return nil, fmt.Errorf("compile error: %w", issues.Err())
	}
	return findMacrosInAST(ast, config.Globals.Functions), nil
}

func findMacrosInAST(ast *cel.Ast, m map[string]string) []string {
	s := findMacrosInExpr(ast.Expr(), m)
	mm := map[string]bool{}
	for _, i := range s {
		mm[i] = true
	}
	s = []string{}
	for k := range mm {
		s = append(s, k)
	}
	return s
}

func findMacrosInExpr(e *v1alpha1.Expr, m map[string]string) []string {
	res := []string{}
	switch exp := e.ExprKind.(type) {
	case *v1alpha1.Expr_ConstExpr:
	case *v1alpha1.Expr_IdentExpr:
	case *v1alpha1.Expr_SelectExpr:
	case *v1alpha1.Expr_CallExpr:
		if _, ok := m[exp.CallExpr.Function]; ok {
			res = append(res, exp.CallExpr.Function)
		} else {
			for _, i := range exp.CallExpr.Args {
				res = append(res, findMacrosInExpr(i, m)...)
			}
		}
	case *v1alpha1.Expr_ListExpr:
		for _, i := range exp.ListExpr.Elements {
			res = append(res, findMacrosInExpr(i, m)...)
		}
	case *v1alpha1.Expr_StructExpr:
	case *v1alpha1.Expr_ComprehensionExpr:
	}
	return res
}
