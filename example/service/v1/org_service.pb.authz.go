// Code generated by protoc-gen-go-authz. DO NOT EDIT.
// versions:
//  protoc-gen-go-authz	v0.0.0
//  protoc 				v3.21.1
// source: example/service/v1/org_service.proto

package v1

import (
	authorize "github.com/Neakxs/protoc-gen-authz/authorize"
	cel "github.com/google/cel-go/cel"
	proto "google.golang.org/protobuf/proto"
)

// Authz global configuration
var _File_example_service_v1_org_service_proto_authzConfiguration = &authorize.FileRule{
	Globals: &authorize.FileRule_Globals{
		Functions: map[string]string{
			"canPong": `headers.get("Xpong") == "yes"`,
		},
		Constants: map[string]string{
			"xpong": `x-pong`,
		},
	},
	Overloads: &authorize.FileRule_Overloads{
		Functions: map[string]*authorize.FileRule_Overloads_Function{
			"do": {
				Args: []*authorize.FileRule_Overloads_Type{
					&authorize.FileRule_Overloads_Type{
						Type: &authorize.FileRule_Overloads_Type_Primitive_{
							Primitive: authorize.FileRule_Overloads_Type_STRING,
						},
					},
					&authorize.FileRule_Overloads_Type{
						Type: &authorize.FileRule_Overloads_Type_Primitive_{
							Primitive: authorize.FileRule_Overloads_Type_STRING,
						},
					},
				},
				Result: &authorize.FileRule_Overloads_Type{
					Type: &authorize.FileRule_Overloads_Type_Primitive_{
						Primitive: authorize.FileRule_Overloads_Type_BOOL,
					},
				},
			},
		},
		Variables: map[string]*authorize.FileRule_Overloads_Type{
			"ping": &authorize.FileRule_Overloads_Type{
				Type: &authorize.FileRule_Overloads_Type_Primitive_{
					Primitive: authorize.FileRule_Overloads_Type_STRING,
				},
			},
		},
	},
}

func NewOrgServiceAuthzInterceptor(opts ...authorize.Options) (authorize.AuthzInterceptor, error) {
	lib := authorize.BuildRuntimeLibrary(_File_example_service_v1_org_service_proto_authzConfiguration, opts...)
	m := map[string]cel.Program{}
	for k, v := range map[string]struct {
		expr string
		req  proto.Message
	}{
		"/service.v1.OrgService/Ping": {expr: `headers.get("ok") == "ok" && do(ping, xpong) && !canPong() && size(request.ping) > 0`, req: &PingRequest{}},
		"/service.v1.OrgService/Pong": {expr: `canPong() && size(request.pong) > 0`, req: &PongRequest{}},
	} {
		if pgr, err := authorize.BuildAuthzProgram(v.expr, v.req, _File_example_service_v1_org_service_proto_authzConfiguration, lib); err != nil {
			return nil, err
		} else {
			m[k] = pgr
		}
	}
	return authorize.NewAuthzInterceptor(m), nil
}
