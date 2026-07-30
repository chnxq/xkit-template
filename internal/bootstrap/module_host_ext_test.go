package bootstrap

import (
	"context"
	"testing"

	crudviewer "github.com/chnxq/x-crud/viewer"
)

type moduleHostTestViewer struct {
	tenantID uint64
	tenant   bool
}

func (v moduleHostTestViewer) UserID() uint64                    { return 1 }
func (v moduleHostTestViewer) TenantID() uint64                  { return v.tenantID }
func (v moduleHostTestViewer) OrgUnitID() uint64                 { return 0 }
func (v moduleHostTestViewer) Permissions() []string             { return nil }
func (v moduleHostTestViewer) Roles() []string                   { return nil }
func (v moduleHostTestViewer) DataScope() []crudviewer.DataScope { return nil }
func (v moduleHostTestViewer) TraceID() string                   { return "" }
func (v moduleHostTestViewer) HasPermission(string, string) bool { return false }
func (v moduleHostTestViewer) IsPlatformContext() bool           { return !v.tenant }
func (v moduleHostTestViewer) IsTenantContext() bool             { return v.tenant }
func (v moduleHostTestViewer) IsSystemContext() bool             { return false }
func (v moduleHostTestViewer) ShouldAudit() bool                 { return false }

func TestViewerTenantScopeResolver(t *testing.T) {
	resolver := viewerTenantScopeResolver{}

	tenantCtx := crudviewer.WithContext(context.Background(), moduleHostTestViewer{tenantID: 7, tenant: true})
	tenantID, err := resolver.TenantID(tenantCtx)
	if err != nil || tenantID != 7 {
		t.Fatalf("tenant scope = %d, %v; want 7, nil", tenantID, err)
	}

	platformCtx := crudviewer.WithContext(context.Background(), moduleHostTestViewer{})
	if _, err := resolver.TenantID(platformCtx); err == nil {
		t.Fatal("platform context must not resolve an implicit module tenant")
	}
	if _, err := resolver.TenantID(context.Background()); err == nil {
		t.Fatal("missing viewer must not resolve a module tenant")
	}
}
