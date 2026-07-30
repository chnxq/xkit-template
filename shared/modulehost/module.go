package modulehost

import (
	"context"
	"sort"
	"sync"
	"time"

	"github.com/chnxq/xkitpkg/secret"
	httptransport "github.com/chnxq/xkitpkg/transport/http"
	"google.golang.org/grpc"
)

type RuntimeKind string

const (
	RuntimeKindInProcess RuntimeKind = "IN_PROCESS"
	RuntimeKindRemote    RuntimeKind = "REMOTE_SERVICE"
)

type DataIsolationMode string

const (
	DataIsolationNone              DataIsolationMode = "NONE"
	DataIsolationShared            DataIsolationMode = "SHARED"
	DataIsolationDatabasePerTenant DataIsolationMode = "DATABASE_PER_TENANT"
	DataIsolationSchemaPerTenant   DataIsolationMode = "SCHEMA_PER_TENANT"
)

type Descriptor struct {
	Code          string
	Name          string
	RuntimeKind   RuntimeKind
	DataIsolation DataIsolationMode
	Version       string
}

type Definition interface {
	Descriptor() Descriptor
	Resources() ResourceBundle
	BuildRuntime(HostServices) (Runtime, error)
}

type Runtime interface {
	Descriptor() Descriptor
	RegisterHTTP(*httptransport.Server)
	RegisterGRPC(grpc.ServiceRegistrar)
	DataPlane() DataPlane
	SyncResources(context.Context, ResourceSyncer) error
	SeedDefaultData(context.Context) error
	Close(context.Context) error
}

type DataPlane interface {
	Validate(context.Context, DeploymentSpec) (ValidationResult, error)
	Provision(context.Context, DeploymentSpec) error
	Migrate(context.Context, DeploymentSpec, string) (string, error)
	Health(context.Context, DeploymentSpec) HealthResult
	Invalidate(DeploymentKey)
}

type DeploymentKey struct {
	ModuleCode string
	TenantID   uint64
}

type DeploymentSpec struct {
	Key                  DeploymentKey
	DeploymentID         uint64
	DeploymentVersion    uint64
	DataSourceID         uint64
	ConfigurationVersion uint64
	SecretVersion        string
	Driver               string
	Endpoint             string
	DatabaseName         string
	SchemaName           string
	DesiredSchemaVersion string
	ActualSchemaVersion  string
	Credential           secret.Ref
	Pool                 PoolPolicy
}

type PoolPolicy struct {
	MaxIdleConnections uint32
	MaxOpenConnections uint32
	ConnectionMaxIdle  time.Duration
	ConnectionMaxLife  time.Duration
}

type ValidationResult struct {
	Valid         bool
	DatabaseInfo  string
	SchemaVersion string
	Warnings      []string
}

type HealthStatus string

const (
	HealthUnknown   HealthStatus = "UNKNOWN"
	HealthHealthy   HealthStatus = "HEALTHY"
	HealthUnhealthy HealthStatus = "UNHEALTHY"
)

type HealthResult struct {
	Status        HealthStatus
	SchemaVersion string
	ErrorCode     string
	CheckedAt     time.Time
}

type DeploymentResolver interface {
	Resolve(context.Context, DeploymentKey) (DeploymentSpec, error)
	Watch(context.Context) (<-chan DeploymentChange, error)
}

type DeploymentChange struct {
	Key     DeploymentKey
	Version uint64
}

type AuditEvent struct {
	Type       string
	ActorID    uint64
	TenantID   uint64
	ModuleCode string
	Operation  string
	Resource   string
	OccurredAt time.Time
	Attributes map[string]string
}

type AuditPublisher interface {
	Publish(context.Context, AuditEvent) error
}

type Logger interface {
	Debugf(string, ...any)
	Infof(string, ...any)
	Warnf(string, ...any)
	Errorf(string, ...any)
}

type LoggerFactory interface {
	NewLogger(string) Logger
}

type HostServices struct {
	Context        context.Context
	Loggers        LoggerFactory
	Deployments    DeploymentResolver
	Secrets        secret.Provider
	Audit          AuditPublisher
	ResourceSyncer ResourceSyncer
}

type ResourceBundle struct {
	OpenAPI []OpenAPIDocument
	Menus   []MenuResource
}

type OpenAPIDocument struct {
	Name     string
	FileName string
	Format   string
	Data     []byte
}

type MenuType string

const (
	MenuTypeCatalog  MenuType = "catalog"
	MenuTypeMenu     MenuType = "menu"
	MenuTypeEmbedded MenuType = "embedded"
	MenuTypeLink     MenuType = "link"
	MenuTypeButton   MenuType = "button"
)

type MenuMeta struct {
	Authority       []string
	Icon            *string
	Link            *string
	OpenInNewWindow *bool
	Title           *string
	TitleAux        *string
}

type MenuResource struct {
	Children  []MenuResource
	Component string
	Meta      MenuMeta
	Name      string
	Path      string
	Redirect  string
	Type      MenuType
}

type ResourceSyncer interface {
	UpsertMenus(context.Context, []MenuResource) error
}

var registry = struct {
	sync.RWMutex
	definitions map[string]Definition
}{definitions: make(map[string]Definition)}

func RegisterDefinition(definition Definition) {
	if definition == nil {
		return
	}
	descriptor := definition.Descriptor()
	if descriptor.Code == "" {
		panic("module definition code is required")
	}
	registry.Lock()
	defer registry.Unlock()
	if _, exists := registry.definitions[descriptor.Code]; exists {
		panic("duplicate module definition: " + descriptor.Code)
	}
	registry.definitions[descriptor.Code] = definition
}

func Definitions() []Definition {
	registry.RLock()
	defer registry.RUnlock()
	result := make([]Definition, 0, len(registry.definitions))
	for _, definition := range registry.definitions {
		result = append(result, definition)
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].Descriptor().Code < result[j].Descriptor().Code
	})
	return result
}

func DefinitionByCode(code string) (Definition, bool) {
	registry.RLock()
	defer registry.RUnlock()
	definition, ok := registry.definitions[code]
	return definition, ok
}
