package main

var CommonErr = ErrorCode{
	Code: 1000000,
}

var Success = ErrorCode{
	Code:    200,
	Message: "Success",
}

var InternalServerErr = ErrorCode{
	Code:    500,
	Message: "Internal Server Err",
}

var FormatErr = ErrorCode{
	Code:    1004001,
	Message: "Parameter err, %v",
}
