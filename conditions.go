package winenv

// Equals with other
func (s WindowsEnvInfo) Equals(other WindowsEnvInfo) bool {
	return s.String() == other.String()
}

// Newer than other
func (s WindowsEnvInfo) Newer(other WindowsEnvInfo) bool {
	return s.compareValue() > other.compareValue()
}

// Older than other
func (s WindowsEnvInfo) Older(other WindowsEnvInfo) bool {
	return s.compareValue() < other.compareValue()
}

// NewerOrEqual with other
func (s WindowsEnvInfo) NewerOrEqual(other WindowsEnvInfo) bool {
	return s.Equals(other) || s.Newer(other)
}

// OlderOrEqual with other
func (s WindowsEnvInfo) OlderOrEqual(other WindowsEnvInfo) bool {
	return s.Equals(other) || s.Older(other)
}
