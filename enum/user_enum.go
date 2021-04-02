package enum

type ResponseType int

const  (
	OperateOK ResponseType = 0
	OperateFail ResponseType = 1
)

func (r ResponseType) String() string {
	switch r {
	case OperateOK :
		return "OK"
	case OperateFail:
		return "Fail"
	default:
		return "UNKNOWN"
	}
}
