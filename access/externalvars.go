package access

// ExternalStruct is strcture to check extrnal access
type ExternalStruct struct {
	internalVal int
	ExternalVal int
}

// ExVar exposes values to outside
var ExVar ExternalStruct

func init() {
	ExVar.internalVal = 10
	ExVar.ExternalVal = 100
}
