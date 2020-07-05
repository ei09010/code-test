package hash_service

import (
	"code-test/server/services/hash_service"
	"testing"
)

func TestGenerate_ReceivesValidValues_ReturnsExpectedHash(t *testing.T) {

	// Arrange
	expectedHash := "2b8605d8"
	valueToHash := "firstTestInput"

	// Act
	executionHashValue := hash_service.Generate(valueToHash)

	// Assert
	if expectedHash != executionHashValue {
		t.Errorf("Expected %v, got %v", expectedHash, executionHashValue)
	}

}

func TestGenerate_ReceivesEmpty_ReturnsExpectedHash(t *testing.T) {

	// Arrange
	expectedHash := "0"
	valueToHash := ""

	// Act
	executionHashValue := hash_service.Generate(valueToHash)

	// Assert
	if expectedHash != executionHashValue {
		t.Errorf("Expected %v, got %v", expectedHash, executionHashValue)
	}

}
