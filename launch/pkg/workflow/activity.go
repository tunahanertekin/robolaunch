package launchflow

import (
	"github.com/robolaunch/robolaunch/launch/pkg/account"
	"github.com/robolaunch/robolaunch/launch/pkg/kubeops"
)

func CreateUserSpace(l LaunchRequest) (string, error) {
	//TODO: Check namespace is available!
	err := kubeops.CheckNamespace(l.Namespace)
	if err != nil {
		// Create namespace
		_, err := kubeops.CreateNamespace(l.Namespace)
		if err != nil {
			return "", err
		}
		//Create role for user
		_, _, err = kubeops.CreateRole(l.Namespace)
		if err != nil {
			return "", err
		}

		//Create keycloak role
		_, err = account.CreateGroup(l.Namespace)
		if err != nil {
			return "", err
		}
		//Bind the role
		err = account.BindGroup(l.Username, l.Namespace)
		if err != nil {
			return "", err
		}
		return "Namespace created: " + l.Namespace, nil
	}
	return "Namespace avaliable", nil
	//TODO: Create Namespace & Role if not available
	//TODO: Create Group and bind them user

}

func CreateLaunch(l LaunchRequest) (LaunchStatus, error) {

	// Launch type not used right now!
	//TODO: Add following functions part as a workflow

	// Check namespace first

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
