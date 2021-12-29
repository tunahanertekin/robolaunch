package main

import (
	"context"
	"log"
	"os"

	"github.com/google/uuid"
	launchPb "github.com/robolaunch/robolaunch/api"
	launchflow "github.com/robolaunch/robolaunch/launch/pkg/workflow"
	"go.temporal.io/sdk/client"
	"google.golang.org/grpc/metadata"
)

func (s *server) CreateLaunch(ctx context.Context, in *launchPb.CreateRequest) (*launchPb.LaunchState, error) {
	//Getting id token from grpc metadata
	headers, _ := metadata.FromIncomingContext(ctx)

	idToken := headers["token-jwt"][0]
	searchAttributes := map[string]interface{}{
		"DeploymentName":      in.Name,
		"DeploymentNamespace": in.Namespace,
	}
	log.Printf("---CreateLaunch---")
	//TODO: Run Workflow!
	c, err := client.NewClient(client.Options{
		HostPort: os.Getenv("TEMPORAL_SERVER_IP"),
	})
	if err != nil {
		return nil, err
	}
	options := client.StartWorkflowOptions{
		ID:               uuid.New().String(),
		TaskQueue:        launchflow.LaunchQueue,
		SearchAttributes: searchAttributes,
	}
	we, err := c.ExecuteWorkflow(ctx, options, launchflow.LaunchWorkflow, launchflow.LaunchRequest{
		Username:   in.GetUsername(),
		Name:       in.GetName(),
		Namespace:  in.GetNamespace(),
		IDToken:    idToken,
		Operation:  in.GetOperation(),
		LaunchType: in.GetLaunchType(),
	})
	if err != nil {
		return nil, err
	}
	//TODO: Query given Workflow

	resp, err := c.QueryWorkflow(context.Background(), we.GetID(), we.GetRunID(), "getLaunch")
	if err != nil {
		return nil, err
	}

	var status launchflow.LaunchStatus
	if err = resp.Get(&status); err != nil {
		return nil, err
	}

	return &launchPb.LaunchState{
		Username:       status.Username,
		Namespace:      status.Namespace,
		Name:           status.Name,
		LaunchType:     status.Namespace,
		WorkloadStatus: status.WorkloadStatus,
		TheiaPort:      status.TheiaPort,
		WebrtcPort:     status.WebRpcPort,
		NodeIp:         status.NodeIp,
	}, nil
}

func (s *server) OperateLaunch(ctx context.Context, in *launchPb.OperateRequest) (*launchPb.LaunchState, error) {
	// log.Printf("---OperateLaunch---")
	// //Getting id token from grpc metadata
	// headers, _ := metadata.FromIncomingContext(ctx)

	// idToken := headers["token-jwt"][0]

	// TODO: Find given workflowId / RunID
	// From workflow list examine Advanced Query Api
	// TODO: Send Start & Stop signal according to incoming request

	// //

	return &launchPb.LaunchState{
		Username:       "",
		Namespace:      "",
		Name:           "",
		LaunchType:     "",
		WorkloadStatus: "",
		TheiaPort:      0,
		WebrtcPort:     0,
		NodeIp:         "",
	}, nil
}

// func (s *server) ListLaunch(ctx context.Context, in *launchPb.Empty) (*launchPb.LaunchList, error) {
// 	log.Printf("---OperateLaunch---")

// 	return &launchPb.LaunchList{
// 		Username:       "",
// 		Namespace:      "",
// 		Name:           "",
// 		LaunchType:     "",
// 		WorkloadStatus: "",
// 	}, nil
// }
