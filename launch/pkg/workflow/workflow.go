package launchflow

import (
	"fmt"
	"time"

	"go.temporal.io/sdk/workflow"
)

func LaunchWorkflow(ctx workflow.Context, req LaunchRequest) error {
	launchState := LaunchStatus{
		Username:       req.Username,
		Name:           req.Name,
		Namespace:      req.Namespace,
		LaunchType:     req.LaunchType,
		WorkloadStatus: "CREATING",
		TheiaPort:      0,
		WebRpcPort:     0,
		NodeIp:         "",
	}
	// Search attribute to find workflows

	options := workflow.ActivityOptions{
		ScheduleToCloseTimeout: time.Minute,
	}
	ctx = workflow.WithActivityOptions(ctx, options)

	// Set query handler for status...
	err := workflow.SetQueryHandler(ctx, "getLaunch", func() (LaunchStatus, error) {
		return launchState, nil
	})
	if err != nil {
		return err
	}
	// Execute CreateDeployment & Create Service Section
	var namespaceStatus string
	err = workflow.ExecuteActivity(ctx, CreateUserSpace, req).Get(ctx, &namespaceStatus)
	if err != nil {
		return err
	}
	fmt.Println(namespaceStatus)

	err = workflow.ExecuteActivity(ctx, CreateLaunch, req).Get(ctx, &launchState)
	if err != nil {
		return err
	}
	signalVal := LaunchRequest{}
	signalName := "CHANGE_LAUNCH"
	signalChan := workflow.GetSignalChannel(ctx, signalName)
	s := workflow.NewSelector(ctx)
	s.AddReceive(signalChan, func(c workflow.ReceiveChannel, more bool) {
		c.Receive(ctx, &signalVal)
		workflow.GetLogger(ctx).Info("Received signal!", "Signal", signalName, "value", signalVal)
		if signalVal.Operation == "DELETE" {
			err = workflow.ExecuteActivity(ctx, DeleteLaunch, signalVal).Get(ctx, &launchState)
			if err != nil {
				fmt.Println(err)
			}
		}
		if signalVal.Operation == "STOP" {
			err = workflow.ExecuteActivity(ctx, ScaleOut, signalVal).Get(ctx, &launchState.WorkloadStatus)
			if err != nil {
				launchState.WorkloadStatus = "FAILED"
				fmt.Println(err)
			}
		}
		if signalVal.Operation == "START" {
			err = workflow.ExecuteActivity(ctx, ScaleUp, signalVal).Get(ctx, &launchState.WorkloadStatus)
			if err != nil {
				launchState.WorkloadStatus = "FAILED"
				fmt.Println(err)

			}
		}
	})
	for {
		s.Select(ctx)
		if signalVal.Operation == "DELETE" {
			return nil
		}

	}
}
