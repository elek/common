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
	_ = x[ObjectLockEndpointsDisabled-10000]
	_ = x[ObjectLockDisabledForProject-10001]
	_ = x[ObjectLockInvalidBucketState-10002]
	_ = x[ObjectLockBucketRetentionConfigurationMissing-10003]
	_ = x[ObjectLockObjectRetentionConfigurationMissing-10004]
	_ = x[ObjectLockObjectProtected-10005]
	_ = x[ObjectLockInvalidObjectState-10006]
	_ = x[ObjectLockInvalidBucketRetentionConfiguration-10007]
}

const (
	_StatusCode_name_0 = "UnknownOKCanceledInvalidArgumentDeadlineExceededNotFoundAlreadyExistsPermissionDeniedResourceExhaustedFailedPreconditionAbortedOutOfRangeUnimplementedInternalUnavailableDataLossUnauthenticatedMethodNotAllowed"
	_StatusCode_name_1 = "ObjectLockEndpointsDisabledObjectLockDisabledForProjectObjectLockInvalidBucketStateObjectLockBucketRetentionConfigurationMissingObjectLockObjectRetentionConfigurationMissingObjectLockObjectProtectedObjectLockInvalidObjectStateObjectLockInvalidBucketRetentionConfiguration"
)

var (
	_StatusCode_index_0 = [...]uint8{0, 7, 9, 17, 32, 48, 56, 69, 85, 102, 120, 127, 137, 150, 158, 169, 177, 192, 208}
	_StatusCode_index_1 = [...]uint16{0, 27, 55, 83, 128, 173, 198, 226, 271}
)

func (i StatusCode) String() string {
	switch {
	case i <= 17:
		return _StatusCode_name_0[_StatusCode_index_0[i]:_StatusCode_index_0[i+1]]
	case 10000 <= i && i <= 10007:
		i -= 10000
		return _StatusCode_name_1[_StatusCode_index_1[i]:_StatusCode_index_1[i+1]]
	default:
		return "StatusCode(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
