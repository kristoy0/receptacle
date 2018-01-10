package main

import (
	"errors"
	"sort"

	unit "github.com/docker/go-units"
	proto "github.com/kristoy0/receptacle/server/proto"
)

type Endpoint struct {
	Address string
	Memory  int64
	CPU     int
	Disk    uint
}

type Result struct {
	Endpoint Endpoint
	Result   int64
}

type Results []Result

func PlaceContainer(req *proto.DeployRequest, endpoints []Endpoint) (Result, error) {
	results := Results{}

	for _, e := range endpoints {
		cpu := req.Resources.CPU
		if cpu == 0 || req.Resources.Memory == "" {
			results = append(results, Result{e, 0})
			continue
		}
		mem, err := unit.FromHumanSize(req.Resources.Memory)
		if err != nil {
			return Result{}, err
		}
		if e.Memory < mem || float32(e.CPU) < cpu {
			continue
		}
		memTotal := e.Memory / mem * 100
		cpuTotal := float32(e.CPU) / cpu * 100
		// diskTotal := resultCalc(e.Disk, uint(req.Resources.Disk))

		total := memTotal + int64(cpuTotal) // + diskTotal
		results = append(results, Result{e, total})
	}

	if len(results) == 0 {
		return Result{}, errors.New("Not enough resources on any hosts")
	}

	sort.Slice(results, func(i, j int) bool {
		return results[i].Result > results[j].Result
	})
	return results[0], nil
}
