package modulehost

import "testing"

type testDefinition struct {
	code string
}

func (d testDefinition) Descriptor() Descriptor {
	return Descriptor{Code: d.code}
}

func (testDefinition) Resources() ResourceBundle { return ResourceBundle{} }
func (testDefinition) BuildRuntime(HostServices) (Runtime, error) {
	return nil, nil
}

func TestDefinitionsAreSorted(t *testing.T) {
	registry.Lock()
	previous := registry.definitions
	registry.definitions = make(map[string]Definition)
	registry.Unlock()
	t.Cleanup(func() {
		registry.Lock()
		registry.definitions = previous
		registry.Unlock()
	})

	RegisterDefinition(testDefinition{code: "zeta"})
	RegisterDefinition(testDefinition{code: "alpha"})
	definitions := Definitions()
	if len(definitions) != 2 {
		t.Fatalf("definition count = %d, want 2", len(definitions))
	}
	if definitions[0].Descriptor().Code != "alpha" || definitions[1].Descriptor().Code != "zeta" {
		t.Fatalf("definition order = %q, %q", definitions[0].Descriptor().Code, definitions[1].Descriptor().Code)
	}
}
