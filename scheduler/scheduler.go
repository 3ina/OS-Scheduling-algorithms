package scheduler

import (
	"OS_Scheduling_algorithms/io"
)

type State string

const (
	New        State = "New"
	Ready      State = "Ready"
	Running    State = "Running"
	Waiting    State = "Waiting"
	Terminated State = "Terminated"
)

type ProcessWithState struct {
	Process        io.Process
	State          State
	WaitTime       int
	ResponseTime   int
	TurnaroundTime int
	StartTime      int
	EndTime        int
	Logs           []string
}

type Metrics struct {
	AverageTurnaroundTime float64
	AverageWaitingTime    float64
	AverageResponseTime   float64
	Throughput            float64
	Utilization           float64
}

type Scheduler struct {
	processes []ProcessWithState
	time      int
	quantum   int
	totalTime int
	logs      []string
}

func NewScheduler(processes []io.Process, quantum int) *Scheduler {
	var processWithState []ProcessWithState
	for _, p := range processes {
		processWithState = append(processWithState, ProcessWithState{
			Process: p,
			State:   New,
		})
	}
	return &Scheduler{
		processes: processWithState,
		time:      0,
		totalTime: 0,
		quantum:   quantum,
		logs:      []string{},
	}
}
