package plugin

import (
	"fmt"

	"github.com/Neakxs/protoc-gen-authz/authorize"
	"github.com/google/cel-go/cel"
	"github.com/google/cel-go/checker/decls"
	"github.com/google/cel-go/interpreter"
	v1alpha1 "google.golang.org/genproto/googleapis/api/expr/v1alpha1"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoregistry"
)

const PluginName = "protoc-gen-go-authz"

const generatedFilenameSuffix = ".pb.authz.go"

const (
	authorizePackage   = protogen.GoImportPath("github.com/Neakxs/protoc-gen-authz/authorize")
	contextPackage     = protogen.GoImportPath("context")
	celPackage         = protogen.GoImportPath("github.com/google/cel-go/cel")
	declsPackage       = protogen.GoImportPath("github.com/google/cel-go/checker/decls")
	interpreterPackage = protogen.GoImportPath("github.com/google/cel-go/interpreter")
	codesPackage       = protogen.GoImportPath("google.golang.org/grpc/codes")
	statusPackage      = protogen.GoImportPath("google.golang.org/grpc/status")
)

func Run(gen *protogen.Plugin) error {
	var files protoregistry.Files
	for _, file := range gen.Files {
		if err := files.RegisterFile(file.Desc); err != nil {
			return err
		}
	}
	celOpts := cel.TypeDescs(&files)
	for _, file := range gen.Files {
		if !file.Generate {
			continue
		}
		g := generateNewFile(gen, file)
		for _, msg := range file.Messages {
			generateMessageMethod(gen, celOpts, g, msg)
		}
	}
	return nil
}

func generateNewFile(gen *protogen.Plugin, file *protogen.File) *protogen.GeneratedFile {
	g := gen.NewGeneratedFile(file.GeneratedFilenamePrefix+generatedFilenameSuffix, file.GoImportPath)
	g.P("// Code generated by ", PluginName, ". DO NOT EDIT.")
	g.P("//")
	g.P("// source: ", file.Desc.Path())
	g.P()
	g.P("package ", file.GoPackageName)
	g.P()
	return g
}

func generateMessageMethod(gen *protogen.Plugin, celOpts cel.EnvOption, g *protogen.GeneratedFile, message *protogen.Message) {
	ok := generateGetProgram(gen, g, message)
	g.P(
		"\n", `func (m *`+message.GoIdent.GoName+`) Authorize(ctx `+g.QualifiedGoIdent(contextPackage.Ident(`Context`))+`) error {`,
	)
	if ok {
		g.P(
			`	pgr, err := `, buildGetProgramSignature(g, message), `()`, "\n",
			`	if err != nil {`, "\n",
			`		return err`, "\n",
			`	}`, "\n",
			`	val, _, err := pgr.ContextEval(ctx, `, g.QualifiedGoIdent(authorizePackage.Ident("BuildProgramVars")), `(ctx, m))`, "\n",
			`	if err != nil {`, "\n",
			`		return err`, "\n",
			`	}`, "\n",
			`	if res, ok := val.Value().(bool); ok {`, "\n",
			`		if res {`, "\n",
			`			return nil`, "\n",
			`		} else {`, "\n",
			`			return `, statusPackage.Ident("Error"), `(`, codesPackage.Ident("PermissionDenied"), `,"")`, "\n",
			`		}`, "\n",
			`	} else {`, "\n",
			`		return `, statusPackage.Ident("Error"), `(`, codesPackage.Ident("Unknown"), `,"")`, "\n",
			`	}`,
		)
	} else {
		g.P(
			`	return nil`,
		)
	}
	g.P(
		`}`,
	)
}

func programVarName(message *protogen.Message) string {
	return `_` + message.GoIdent.GoName + `_celProgram`
}

func buildGetProgramSignature(g *protogen.GeneratedFile, message *protogen.Message) string {
	return `_get` + message.GoIdent.GoName + `CELProgram`
}

func buildGetProgramSignatureDef(g *protogen.GeneratedFile, message *protogen.Message) string {
	return `func ` + buildGetProgramSignature(g, message) + `() (` + g.QualifiedGoIdent(celPackage.Ident("Program")) + `, error)`
}

func generateGetProgram(gen *protogen.Plugin, g *protogen.GeneratedFile, message *protogen.Message) bool {
	rule := proto.GetExtension(message.Desc.Options(), authorize.E_Rule).(*authorize.Rule)
	if rule == nil || len(rule.Expr) == 0 {
		return false
	} else {
		g.P(
			`var `, programVarName(message), ` `, g.QualifiedGoIdent(celPackage.Ident("Program")), ` = nil`, "\n",
			buildGetProgramSignatureDef(g, message), `{`,
		)
		env, err := cel.NewEnv(
			cel.Types(&authorize.AuthorizationContext{}),
			cel.DeclareContextProto(message.Desc),
			cel.Declarations(
				decls.NewVar("_ctx", decls.NewObjectType(string((&authorize.AuthorizationContext{}).ProtoReflect().Descriptor().FullName()))),
			),
		)
		if err != nil {
			gen.Error(err)
			return false
		}
		ast, issues := env.Compile(rule.Expr)
		if issues != nil && issues.Err() != nil {
			gen.Error(issues.Err())
			return false
		}
		switch ast.ResultType().TypeKind.(type) {
		case *v1alpha1.Type_Primitive:
			if v1alpha1.Type_PrimitiveType(ast.ResultType().GetPrimitive().Number()) != v1alpha1.Type_BOOL {
				gen.Error(fmt.Errorf("result type not bool"))
				return false
			}
		default:
			gen.Error(fmt.Errorf("result type not bool"))
			return false
		}
		_, err = env.Program(ast, cel.OptimizeRegex(interpreter.MatchesRegexOptimization))
		if err != nil {
			gen.Error(err)
		}
		authorizeAuthorizationContext := g.QualifiedGoIdent(authorizePackage.Ident("AuthorizationContext"))
		celNewEnv := g.QualifiedGoIdent(celPackage.Ident("NewEnv"))
		celTypes := g.QualifiedGoIdent(celPackage.Ident("Types"))
		celDeclarations := g.QualifiedGoIdent(celPackage.Ident("Declarations"))
		celDeclareContextProto := g.QualifiedGoIdent(celPackage.Ident("DeclareContextProto"))
		celOptimizeRegex := g.QualifiedGoIdent(celPackage.Ident("OptimizeRegex"))
		declsNewVar := g.QualifiedGoIdent(declsPackage.Ident("NewVar"))
		declsNewObjectType := g.QualifiedGoIdent(declsPackage.Ident("NewObjectType"))
		interpreterMatchesRegexOptimization := g.QualifiedGoIdent(interpreterPackage.Ident("MatchesRegexOptimization"))
		g.P(
			`	if `, programVarName(message), ` == nil {`, "\n",
			`		env, _ := `, celNewEnv, `(`, "\n",
			`			`, celTypes, `(&`, authorizeAuthorizationContext, `{}),`, "\n",
			`			`, celDeclareContextProto, `((&`, message.GoIdent.GoName, `{}).ProtoReflect().Descriptor()),`, "\n",
			`			`, celDeclarations, `(`, "\n",
			`				`, declsNewVar, `("_ctx", `, declsNewObjectType, `(string((&`, authorizeAuthorizationContext, `{}).ProtoReflect().Descriptor().FullName()))),`, "\n",
			`			),`, "\n",
			`		)`, "\n",
			`		ast, _ := env.Compile(`, "`", rule.Expr, "`", `)`, "\n",
			`		`, programVarName(message), `, _ = env.Program(ast, `, celOptimizeRegex, `(`, interpreterMatchesRegexOptimization, `)`, `)`, "\n",
			`	}`, "\n",
			`	return `, programVarName(message), `, nil`,
		)
		g.P(
			`}`,
		)
	}
	return true
}
