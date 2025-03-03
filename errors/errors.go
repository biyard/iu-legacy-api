package errors

import (
	"net/http"

	"biyard.co/common/base"
)

var (
	ErrRandomGeneration      = base.NewError(1000, http.StatusInternalServerError, "random generation error")
	ErrAccountAlreadyExists  = base.NewError(1001, http.StatusBadRequest, "account already exists")
	ErrInvalidCode           = base.NewError(1002, http.StatusBadRequest, "invalid code")
	ErrFailedToUpdateAddress = base.NewError(1003, http.StatusInternalServerError, "failed to update address")

	ErrVerificationAlreadyExists = base.NewError(1004, http.StatusBadRequest, "verification key already exists")
	ErrVerificationNotExists     = base.NewError(1005, http.StatusBadRequest, "verification key not exists")
	ErrVerificationIsUsed        = base.NewError(1005, http.StatusBadRequest, "verification key is already used")
	ErrTransactionSend           = base.NewError(1006, http.StatusBadRequest, "transaction failed")

	ErrAlreadyPlay       = base.NewError(1006, http.StatusBadRequest, "User is already play")
	ErrUserAlreadyExists = base.NewError(1007, http.StatusBadRequest, "user is already exists")
	ErrUserNotExists     = base.NewError(1008, http.StatusBadRequest, "user is not exists")
	ErrDeleteFailed      = base.NewError(1009, http.StatusBadRequest, "deleted failed")

	ErrInsertFailed = base.NewError(1010, http.StatusBadRequest, "insert failed")
	ErrProductEmpty = base.NewError(1011, http.StatusBadRequest, "product is empty")
	ErrUserEmpty    = base.NewError(1011, http.StatusBadRequest, "user is empty")

	ErrMsg                           = base.NewError(1012, http.StatusBadRequest, "error is happening")
	ErrAllowedError                  = base.NewError(1013, http.StatusBadRequest, "get allowed failed")
	ErrUpdateFailed                  = base.NewError(1014, http.StatusBadRequest, "update failed")
	ErrNotExists                     = base.NewError(1015, http.StatusBadRequest, "data not exists")
	ErrGetClosedTokenClaimInfoFailed = base.NewError(1016, http.StatusBadRequest, "get closedTokenClaim info failed")
	ErrGetAllDataFailed              = base.NewError(1017, http.StatusBadRequest, "data get failed")

	ErrContractCallFailed = base.NewError(1018, http.StatusBadRequest, "contract call failed")

	ErrMissingRequiredParameters = base.NewError(1019, http.StatusBadRequest, "missing required parameters")

	ErrInsertStartNumberFailed = base.NewError(1020, http.StatusBadRequest, "insert start number failed")

	ErrUpdateStartNumberFailed = base.NewError(1021, http.StatusBadRequest, "update start number failed")
)
