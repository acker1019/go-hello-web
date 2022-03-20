package validator

func IsBinary(operands []interface{}) bool {
	return len(operands) == 2
}
