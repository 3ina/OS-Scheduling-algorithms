package scheduler

import "OS_Scheduling_algorithms/io"

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
