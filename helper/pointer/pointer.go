package helper

// ToStringPtr returns a nil string pointer if s is empty.
func ToStringPtr(s string) *string {
	if len(s) == 0 {
		return nil
	}

	return &s
}
