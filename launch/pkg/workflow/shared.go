package launchflow

const LaunchQueue = "LAUNCH_QUEUE"

type LaunchRequest struct {
	Username   string
	Name       string
	LaunchType string
	Namespace  string
	IDToken    string
	Operation  string
}

type LaunchStatus struct {
	Username       string
	Name           string
	LaunchType     string
	Namespace      string
	WorkloadStatus string
	TheiaPort      int32
	WebRpcPort     int32
	NodeIp         string
}
