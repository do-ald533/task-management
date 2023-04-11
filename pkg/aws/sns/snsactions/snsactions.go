package snsactions

type SNS_ACTIONS string

const (
	CREATION     SNS_ACTIONS = "CREATION"
	UPDATE       SNS_ACTIONS = "UPDATE"
	REMOVAL      SNS_ACTIONS = "REMOVAL"
	CANCELLATION SNS_ACTIONS = "CANCELLATION"
)
