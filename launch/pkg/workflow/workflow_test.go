package launchflow

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"

	"go.temporal.io/sdk/testsuite"
)

type UnitTestSuite struct {
	suite.Suite
	testsuite.WorkflowTestSuite

	env *testsuite.TestWorkflowEnvironment
}

func (s *UnitTestSuite) SetupTest() {
	s.env = s.NewTestWorkflowEnvironment()
}

func (s *UnitTestSuite) AfterTest(suiteName, testName string) {
	s.env.AssertExpectations(s.T())
}

type IntegrationTestSuite struct {
	suite.Suite
	testsuite.WorkflowTestSuite

	env *testsuite.TestWorkflowEnvironment
}

func (s *IntegrationTestSuite) SetupTest() {
	s.env = s.NewTestWorkflowEnvironment()
}

func (s *IntegrationTestSuite) AfterTest(suiteName, testName string) {
	s.env.AssertExpectations(s.T())
}

func (s *UnitTestSuite) Test_CreateWorkflow() {
	//TODO: Add persistent token to run workflow
	launch := LaunchRequest{
		Name:       "testname",
		Namespace:  "testplace",
		LaunchType: "Test-Type",
		Username:   "test_name",
		Operation:  "CREATE",
	}
	launchStatus := LaunchStatus{}

	s.env.ExecuteWorkflow(LaunchWorkflow, launch)
	res, err := s.env.QueryWorkflow("getLaunch")
	s.NoError(err)
	err = res.Get(&launchStatus)
	s.NoError(err)

	s.Equal("CREATING", launchStatus.WorkloadStatus)

	s.True(s.env.IsWorkflowCompleted())

}

func (s *UnitTestSuite) Test_CreateLaunch() {
	//TODO: Add persistent token to run workflow
	launch := LaunchRequest{
		Name:       "testname",
		Namespace:  "testplace",
		LaunchType: "Test-Type",
		Username:   "test_name",
		Operation:  "CREATE",
	}
	launchStatus := LaunchStatus{}

	s.env.OnActivity(CreateLaunch, launch).Return(func(req LaunchRequest) (LaunchStatus, error) {
		return LaunchStatus{
			Name:           req.Name,
			Namespace:      req.Namespace,
			Username:       req.Username,
			WorkloadStatus: "RUNNING",
		}, nil
	})
	s.env.ExecuteWorkflow(LaunchWorkflow, launch)
	res, err := s.env.QueryWorkflow("getLaunch")
	s.NoError(err)
	err = res.Get(&launchStatus)
	s.NoError(err)

	s.Equal("RUNNING", launchStatus.WorkloadStatus)

	s.True(s.env.IsWorkflowCompleted())

	fmt.Println(launchStatus.WorkloadStatus)
}

func (s *UnitTestSuite) Test_DeleteLaunch() {
	//TODO: Add persistent token to run workflow
	launch := LaunchRequest{
		Name:       "testname",
		Namespace:  "testplace",
		LaunchType: "Test-Type",
		Username:   "test_name",
		Operation:  "DELETE",
	}
	launchStatus := LaunchStatus{}

	s.env.OnActivity(CreateLaunch, launch).Return(func(req LaunchRequest) (LaunchStatus, error) {
		return LaunchStatus{
			Name:           req.Name,
			Namespace:      req.Namespace,
			Username:       req.Username,
			WorkloadStatus: "RUNNING",
		}, nil
	})
	s.env.OnActivity(DeleteLaunch, launch).Return(func(req LaunchRequest) (LaunchStatus, error) {
		if req.Operation == "DELETE" {
			return LaunchStatus{
				Name:           req.Name,
				Namespace:      req.Namespace,
				Username:       req.Username,
				WorkloadStatus: "DELETED",
				TheiaPort:      0,
				WebRpcPort:     0,
				NodeIp:         "",
			}, nil
		}
		return LaunchStatus{}, errors.New("Error happend")

	})

	s.env.RegisterDelayedCallback(func() {

		s.env.SignalWorkflow("CHANGE_LAUNCH", launch)

	}, time.Millisecond*1)
	s.env.ExecuteWorkflow(LaunchWorkflow, launch)
	res, err := s.env.QueryWorkflow("getLaunch")
	s.NoError(err)
	err = res.Get(&launchStatus)
	s.NoError(err)
	s.Equal("DELETED", launchStatus.WorkloadStatus)
	s.True(s.env.IsWorkflowCompleted())

}

func (s *UnitTestSuite) Test_StopLaunch() {
	//TODO: Add persistent token to run workflow
	launch := LaunchRequest{
		Name:       "testname",
		Namespace:  "testplace",
		LaunchType: "Test-Type",
		Username:   "test_name",
		Operation:  "STOP",
	}
	launchStatus := LaunchStatus{}

	s.env.OnActivity(CreateLaunch, launch).Return(func(req LaunchRequest) (LaunchStatus, error) {
		return LaunchStatus{
			Name:           req.Name,
			Namespace:      req.Namespace,
			Username:       req.Username,
			WorkloadStatus: "RUNNING",
		}, nil
	})
	s.env.OnActivity(ScaleOut, launch).Return(func(req LaunchRequest) (string, error) {
		if req.Operation == "STOP" {
			return "STOPPED", nil
		}
		return "", errors.New("Error happend")

	})

	s.env.RegisterDelayedCallback(func() {

		s.env.SignalWorkflow("CHANGE_LAUNCH", launch)

	}, time.Millisecond*1)
	s.env.ExecuteWorkflow(LaunchWorkflow, launch)
	res, err := s.env.QueryWorkflow("getLaunch")
	s.NoError(err)
	err = res.Get(&launchStatus)
	s.NoError(err)
	s.Equal("STOPPED", launchStatus.WorkloadStatus)
	s.True(s.env.IsWorkflowCompleted())

}

func (s *UnitTestSuite) Test_StartLaunch() {
	// Scenario
	// Create flow
	// Automatically run
	// Stop it
	// Start Again
	launch := LaunchRequest{
		Name:       "testname",
		Namespace:  "testplace",
		LaunchType: "Test-Type",
		Username:   "test_name",
		Operation:  "RUNNING",
	}
	launchStatus := LaunchStatus{}

	s.env.OnActivity(CreateLaunch, launch).Return(func(req LaunchRequest) (LaunchStatus, error) {
		return LaunchStatus{
			Name:           req.Name,
			Namespace:      req.Namespace,
			Username:       req.Username,
			WorkloadStatus: "RUNNING",
		}, nil
	})
	s.env.OnActivity(ScaleUp, launch).Return(func(req LaunchRequest) (string, error) {
		if req.Operation == "START" {
			return "RUNNING", nil
		}
		return "", errors.New("Error happend")

	})
	s.env.OnActivity(ScaleOut, launch).Return(func(req LaunchRequest) (string, error) {
		if req.Operation == "DELETE" {
			return "STOPPED", nil
		}
		return "", errors.New("Error happend")

	})

	s.env.RegisterDelayedCallback(func() {
		stopReq := launch
		stopReq.Operation = "STOP"
		s.env.SignalWorkflow("CHANGE_LAUNCH", stopReq)

	}, time.Millisecond*1)
	s.env.RegisterDelayedCallback(func() {
		startReq := launch
		startReq.Operation = "STOP"

		s.env.SignalWorkflow("CHANGE_LAUNCH", startReq)

	}, time.Millisecond*1)
	s.True(s.env.IsWorkflowCompleted())

	s.env.ExecuteWorkflow(LaunchWorkflow, launch)
	res, err := s.env.QueryWorkflow("getLaunch")
	s.NoError(err)
	err = res.Get(&launchStatus)
	s.NoError(err)
	s.Equal("RUNNING", launchStatus.WorkloadStatus)

}

func TestUnitTestSuite(t *testing.T) {
	suite.Run(t, new(UnitTestSuite))
}
