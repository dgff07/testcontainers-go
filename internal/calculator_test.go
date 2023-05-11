package internal

import (
	"testing"

	"github.com/dgff07/testcontainers-go/internal/mocks"
	"github.com/dgff07/testcontainers-go/internal/tests"

	_ "github.com/lib/pq"
)

func TestMain(m *testing.M) {
	tests.InitContainers(m)
}

func TestSum(t *testing.T) {
	tests.DB.Ping()
	// Create a new instance of the mock PrintAPI
	mockPrintAPI := &mocks.PrintAPI{}

	// Create a new instance of the Calculator struct with the mock PrintAPI
	calculator := &Calculator{printAPI: mockPrintAPI}

	// Set the expectation that the PrintSomething method will be called once
	mockPrintAPI.On("PrintSomething").Once()

	// Call the Sum method
	result := calculator.Sum(1, 2)

	// Assert that the result is correct
	if result != 3 {
		t.Errorf("Unexpected result: %v", result)
	}

	// Assert that the PrintSomething method was called as expected
	mockPrintAPI.AssertExpectations(t)
}
