package test

type TestComment interface {
	// Comment of Add
	Add(a, b int) int
	// Comment of Test
	Test(s string) error
}
