/*
####### sdk.jw (c) 2025 Archivage Numérique ######################################################## MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package jw

import (
	"time"

	"github.com/archnum/sdk.base/failure"
	"github.com/archnum/sdk.base/kv"
	"github.com/archnum/sdk.base/logger"
	"github.com/archnum/sdk.base/logger/level"
	"github.com/archnum/sdk.base/uuid"
)

const (
	_defaultGroup = "default"
	_maxAttempts  = 5
)

//////////////////////
/// JobCore //////////
//////////////////////

type (
	// TODO: un pool ?
	JobCore struct {
		RunAfter    time.Time      `json:"run_after"`
		Config      map[string]any `json:"config"`
		ID          uuid.UUID      `json:"id"`
		Namespace   string         `json:"namespace"`
		Group       string         `json:"group"`
		Type        string         `json:"type"`
		Name        string         `json:"name"`
		UniqueID    string         `json:"unique_id"`
		ExternalID  string         `json:"external_id"`
		Origin      string         `json:"origin"`
		Priority    Priority       `json:"priority"`
		MaxAttempts uint           `json:"max_attempts"`
	}
)

func (jc *JobCore) Validate() error {
	if !jc.ID.Validate() {
		return failure.New("the job identifier is not a UUID", kv.String("id", string(jc.ID))) /////////////////////////
	}

	if jc.Namespace == "" {
		return failure.New("the job namespace cannot be empty") ////////////////////////////////////////////////////////
	}

	if jc.Type == "" {
		return failure.New("the job type cannot be empty") /////////////////////////////////////////////////////////////
	}

	if jc.Group == "" {
		jc.Group = _defaultGroup
	}

	if jc.Name == "" {
		jc.Name = _defaultName
	}

	if jc.Origin == "" {
		jc.Origin = _defaultOrigin
	}

	jc.Priority.Fix()

	if jc.MaxAttempts > _maxAttempts {
		jc.MaxAttempts = _maxAttempts
	}

	return nil
}

func (jc *JobCore) NewJob() *Job {
	return &Job{
		JobCore: jc,
		Status:  StatusTodo,
	}
}

func (jc *JobCore) LogError(err error, logger *logger.Logger, msg string) {
	logger.Error( //::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
		msg,
		kv.String("id", jc.ID.String()),
		kv.String("namespace", jc.Namespace),
		kv.String("group", jc.Group),
		kv.String("type", jc.Type),
		kv.String("name", jc.Name),
		kv.String("origin", jc.Origin),
		kv.String("priority", jc.Priority.String()),
		kv.String("unique_id", jc.UniqueID),
		kv.Time("run_after", jc.RunAfter),
		kv.Uint("max_attempts", uint64(jc.MaxAttempts)),
		kv.String("external_id", jc.ExternalID),
		kv.Error(err),
	)
}

func (jc *JobCore) Log(logger *logger.Logger, level level.Level, msg string) {
	logger.Log( //::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
		level,
		msg,
		kv.String("id", jc.ID.String()),
		kv.String("namespace", jc.Namespace),
		kv.String("group", jc.Group),
		kv.String("type", jc.Type),
		kv.String("name", jc.Name),
		kv.String("origin", jc.Origin),
		kv.String("priority", jc.Priority.String()),
		kv.String("unique_id", jc.UniqueID),
		kv.Time("run_after", jc.RunAfter),
		kv.Uint("max_attempts", uint64(jc.MaxAttempts)),
		kv.String("external_id", jc.ExternalID),
	)
}

//////////////////////
/// Job //////////////
//////////////////////

type (
	// TODO: un pool ?
	Job struct {
		*JobCore
		Status Status `json:"status"`
	}
)

func (job *Job) Log(logger *logger.Logger, level level.Level, msg string) {
	logger.Log( //::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
		level,
		msg,
		kv.String("id", job.ID.String()),
		kv.String("namespace", job.Namespace),
		kv.String("group", job.Group),
		kv.String("type", job.Type),
		kv.String("name", job.Name),
		kv.String("origin", job.Origin),
		kv.String("priority", job.Priority.String()),
		kv.String("unique_id", job.UniqueID),
		kv.Time("run_after", job.RunAfter),
		kv.Uint("max_attempts", uint64(job.MaxAttempts)),
		kv.String("external_id", job.ExternalID),
		kv.String("status", job.Status.String()),
	)
}

/*
####### END ############################################################################################################
*/
