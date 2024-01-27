package main

import (
	"reflect"
	"testing"

	"sandbox3.0/task"
)

func TestCreateNIP(t *testing.T) {
	// Positive test cases
	t.Run("valid akhwat", func(t *testing.T) {
		output, err := task.CreateNIP(task.Akhwat, 2024, 6, 1)
		expected := "ART241-00001"
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if output != expected {
			t.Errorf("Expected %s, got %s", expected, output)
		}
	})
	t.Run("valid ikhwan", func(t *testing.T) {
		output, err := task.CreateNIP(task.Ikhwan, 2024, 12, 99999)
		expected := "ARN242-99999"
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if output != expected {
			t.Errorf("Expected %s, got %s", expected, output)
		}
	})

	// Negative test cases
	t.Run("id > 99999", func(t *testing.T) {
		_, err := task.CreateNIP(task.Ikhwan, 2024, 1, 100000)
		if err == nil {
			t.Error("Expected an error, got nil")
		}
	})
	t.Run("id < 1", func(t *testing.T) {
		_, err := task.CreateNIP(task.Akhwat, 2024, 1, 0)
		if err == nil {
			t.Error("Expected an error, got nil")
		}
	})
}

func TestGenerateNIPs(t *testing.T) {
	// Positive test cases
	t.Run("valid akhwat", func(t *testing.T) {
		output, err := task.GenerateNIPs(task.Akhwat, 2024, 1, 5, 1)
		expected := []string{"ART241-00001", "ART241-00002", "ART241-00003", "ART241-00004", "ART241-00005"}
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if !reflect.DeepEqual(output, expected) {
			t.Errorf("Expected %v, got %v", expected, output)
		}
	})
	t.Run("valid akhwat", func(t *testing.T) {
		output, err := task.GenerateNIPs(task.Ikhwan, 2024, 7, 6, 7)
		expected := []string{"ARN242-00007", "ARN242-00008", "ARN242-00009", "ARN242-00010", "ARN242-00011", "ARN242-00012"}
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if !reflect.DeepEqual(output, expected) {
			t.Errorf("Expected %v, got %v", expected, output)
		}
	})

	// Negative test cases
	t.Run("invalid month", func(t *testing.T) {
		_, err := task.GenerateNIPs(task.Ikhwan, 2024, 13, 1, 1)
		if err == nil {
			t.Error("Expected an error, got nil")
		}
	})
}

func TestCreateNextNIP(t *testing.T) {
	// Positive test cases
	t.Run("valid ikhwan", func(t *testing.T) {
		input := "ARN242-00001"
		expected := "ARN242-00002"
		output, err := task.CreateNextNIP(input)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if output != expected {
			t.Errorf("Expected %s, got %s", expected, output)
		}
	})
	t.Run("valid akhwat", func(t *testing.T) {
		input := "ART241-99998"
		expected := "ART241-99999"
		output, err := task.CreateNextNIP(input)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if output != expected {
			t.Errorf("Expected %s, got %s", expected, output)
		}
	})

	// Negative test cases
	t.Run("invalid NIP", func(t *testing.T) {
		input := "ARX242-00001"
		_, err := task.CreateNextNIP(input)
		if err == nil {
			t.Error("Expected an error, got nil")
		}
	})
}

func TestGenerateNextNIPs(t *testing.T) {
	// Positive test cases
	t.Run("valid ikhwan", func(t *testing.T) {
		nip := "ARN242-00001"
		count := 5
		nips, err := task.GenerateNextNIPs(nip, count)
		expected := []string{"ARN242-00002", "ARN242-00003", "ARN242-00004", "ARN242-00005", "ARN242-00006"}
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if !reflect.DeepEqual(nips, expected) {
			t.Errorf("Expected %v, got %v", expected, nips)
		}
	})
	t.Run("valid akhwat", func(t *testing.T) {
		nip := "ART241-00990"
		count := 6
		nips, err := task.GenerateNextNIPs(nip, count)
		expected := []string{"ART241-00991", "ART241-00992", "ART241-00993", "ART241-00994", "ART241-00995", "ART241-00996"}
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if !reflect.DeepEqual(nips, expected) {
			t.Errorf("Expected %v, got %v", expected, nips)
		}
	})

	// Negative test cases
	t.Run("invalid NIP", func(t *testing.T) {
		nip := "ART243-00000"
		count := 5
		_, err := task.GenerateNextNIPs(nip, count)
		if err == nil {
			t.Error("Expected an error, got nil")
		}
	})
	t.Run("count less than 1", func(t *testing.T) {
		nip := "ART242-99999"
		count := 0
		nips, err := task.GenerateNextNIPs(nip, count)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if len(nips) != 0 {
			t.Errorf("Expected 0 NIPs, got %d", len(nips))
		}
	})
}
