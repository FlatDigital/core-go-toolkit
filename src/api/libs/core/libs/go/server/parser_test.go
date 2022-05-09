package server_test

import (
	"testing"

	"github.com/FlatDigital/flat-go-toolkit/src/api/libs/core/libs/go/server"
)

func TestContextFromScopeString(t *testing.T) {
	tt := []struct {
		Scope         string
		ExpectedRole  server.Role
		ExpectedScope server.Environment
		ExpectedTag   string
		ExpectedErr   bool
	}{
		{"production-read", server.RoleRead, server.EnvProduction, "", false},
		{"sandbox-read", server.RoleRead, server.EnvSandbox, "", false},
		{"develop-read", server.RoleRead, server.EnvDevelop, "", false},
		{"test-read", server.RoleRead, server.EnvTest, "", false},

		{"production-read-tag", server.RoleRead, server.EnvProduction, "tag", false},
		{"production-read-feature-new-search", server.RoleRead, server.EnvProduction, "feature-new-search", false},

		{"custom-read", server.RoleRead, server.Environment("custom"), "", false},
		{"custom-read-appname", server.RoleRead, server.Environment("custom"), "appname", false},
		{"test-custom", server.Role("custom"), server.EnvTest, "", false},
		{"test-custom-appname", server.Role("custom"), server.EnvTest, "appname", false},

		{"invalid", server.RoleRead, server.EnvTest, "", true},
	}

	for _, tc := range tt {
		t.Run(tc.Scope, func(t *testing.T) {
			ctx, err := server.ContextFromScopeString(tc.Scope)

			if err != nil {
				if !tc.ExpectedErr {
					t.Fatalf("Unexpected error returned: %v", err)
				}

				return
			}

			if ctx.Role != tc.ExpectedRole {
				t.Fatalf("Expected role to be %v, got: %v", tc.ExpectedRole, ctx.Role)
			}

			if ctx.Environment != tc.ExpectedScope {
				t.Fatalf("Expected scope to be %v, got: %v", tc.ExpectedScope, ctx.Environment)
			}

			if ctx.Tag != tc.ExpectedTag {
				t.Fatalf(`Expected tag to be "%v", got: "%v"`, tc.ExpectedTag, ctx.Tag)
			}
		})
	}
}
