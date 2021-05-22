package constant

var NotAuthorizedErrorTemplate = &errorValue{
	Code:           1,
	Message:        "Not authorized",
	DetailTemplate: "%s",
}

type errorValue struct {
	Code           int
	Message        string
	DetailTemplate string
}
