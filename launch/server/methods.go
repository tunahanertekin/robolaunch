package main

import (
	"context"
	"log"

	launchPb "github.com/robolaunch/robolaunch/api"
)

func (s *server) CreateLaunch(ctx context.Context, in *launchPb.CreateRequest) (*launchPb.LaunchState, error) {
	log.Printf("---CreateLaunch---")

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

func (s *server) OperateLaunch(ctx context.Context, in *launchPb.OperateRequest) (*launchPb.LaunchState, error) {
	log.Printf("---OperateLaunch---")

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
