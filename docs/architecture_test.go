package docs

import (
	"testing"

	archgo "github.com/fdaines/arch-go/api"
	config "github.com/fdaines/arch-go/api/configuration"
)

func TestArchitecture(t *testing.T) {
	configuration := config.Config{
		DependenciesRules: []*config.DependenciesRule{
			{
				Package: "cmd/myapp",
				ShouldOnlyDependsOn: &config.Dependencies{
					Internal: []string{"initializer"},
				},
			},
			{
				Package: "initializer",
				ShouldOnlyDependsOn: &config.Dependencies{
					Internal: []string{"internal/**", "pkg", "db"},
				},
			},
			{
				Package: "db",
				ShouldOnlyDependsOn: &config.Dependencies{
					Internal: []string{"pkg"},
				},
			},
			{
				Package: "internal/cron-jobs/**",
				ShouldOnlyDependsOn: &config.Dependencies{
					Internal: []string{"internal/**", "pkg/constants"},
				},
			},
			{
				Package: "internal/external-api/rate/**",
				ShouldOnlyDependsOn: &config.Dependencies{
					Internal: []string{"pkg/constants", "pkg/util"},
				},
			},
			{
				Package: "internal/mails/**",
				ShouldOnlyDependsOn: &config.Dependencies{
					Internal: []string{
						"internal/cron-jobs/**", "internal/rates/**", "internal/subscribers/**",
						"pkg/constants", "pkg/util", "pkg/models",
					},
				},
			},
			{
				Package: "internal/rates/**",
				ShouldOnlyDependsOn: &config.Dependencies{
					Internal: []string{"pkg/models", "pkg/constants", "infra/external-api/rate/**"},
				},
			},
			{
				Package: "internal/subscribers",
				ShouldOnlyDependsOn: &config.Dependencies{
					Internal: []string{"pkg/models"},
				},
			},
			{
				Package:             "pkg/constants",
				ShouldOnlyDependsOn: &config.Dependencies{},
			},
			{
				Package:             "pkg/models",
				ShouldOnlyDependsOn: &config.Dependencies{},
			},
			{
				Package:             "pkg/util",
				ShouldOnlyDependsOn: &config.Dependencies{},
			},
		},
	}

	moduleInfo := config.Load("github.com/Gurmigou/se-school-case-2024")

	result := archgo.CheckArchitecture(moduleInfo, configuration)

	if !result.Passes {
		t.Fatal("Architecture tests failed!")
	}
}
