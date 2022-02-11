package plugin

import (
	"fmt"

	"github.com/Neakxs/protoc-gen-authz/api"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoregistry"
)

const PluginName = "protoc-gen-go-grpc-authz"

const generatedFilenameSuffix = "_grpc_authz.pb.go"

const (
	contextPackage  = protogen.GoImportPath("context")
	celPackage      = protogen.GoImportPath("github.com/google/cel-go/cel")
	declsPackage    = protogen.GoImportPath("github.com/google/cel-go/checker/decls")
	exprPackage     = protogen.GoImportPath("google.golang.org/genproto/googleapis/api/expr/v1alpha1")
	grpcPackage     = protogen.GoImportPath("google.golang.org/grpc")
	codesPackage    = protogen.GoImportPath("google.golang.org/grpc/codes")
	statusPackage   = protogen.GoImportPath("google.golang.org/grpc/status")
	metadataPackage = protogen.GoImportPath("google.golang.org/grpc/metadata")
)

func Run(gen *protogen.Plugin) error {
	var files protoregistry.Files
	for _, file := range gen.Files {
		if err := files.RegisterFile(file.Desc); err != nil {
			return err
		}
	}
	for _, file := range gen.Files {
		if !file.Generate {
			continue
		}
		g := generateNewFile(gen, file)
		for _, svc := range file.Services {
			m := []*protogen.Method{}
			for _, mth := range svc.Methods {
				if generateBuilder(g, mth) {
					m = append(m, mth)
				}
			}
			generateInterceptors(g, svc, m)
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

func authzBuilderSignature(g *protogen.GeneratedFile, method *protogen.Method) string {
	return `_` + method.Parent.GoName + `_` + method.GoName + `_AuthzBuilder() (` + g.QualifiedGoIdent(celPackage.Ident("Program")) + `, error)`
}

func generateBuilder(g *protogen.GeneratedFile, method *protogen.Method) bool {
	rule := proto.GetExtension(method.Desc.Options(), api.E_AuthorizationRule).(string)
	if rule == "" {
		return false
	}
	celNewEnv := g.QualifiedGoIdent(celPackage.Ident("NewEnv"))
	celTypes := g.QualifiedGoIdent(celPackage.Ident("Types"))
	celDeclarations := g.QualifiedGoIdent(celPackage.Ident("Declarations"))
	declsNewVar := g.QualifiedGoIdent(declsPackage.Ident("NewVar"))
	declsNewObjectType := g.QualifiedGoIdent(declsPackage.Ident("NewObjectType"))
	declsNewMapType := g.QualifiedGoIdent(declsPackage.Ident("NewMapType"))
	declsNewListType := g.QualifiedGoIdent(declsPackage.Ident("NewListType"))
	exprType := g.QualifiedGoIdent(exprPackage.Ident("Type"))
	exprTypePrimitive := g.QualifiedGoIdent(exprPackage.Ident("Type_Primitive"))
	exprTypeString := g.QualifiedGoIdent(exprPackage.Ident("Type_STRING"))

	g.P(`
func `, authzBuilderSignature(g, method), ` {
	env, err := `, celNewEnv, `(
		`, celTypes, `(&`, g.QualifiedGoIdent(method.Input.GoIdent), `{}),
		`, celDeclarations, `(
			`, declsNewVar, `("request", `, declsNewObjectType, `("`, method.Input.Desc.FullName(), `")),
			`, declsNewVar, `("metadata", `, declsNewMapType, `(&`, exprType, `{TypeKind: &`, exprTypePrimitive, `{Primitive: `, exprTypeString, `}}, `, declsNewListType, `(&`, exprType, `{TypeKind: &`, exprTypePrimitive, `{Primitive: `, exprTypeString, `}}))),
		),
	)
	if err != nil {
		return nil, err
	}
	ast, issues := env.Compile(`, "`", rule, "`", `)
	if issues != nil && issues.Err() != nil {
		return nil, issues.Err()
	}
	return env.Program(ast)
}
	`)
	return true
}

func authzMapSignature(g *protogen.GeneratedFile, service *protogen.Service) string {
	return "_" + service.GoName + "_Authz = map[string]func() (" + g.QualifiedGoIdent(celPackage.Ident("Program")) + ", error)"
}

func unaryInterceptorSignature(g *protogen.GeneratedFile, service *protogen.Service) string {
	return `New` + service.GoName + `AuthzUnaryServerInterceptor() (` + g.QualifiedGoIdent(grpcPackage.Ident("UnaryServerInterceptor")) + `, error)`
}

func generateInterceptors(g *protogen.GeneratedFile, service *protogen.Service, methods []*protogen.Method) {
	g.P("var ", authzMapSignature(g, service), " {")
	for _, method := range methods {
		g.P(`	"`, fmt.Sprintf("/%s/%s", method.Parent.Desc.FullName(), method.Desc.Name()), `": _`, method.Parent.GoName, `_`, method.GoName, `_AuthzBuilder,`)
	}
	g.P("}")
	g.P()
	celProgram := g.QualifiedGoIdent(celPackage.Ident("Program"))
	contextContext := g.QualifiedGoIdent(contextPackage.Ident("Context"))
	grpcUnaryServerInfo := g.QualifiedGoIdent(grpcPackage.Ident("UnaryServerInfo"))
	grpcUnaryHandler := g.QualifiedGoIdent(grpcPackage.Ident("UnaryHandler"))
	metadataFromIncomingContext := g.QualifiedGoIdent(metadataPackage.Ident("FromIncomingContext"))
	metadataNew := g.QualifiedGoIdent(metadataPackage.Ident("New"))
	statusError := g.QualifiedGoIdent(statusPackage.Ident("Error"))
	codesPermissionDenied := g.QualifiedGoIdent(codesPackage.Ident("PermissionDenied"))
	g.P(`
func `, unaryInterceptorSignature(g, service), ` {
	m := map[string]`, celProgram, `{}
	for k,v := range _`, service.GoName, `_Authz {
		if pgr, err := v(); err == nil {
			m[k] = pgr
		} else {
			return nil, err
		}
	}
	return func(ctx `, contextContext, `, req interface{}, info *`, grpcUnaryServerInfo, `, handler `, grpcUnaryHandler, `) (resp interface{}, err error) {
		if pgr, ok := m[info.FullMethod]; !ok {
			return handler(ctx, req)
		} else {
			md, ok := `, metadataFromIncomingContext, `(ctx)
			if !ok {
				md = `, metadataNew, `(map[string]string{})
			}
			if val, _, err := pgr.Eval(map[string]interface{}{
				"request": req,
				"metadata": md,
			}); err != nil {
				return nil, err
			} else if b, ok := val.Value().(bool); !ok || !b {
				return nil, `, statusError, `(`, codesPermissionDenied, `, "Unauthorized")
			}
			return handler(ctx, req)
		}
	}, nil
}
	`)
}
