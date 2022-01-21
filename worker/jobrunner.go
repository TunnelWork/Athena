package main

import (
	"errors"
	"time"

	"github.com/TunnelWork/Athena/protobuf"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var (
	ErrJobExpired = errors.New("JobRunner: job has expired and is therefore rejected")
)

type JobRunner struct {
	Ipv4Capable bool
	Ipv6Capable bool
}

func (jr *JobRunner) Run(job *protobuf.Job) (*protobuf.Response, error) {
	var resp = &protobuf.Response{
		Uuid:       job.Uuid,
		Relays:     []string{},
		Results:    []*protobuf.Result{},
		AcceptedAt: timestamppb.Now(),
	}

	// First, check if the Job is expired
	if job.GetDeadline().AsTime().Before(time.Now()) {
		return nil, ErrJobExpired
	}

	resp.ReportedAt = timestamppb.Now()

	return resp, nil
}
