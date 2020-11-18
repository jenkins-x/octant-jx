package links

import (
	"fmt"
	"strings"

	"github.com/jenkins-x/octant-jx/pkg/plugin"
)

func GetDeploymentsLink(namespace string) string {
	return plugin.PathPrefix + "/overview/namespace/" + namespace + "/workloads/deployments"
}

func GetDeploymentLink(namespace, name string) string {
	return GetDeploymentsLink(namespace) + "/" + name
}

func GetJobsLink(namespace string) string {
	return plugin.PathPrefix + "/overview/namespace/" + namespace + "/workloads/jobs"
}

func GetJobLink(namespace, name string) string {
	return GetJobsLink(namespace) + "/" + name
}

func GetPodsLink(namespace string) string {
	return plugin.PathPrefix + "/overview/namespace/" + namespace + "/workloads/pods"
}

func GetPodLink(namespace, name string) string {
	return GetPodsLink(namespace) + "/" + name
}

func GetPodTerminalLink(namespace, name string) string {
	return GetPodLink(namespace, name) + "#terminal"
}

func GetSecretsLink(namespace string) string {
	return plugin.PathPrefix + "/overview/namespace/" + namespace + "/config-and-storage/secrets"
}

func GetSecretLink(namespace, name string) string {
	return GetSecretsLink(namespace) + "/" + name
}

// GetImageLink returns the HTML link for the given container image name
func GetImageLink(image string) string {
	if image == "" {
		return ""
	}
	paths := strings.Split(image, "/")
	registry := paths[0]
	switch len(paths) {
	case 1:
		return fmt.Sprintf("https://hub.docker.com/_/%s", image)
	case 2:
		return fmt.Sprintf("https://hub.docker.com/r/%s/%s", paths[0], paths[1])
	case 3:
		if registry == "index.docker.io" {
			return fmt.Sprintf("https://hub.docker.com/r/%s/%s", paths[1], paths[2])
		}
	}
	if registry == "gcr.io" || strings.HasSuffix(registry, ".gcr.io") {
		return "https://" + image
	}
	return ""
}
