package request

// 创建操作日志
type SysOperLogInsertRequest struct {
	Title         string
	BusinessType  int
	Method        string
	RequestMethod string
	OperatorType  int
	OperName      string
	DeptName      string
	OperUrl       string
	OperIp        string
	OperLocation  string
	OperParam     string
	JsonResult    string
	Status        string
	ErrorMsg      string
	CostTime      int
}
