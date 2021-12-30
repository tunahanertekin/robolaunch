package launchflow

import (
	"github.com/robolaunch/robolaunch/launch/pkg/kubeops"
)

func CreateLaunch(l LaunchRequest) (LaunchStatus, error) {

	// Launch type not used right now!
	udpPort, theiaPort, err := kubeops.CreateDeploymentService(l.Name, l.Namespace, l.IDToken)
	if err != nil {
		return LaunchStatus{}, err
	}
	return LaunchStatus{
		Name:           l.Name,
		Namespace:      l.Namespace,
		LaunchType:     "",
		WorkloadStatus: "RUNNING",
		NodeIp:         "", // TODO: Add Get Node IP ops function
		TheiaPort:      theiaPort,
		WebRpcPort:     udpPort,
	}, nil
}

func DeleteLaunch(l LaunchRequest) (LaunchStatus, error) {

	err := kubeops.DeleteDeploymentService(l.Name, l.Namespace, l.IDToken)
	if err != nil {
		return LaunchStatus{}, err
	}
	return LaunchStatus{
		Name:           l.Name,
		Namespace:      l.Namespace,
		LaunchType:     "",
		WorkloadStatus: "DELETED",
		NodeIp:         "", // For a moment it would be static
		TheiaPort:      0,
		WebRpcPort:     0,
	}, nil

}

func ScaleOut(l LaunchRequest) (string, error) {
	err := kubeops.ScaleDeployment(l.Name, l.Namespace, 0, l.IDToken)
	if err != nil {
		return "", nil
	}
	return "STOPPED", nil
}

func ScaleUp(l LaunchRequest) (string, error) {
	err := kubeops.ScaleDeployment(l.Name, l.Namespace, 1, l.IDToken)
	if err != nil {
		return "", nil
	}
	return "RUNNING", nil
}
