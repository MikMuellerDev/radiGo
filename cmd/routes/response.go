package routes

type ResponseStruct struct {
	Success   bool
	ErrorCode int
	Title     string
	Message   string
}

type StatusStruct struct {
	Mode string
}
