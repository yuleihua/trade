package errcode

const SuccessMsg = "success"
const (
	SuccessCode   = 9999
	RiskTradeCode = 6666
)

type Error interface {
	Error() string // returns the error message
	Code() int     // returns the error code
}

type SuccessError struct{ Message string }

func (e *SuccessError) Code() int     { return 9999 }
func (e *SuccessError) Error() string { return SuccessMsg }

type InvalidRequestError struct{ Message string }

func (e *InvalidRequestError) Code() int     { return 1100 }
func (e *InvalidRequestError) Error() string { return e.Message }

type InvalidTokenError struct{ Message string }

func (e *InvalidTokenError) Code() int     { return 1101 }
func (e *InvalidTokenError) Error() string { return e.Message }

type InvalidCacheError struct{ Message string }

func (e *InvalidCacheError) Code() int     { return 1200 }
func (e *InvalidCacheError) Error() string { return e.Message }

type InvalidDatabaseError struct{ Message string }

func (e *InvalidDatabaseError) Code() int     { return 1300 }
func (e *InvalidDatabaseError) Error() string { return e.Message }

type InvalidMQError struct{ Message string }

func (e *InvalidMQError) Code() int     { return 1400 }
func (e *InvalidMQError) Error() string { return e.Message }

type InvalidMetricError struct{ Message string }

func (e *InvalidMetricError) Code() int     { return 1500 }
func (e *InvalidMetricError) Error() string { return e.Message }

type InvalidTracingError struct{ Message string }

func (e *InvalidTracingError) Code() int     { return 1600 }
func (e *InvalidTracingError) Error() string { return e.Message }

type InvalidInternalError struct{ Message string }

func (e *InvalidInternalError) Code() int     { return 1900 }
func (e *InvalidInternalError) Error() string { return e.Message }

//-------------------biz-------------------------------------------------------

type InvalidParamsError struct{ Message string }

func (e *InvalidParamsError) Code() int     { return 2100 }
func (e *InvalidParamsError) Error() string { return e.Message }

type InvalidAccountError struct{ Message string }

func (e *InvalidAccountError) Code() int     { return 2200 }
func (e *InvalidAccountError) Error() string { return e.Message }

type InvalidFundError struct{ Message string }

func (e *InvalidFundError) Code() int     { return 2300 }
func (e *InvalidFundError) Error() string { return e.Message }

type InvalidRiskError struct{ Message string }

func (e *InvalidRiskError) Code() int     { return 2400 }
func (e *InvalidRiskError) Error() string { return e.Message }

type InvalidBizError struct{ Message string }

func (e *InvalidBizError) Code() int     { return 2900 }
func (e *InvalidBizError) Error() string { return e.Message }

//-----------------------------------------------------------------------------

const InitSequenceCode = 3000

type InvalidInitSequenceError struct{ Message string }

func (e *InvalidInitSequenceError) Code() int     { return InitSequenceCode }
func (e *InvalidInitSequenceError) Error() string { return e.Message }
