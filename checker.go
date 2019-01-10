package checker

// Checker is the interface for health checking
type Checker interface {
	// Check returns error if health checking was failed
	Check() error
}
