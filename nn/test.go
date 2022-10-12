package nn

type TestData struct {
	TestInputData []float64
	TestResult    []float64
}

func NewTestData(testIn, testOut []float64) *TestData {
	return &TestData{
		TestInputData: testIn,
		TestResult:    testOut,
	}
}
