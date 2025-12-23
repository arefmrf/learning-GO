package domain

import "errors"

type HostStatus int

const (
	Draft HostStatus = iota + 1
	Pending
	Rejected
	Approved
	Suspension
	Deleted
)

type CancelLevel int

const (
	Easy CancelLevel = 1
	Normal
	Hard
)

var AllowedTransitions = map[[2]HostStatus]bool{
	{Draft, Pending}:       true,
	{Pending, Rejected}:    true,
	{Pending, Approved}:    true,
	{Rejected, Pending}:    true,
	{Rejected, Approved}:   true,
	{Approved, Suspension}: true,
	{Approved, Deleted}:    true,
	{Suspension, Approved}: true,
	{Suspension, Deleted}:  true,
}

type Host struct {
	Status    HostStatus
	oldStatus HostStatus
}

func (h *Host) ChangeStatus(newStatus HostStatus) error {
	if !AllowedTransitions[[2]HostStatus{h.Status, newStatus}] {
		return errors.New("er2021")
	}
	h.oldStatus = h.Status
	h.Status = newStatus
	return nil
}
