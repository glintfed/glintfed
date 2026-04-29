package libstr

func ToPtr(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

func FromPtr(s *string) string {
	if s == nil {
		return ""
	}

	return *s
}
