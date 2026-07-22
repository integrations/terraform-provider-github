package value

import "time"

type Kind string

const (
	KindText         Kind = "text"
	KindNumber       Kind = "number"
	KindDate         Kind = "date"
	KindSingleSelect Kind = "single_select"
	KindIteration    Kind = "iteration"
)

type Result struct {
	Kind        Kind
	Text        string
	Number      float64
	Date        time.Time
	OptionID    string
	IterationID string
}

type SetInput struct {
	ProjectID string
	ItemID    string
	FieldID   string
	Value     Result
}
