package scheduler

import (
	"OS_Scheduling_algorithms/io"
	"fmt"
	"sort"
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

func (s *Scheduler) ioCompletion(p *ProcessWithState) {
	s.logStateChange(p, Ready)
}

func (s *Scheduler) terminate(p *ProcessWithState) {
	s.logStateChange(p, Terminated)
	p.EndTime = s.time
	p.TurnaroundTime = p.EndTime - p.Process.ArrivalTime
	p.WaitTime = p.TurnaroundTime - p.Process.CpuBurstTime1 - p.Process.IoBurstTime - p.Process.CpuBurstTime2
}

func (s *Scheduler) RunFCFS() []ProcessWithState {
	var result []ProcessWithState
	for i, _ := range s.processes {
		process := &s.processes[i]
		if s.time < process.Process.ArrivalTime {
			s.time = process.Process.ArrivalTime
		}
		s.dispatch(process)
		s.time += process.Process.CpuBurstTime1 + process.Process.CpuBurstTime2 + process.Process.IoBurstTime
		s.terminate(process)
		result = append(result, *process)

	}
	return result
}

func (s *Scheduler) RunSJF() []ProcessWithState {
	var result []ProcessWithState
	remaining := make([]*ProcessWithState, len(s.processes))

	for i, _ := range s.processes {
		remaining[i] = &s.processes[i]
	}

	for len(remaining) > 0 {
		ready := []*ProcessWithState{}

		for _, process := range remaining {
			if process.Process.ArrivalTime <= s.time {
				ready = append(ready, process)
			}
		}

		sort.Slice(ready, func(i, j int) bool {
			return ready[i].Process.CpuBurstTime1 < ready[j].Process.CpuBurstTime1
		})

		if len(ready) == 0 {
			s.time = remaining[0].Process.ArrivalTime
			continue
		}

		process := ready[0]

		for i, p := range remaining {
			if p == process {
				remaining = append(remaining[:i], remaining[i+1:]...)
				break
			}
		}

		s.dispatch(process)
		s.time += process.Process.CpuBurstTime1
		s.time += process.Process.IoBurstTime
		s.time += process.Process.CpuBurstTime2
		s.terminate(process)
		result = append(result, *process)

	}

	return result
}
