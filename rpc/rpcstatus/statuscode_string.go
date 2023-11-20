// Code generated by "stringer -type StatusCode ."; DO NOT EDIT.

package rpcstatus

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Unknown-0]
	_ = x[OK-1]
	_ = x[Canceled-2]
	_ = x[InvalidArgument-3]
	_ = x[DeadlineExceeded-4]
	_ = x[NotFound-5]
	_ = x[AlreadyExists-6]
	_ = x[PermissionDenied-7]
	_ = x[ResourceExhausted-8]
	_ = x[FailedPrecondition-9]
	_ = x[Aborted-10]
	_ = x[OutOfRange-11]
	_ = x[Unimplemented-12]
	_ = x[Internal-13]
	_ = x[Unavailable-14]
	_ = x[DataLoss-15]
	_ = x[Unauthenticated-16]
	_ = x[MethodNotAllowed-17]
}

const _StatusCode_name = "UnknownOKCanceledInvalidArgumentDeadlineExceededNotFoundAlreadyExistsPermissionDeniedResourceExhaustedFailedPreconditionAbortedOutOfRangeUnimplementedInternalUnavailableDataLossUnauthenticatedMethodNotAllowed"

var _StatusCode_index = [...]uint8{0, 7, 9, 17, 32, 48, 56, 69, 85, 102, 120, 127, 137, 150, 158, 169, 177, 192, 208}

func (i StatusCode) String() string {
	if i >= StatusCode(len(_StatusCode_index)-1) {
		return "StatusCode(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _StatusCode_name[_StatusCode_index[i]:_StatusCode_index[i+1]]
}
