package plugin

import (
	"github.com/vmware-tanzu/octant/pkg/navigation"
)

const (
	Name = "jx" // This should stay lowercase for routing purposes

	// PathPrefix the initial prefix of all paths
	PathPrefix = "/#"

	AppsPath               = "apps"
	EnvironmentsPath       = "environments"
	HelmPath               = "helm"
	JenkinsPath            = "jenkins"
	LogsPath               = "logs"
	OverviewPath           = "overview"
	PipelinesPath          = "pipelines"
	PipelineContainersPath = "pipeline/containers"
	PipelineContainerPath  = "pipeline/container"
	PipelineTerminalPath   = "pipeline/terminal"
	PipelinesRecentPath    = "pipelines-recent"
	PreviewsPath           = "previews"
	RepositoriesPath       = "repositories"

	// RootBreadcrumb the root breadcrumb for the developer plugin
	RootBreadcrumb = `<a href="/#/jx/overview">Jenkins X</a>`
)

var (
	// Navigations the default navigations
	Navigations = []navigation.Navigation{
		{
			Title: "Apps",
			Path:  Name + "/" + AppsPath,
		},
		{
			Title: "Environments",
			Path:  Name + "/" + EnvironmentsPath,
		},
		{
			Title: "Helm",
			Path:  Name + "/" + HelmPath,
		},
		{
			Title: "Jenkins",
			Path:  Name + "/" + JenkinsPath,
		},
		{
			Title: "Pipelines",
			Path:  Name + "/" + PipelinesPath,
		},
		{
			Title: "Pipelines: Recent",
			Path:  Name + "/" + PipelinesRecentPath,
		},
		{
			Title: "Previews",
			Path:  Name + "/" + PreviewsPath,
		},
		{
			Title: "Repositories",
			Path:  Name + "/" + RepositoriesPath,
		},
	}
)
