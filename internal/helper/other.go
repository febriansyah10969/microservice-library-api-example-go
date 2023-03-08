package helper

// Generic Function to return address value of variable
func GetAddress[V int | float64 | string](val V) *V {
	return &val
}