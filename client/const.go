package client

const (
	applicationName = "k8ctl"        // Application name.
	serverName      = "k8ctl-server" // Server name.
	version         = "1.0.0"        // Application version.

	// Helm related
	httpRouteReleases        = "/releases"             // List, deploy releases.(?e=environment; POST)
	httpRouteRelease         = "/releases/%s"          // Get, or Delete a release. (GET=status; DELETE=delete)
	httpRouteReleaseRollback = "/releases/%s/rollback" // Rollback a release. (PUT=rollback)
	httpRouteReleaseHistory  = "/releases/%s/history"  // Display the history of a release.

	// Kube related
	httpRouteCronjobs          = "/cronjobs"               // Display a list of cronjobs.
	httpRouteCronjob           = "/cronjobs/%s"            // Display details of a cronjob.
	httpRouteDeployments       = "/deployments"            // Display a list of deployments.
	httpRouteDeployment        = "/deployments/%s"         // Display details of a deployment.
	httpRouteDeploymentRestart = "/deployments/%s/restart" // Restart a deployment and its pods (PATCH)
	httpRouteIngresses         = "/ingresses"              // Display a list of ingresses.
	httpRouteIngress           = "/ingresses/%s"           // Display details of an ingress.
	httpRouteJobs              = "/jobs"                   // Display a list of jobs.
	httpRouteJob               = "/jobs/%s"                // Display details of a jobs.
	httpRoutePods              = "/pods"                   // Display a list of running pods.
	httpRoutePod               = "/pods/%s"                // Display details of a running pod.
	httpRouteServices          = "/services"               // Display a list of running services.
	httpRouteService           = "/services/%s"            // Display details of a running service.

	// Other
	httpRouteGuide = "/guide" // Get information on how to use this application from the server.

	// API Versions
	httpRouteReleasesVersion    = "v1.0.0"
	httpRouteCronjobsVersion    = "v1.0.0"
	httpRouteDeploymentsVersion = "v1.0.0"
	httpRouteIngressesVersion   = "v1.0.0"
	httpRouteJobsVersion        = "v1.0.0"
	httpRoutePodsVersion        = "v1.0.0"
	httpRouteServicesVersion    = "v1.0.0"
	httpRouteGuideVersion       = "v1.0.0"

	httpGet    = "GET"
	httpPatch  = "PATCH"
	httpPost   = "POST"
	httpPut    = "PUT"
	httpDelete = "DELETE"
)
