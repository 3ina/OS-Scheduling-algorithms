package io

import (
	"encoding/csv"
	"os"
	"strconv"
)

type Process struct {
	Id            int
	ArrivalTime   int
	CpuBurstTime1 int
	IoBurstTime   int
	CpuBurstTime2 int
}

func ReadProcess(filename string) ([]Process, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	var processes []Process
	for _, record := range records {
		id, _ := strconv.Atoi(record[0])
		arrival, _ := strconv.Atoi(record[1])
		cpuBurst1, _ := strconv.Atoi(record[2])
		ioBurst, _ := strconv.Atoi(record[3])
		cpuBurst2, _ := strconv.Atoi(record[4])

		process := Process{
			Id:            id,
			ArrivalTime:   arrival,
			CpuBurstTime1: cpuBurst1,
			IoBurstTime:   ioBurst,
			CpuBurstTime2: cpuBurst2,
		}
		processes = append(processes, process)
	}

	return processes, nil

}
