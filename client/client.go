package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Client represents an instance of a connection to the server.
type Client struct {
	Token string `json:"bearerToken"` // The API authorization token to the server.
	Url   string `json:"URL"`         // The URL to the server endpoint.
}

type DeployRequest struct {
	Memo       string `json:"memo"`       // Optional text to display in slack etc.
	Name       string `json:"name"`       // The application/chart name to deploy.
	Namespace  string `json:"namespace"`  // The namespace to deploy.
	VersionTag string `json:"versionTag"` // The docker version tag.
}

type RestartRequest struct {
	Namespace string `json:"namespace"` // The namespace where the deployment is running.
}

type RollbackRequest struct {
	Revision string `json:"revision"` // The revision to roll back to (optional)
}

// New is a factory function that returns a new client instance.
func NewClient(url string, apiToken string) *Client {
	return &Client{
		Token: apiToken,
		Url:   url,
	}
}

// Version prints the version of the client.
func (c *Client) Version() string {
	return fmt.Sprintf("%s version %s\n", applicationName, version)
}

// Helm related commands.

// Delete removes a deployed release from the cluster.
func (c *Client) Delete(release string) (*Response, error) {
	// Send the request.
	req, err := http.NewRequest(httpDelete, fmt.Sprintf("%s%s", c.Url, fmt.Sprintf(httpRouteRelease, release)), nil)
	if err != nil {
		return nil, err
	}
	return c.sendRequest(req, "releases", httpRouteReleasesVersion)
}

// Deploy submits a deploy request to the server.
func (c *Client) Deploy(name string, versionTag string, namespace string, memo string) (*Response, error) {
	// Create the payload.
	dr := &DeployRequest{
		Memo:       memo,
		Name:       name,
		Namespace:  namespace,
		VersionTag: versionTag,
	}
	payload, err := json.Marshal(dr)
	if err != nil {
		return nil, err
	}

	// Send the request.
	req, err := http.NewRequest(httpPost, fmt.Sprintf("%s%s", c.Url, httpRouteReleases),
		bytes.NewBuffer([]byte(payload)))
	if err != nil {
		return nil, err
	}
	return c.sendRequest(req, "releases", httpRouteReleasesVersion)
}

// History prints out the detail historical activity for a release.
func (c *Client) History(release string, format string) (*Response, error) {
	req, err := http.NewRequest(httpGet, fmt.Sprintf("%s%s", c.Url, fmt.Sprintf(httpRouteReleaseHistory, release)), nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	q.Add("f", format)
	req.URL.RawQuery = q.Encode()
	return c.sendRequest(req, "releases", httpRouteReleasesVersion)
}

// List prints out a list of releases.
func (c *Client) List(namespace string, format string) (*Response, error) {
	req, err := http.NewRequest(httpGet, fmt.Sprintf("%s%s", c.Url, httpRouteReleases), nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	q.Add("n", namespace)
	q.Add("f", format)
	req.URL.RawQuery = q.Encode()
	return c.sendRequest(req, "releases", httpRouteReleasesVersion)
}

// Rollback removes a deployed release froms the cluster and restarts the previous one in history.
func (c *Client) Rollback(release string, revision string) (*Response, error) {
	// Create the payload.
	dr := &RollbackRequest{
		Revision: revision,
	}
	payload, err := json.Marshal(dr)
	if err != nil {
		return nil, err
	}

	// Send the request.
	req, err := http.NewRequest(httpPut, fmt.Sprintf("%s%s", c.Url, fmt.Sprintf(httpRouteReleaseRollback, release)),
		bytes.NewBuffer([]byte(payload)))
	if err != nil {
		return nil, err
	}
	return c.sendRequest(req, "releases", httpRouteReleasesVersion)
}

// Status gets the details of a release.
func (c *Client) Status(release string, format string) (*Response, error) {
	req, err := http.NewRequest(httpGet, fmt.Sprintf("%s%s", c.Url, fmt.Sprintf(httpRouteRelease, release)), nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	q.Add("f", format)
	req.URL.RawQuery = q.Encode()
	return c.sendRequest(req, "releases", httpRouteReleasesVersion)
}

// Kube related commands

// Configmap prints out the details of a configmap.
func (c *Client) Configmap(name string, namespace string) (*Response, error) {
	req, err := http.NewRequest(httpGet, fmt.Sprintf("%s%s", c.Url, fmt.Sprintf(httpRouteConfigmap, name)), nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	q.Add("n", namespace)
	req.URL.RawQuery = q.Encode()
	return c.sendRequest(req, "configmaps", httpRouteConfigmapsVersion)
}

// Configmaps prints out a list of configmaps.
func (c *Client) Configmaps(namespace string, format string) (*Response, error) {
	req, err := http.NewRequest(httpGet, fmt.Sprintf("%s%s", c.Url, httpRouteConfigmaps), nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	q.Add("n", namespace)
	q.Add("f", format)
	req.URL.RawQuery = q.Encode()
	return c.sendRequest(req, "configmaps", httpRouteConfigmapsVersion)
}

// Cronjob prints out the details of a running cronjob.
func (c *Client) Cronjob(name string, namespace string) (*Response, error) {
	req, err := http.NewRequest(httpGet, fmt.Sprintf("%s%s", c.Url, fmt.Sprintf(httpRouteCronjob, name)), nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	q.Add("n", namespace)
	req.URL.RawQuery = q.Encode()
	return c.sendRequest(req, "cronjobs", httpRouteCronjobsVersion)
}

// Cronjobs prints out a list of cronjobs.
func (c *Client) Cronjobs(namespace string, format string) (*Response, error) {
	req, err := http.NewRequest(httpGet, fmt.Sprintf("%s%s", c.Url, httpRouteCronjobs), nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	q.Add("n", namespace)
	q.Add("f", format)
	req.URL.RawQuery = q.Encode()
	return c.sendRequest(req, "cronjobs", httpRouteCronjobsVersion)
}

// Deployment prints out the details of a deployment.
func (c *Client) Deployment(name string, namespace string) (*Response, error) {
	req, err := http.NewRequest(httpGet, fmt.Sprintf("%s%s", c.Url, fmt.Sprintf(httpRouteDeployment, name)), nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	q.Add("n", namespace)
	req.URL.RawQuery = q.Encode()
	return c.sendRequest(req, "deployments", httpRouteDeploymentsVersion)
}

// Deployments prints out a list of deployments.
func (c *Client) Deployments(namespace string, format string) (*Response, error) {
	req, err := http.NewRequest(httpGet, fmt.Sprintf("%s%s", c.Url, httpRouteDeployments), nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	q.Add("n", namespace)
	q.Add("f", format)
	req.URL.RawQuery = q.Encode()
	return c.sendRequest(req, "deployments", httpRouteDeploymentsVersion)
}

// DeploymentRestart restarts all pods in a deployment
func (c *Client) DeploymentRestart(name string, namespace string) (*Response, error) {
	// Create the payload.
	dr := &RestartRequest{
		Namespace: namespace,
	}
	payload, err := json.Marshal(dr)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(httpPatch, fmt.Sprintf("%s%s", c.Url, fmt.Sprintf(httpRouteDeploymentRestart, name)),
		bytes.NewBuffer([]byte(payload)))
	if err != nil {
		return nil, err
	}
	return c.sendRequest(req, "deployments", httpRouteDeploymentsVersion)
}

// Ingress prints out the details of an ingress.
func (c *Client) Ingress(name string, namespace string) (*Response, error) {
	req, err := http.NewRequest(httpGet, fmt.Sprintf("%s%s", c.Url, fmt.Sprintf(httpRouteIngress, name)), nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	q.Add("n", namespace)
	req.URL.RawQuery = q.Encode()
	return c.sendRequest(req, "ingresses", httpRouteIngressesVersion)
}

// Ingresses prints out a list of ingresses.
func (c *Client) Ingresses(namespace string, format string) (*Response, error) {
	req, err := http.NewRequest(httpGet, fmt.Sprintf("%s%s", c.Url, httpRouteIngresses), nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	q.Add("n", namespace)
	q.Add("f", format)
	req.URL.RawQuery = q.Encode()
	return c.sendRequest(req, "ingresses", httpRouteIngressesVersion)
}

// Job prints out the details of a running job.
func (c *Client) Job(name string, namespace string) (*Response, error) {
	req, err := http.NewRequest(httpGet, fmt.Sprintf("%s%s", c.Url, fmt.Sprintf(httpRouteJob, name)), nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	q.Add("n", namespace)
	req.URL.RawQuery = q.Encode()
	return c.sendRequest(req, "jobs", httpRouteJobsVersion)
}

// Jobs prints out a list of jobs.
func (c *Client) Jobs(namespace string, format string) (*Response, error) {
	req, err := http.NewRequest(httpGet, fmt.Sprintf("%s%s", c.Url, httpRouteJobs), nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	q.Add("n", namespace)
	q.Add("f", format)
	req.URL.RawQuery = q.Encode()
	return c.sendRequest(req, "jobs", httpRouteJobsVersion)
}

// Pod prints out the details of a running pod.
func (c *Client) Pod(name string, namespace string) (*Response, error) {
	req, err := http.NewRequest(httpGet, fmt.Sprintf("%s%s", c.Url, fmt.Sprintf(httpRoutePod, name)), nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	q.Add("n", namespace)
	req.URL.RawQuery = q.Encode()
	return c.sendRequest(req, "pods", httpRoutePodsVersion)
}

// Pods prints out a list of pods.
func (c *Client) Pods(namespace string, format string) (*Response, error) {
	req, err := http.NewRequest(httpGet, fmt.Sprintf("%s%s", c.Url, httpRoutePods), nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	q.Add("n", namespace)
	q.Add("f", format)
	req.URL.RawQuery = q.Encode()
	return c.sendRequest(req, "pods", httpRoutePodsVersion)
}

// Service prints out details of a service.
func (c *Client) Service(name string, namespace string) (*Response, error) {
	req, err := http.NewRequest(httpGet, fmt.Sprintf("%s%s", c.Url, fmt.Sprintf(httpRouteService, name)), nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	q.Add("n", namespace)
	req.URL.RawQuery = q.Encode()
	return c.sendRequest(req, "services", httpRouteServicesVersion)
}

// Services prints out a list of services.
func (c *Client) Services(namespace string, format string) (*Response, error) {
	req, err := http.NewRequest(httpGet, fmt.Sprintf("%s%s", c.Url, httpRouteServices), nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	q.Add("n", namespace)
	q.Add("f", format)
	req.URL.RawQuery = q.Encode()
	return c.sendRequest(req, "services", httpRouteServicesVersion)
}

// Guide retrieves the user guide from the server.
func (c *Client) Guide() (*Response, error) {
	req, err := http.NewRequest(httpGet, fmt.Sprintf("%s%s", c.Url, httpRouteGuide), nil)
	if err != nil {
		return nil, err
	}

	return c.sendRequest(req, "guide", httpRouteGuideVersion)
}

// Private Methods.

// Response back from the server.
type Response struct {
	Status  string `json:"status"`  // Status response from the server.
	Message string `json:"message"` // Message and or data back from the server.
}

// sendRequest adds some metadata, sends the request to the server, and returns the response.
func (c *Client) sendRequest(req *http.Request, resource string, apiVersion string) (*Response, error) {
	req.Header.Add("Accept", fmt.Sprintf("application/vnd.%s.%s-%s+json", serverName, resource, apiVersion))
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.Token))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-Request-ID", createV4UUID()) // For logging/sync purposes.

	cl := &http.Client{}
	resp, err := cl.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result Response
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
