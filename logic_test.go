package main

import (
	"reflect"
	"testing"
)

// TestSortNumbers_RandomNumbers menguji fungsi SortNumbers dengan angka acak
func TestSortNumbers_RandomNumbers(t *testing.T) {
	input := []int{64, 34, 25, 12, 22, 11, 90}
	expected := []int{11, 12, 22, 25, 34, 64, 90}
	
	result := SortNumbers(input)
	
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("SortNumbers(%v) = %v; expected %v", input, result, expected)
	}
}

// TestSortNumbers_AlreadySorted menguji fungsi SortNumbers dengan angka yang sudah terurut
func TestSortNumbers_AlreadySorted(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	expected := []int{1, 2, 3, 4, 5}
	
	result := SortNumbers(input)
	
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("SortNumbers(%v) = %v; expected %v", input, result, expected)
	}
}

// TestSortNumbers_ReverseSorted menguji fungsi SortNumbers dengan angka terurut terbalik
func TestSortNumbers_ReverseSorted(t *testing.T) {
	input := []int{5, 4, 3, 2, 1}
	expected := []int{1, 2, 3, 4, 5}
	
	result := SortNumbers(input)
	
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("SortNumbers(%v) = %v; expected %v", input, result, expected)
	}
}

// TestSortNumbers_DuplicateNumbers menguji fungsi SortNumbers dengan angka duplikat
func TestSortNumbers_DuplicateNumbers(t *testing.T) {
	input := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3}
	expected := []int{1, 1, 2, 3, 3, 4, 5, 5, 6, 9}
	
	result := SortNumbers(input)
	
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("SortNumbers(%v) = %v; expected %v", input, result, expected)
	}
}

// TestSortNumbers_EmptySlice menguji fungsi SortNumbers dengan slice kosong
func TestSortNumbers_EmptySlice(t *testing.T) {
	input := []int{}
	expected := []int{}
	
	result := SortNumbers(input)
	
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("SortNumbers(%v) = %v; expected %v", input, result, expected)
	}
}

// TestSortNumbers_SingleElement menguji fungsi SortNumbers dengan satu elemen
func TestSortNumbers_SingleElement(t *testing.T) {
	input := []int{42}
	expected := []int{42}
	
	result := SortNumbers(input)
	
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("SortNumbers(%v) = %v; expected %v", input, result, expected)
	}
}

// TestSortNumbers_NegativeNumbers menguji fungsi SortNumbers dengan angka negatif
func TestSortNumbers_NegativeNumbers(t *testing.T) {
	input := []int{-3, -1, -4, -1, -5, 0, 2, -6}
	expected := []int{-6, -5, -4, -3, -1, -1, 0, 2}
	
	result := SortNumbers(input)
	
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("SortNumbers(%v) = %v; expected %v", input, result, expected)
	}
}

// TestSortNumbers_OriginalNotModified menguji bahwa slice asli tidak dimodifikasi
func TestSortNumbers_OriginalNotModified(t *testing.T) {
	original := []int{3, 1, 4, 1, 5}
	originalCopy := make([]int, len(original))
	copy(originalCopy, original)
	
	SortNumbers(original)
	
	if !reflect.DeepEqual(original, originalCopy) {
		t.Errorf("Original slice was modified. Original: %v, Expected: %v", original, originalCopy)
	}
}
