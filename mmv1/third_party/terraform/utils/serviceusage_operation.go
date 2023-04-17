package google

import (
	"encoding/json"
	"time"

	"google.golang.org/api/googleapi"
	"google.golang.org/api/serviceusage/v1"
)

func serviceUsageOperationWait(config *transport_tpg.Config, op *serviceusage.Operation, project, activity, userAgent string, timeout time.Duration) error {
	// maintained for compatibility with old code that was written before the
	// autogenerated waiters.
	b, err := op.MarshalJSON()
	if err != nil {
		return err
	}
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return ServiceUsageOperationWaitTime(config, m, project, activity, userAgent, timeout)
}

func handleServiceUsageRetryableError(err error) error {
	if err == nil {
		return nil
	}
	if gerr, ok := err.(*googleapi.Error); ok {
		if (gerr.Code == 400 || gerr.Code == 412) && gerr.Message == "Precondition check failed." {
			return &googleapi.Error{
				Code:    503,
				Message: "api returned \"precondition failed\" while enabling service",
			}
		}
	}
	return err
}
