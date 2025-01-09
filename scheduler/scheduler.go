package scheduler

import (
	"OS_Scheduling_algorithms/io"
	"fmt"
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

func (s *Scheduler) logStateChange(p *ProcessWithState, newState State) {
	timestamp := fmt.Sprintf("Time %d: Process %d moved to %s state", s.time, p.Process.Id, newState)
	p.Logs = append(p.Logs, timestamp)
	s.logs = append(s.logs, timestamp)
	p.State = newState

}

func (s *Scheduler) admit(p *ProcessWithState) {
	s.logStateChange(p, Ready)
}

func (s *Scheduler) dispatch(p *ProcessWithState) {
	if p.StartTime == 0 {
		p.StartTime = s.time
		p.ResponseTime = s.time - p.Process.ArrivalTime
	}
	s.logStateChange(p, Running)
}

func (s *Scheduler) preempt(p *ProcessWithState) {
	s.logStateChange(p, Ready)
}

func (s *Scheduler) ioRequest(p *ProcessWithState) {
	s.logStateChange(p, Waiting)
	s.time += p.Process.IoBurstTime
}
