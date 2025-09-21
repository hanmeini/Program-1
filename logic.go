package main

// SortNumbers mengurutkan slice of integers dari kecil ke besar
// Parameter: numbers - slice of integers yang akan diurutkan
// Return: slice of integers baru yang sudah terurut dari kecil ke besar
func SortNumbers(numbers []int) []int {
	// Membuat copy dari slice input untuk menghindari modifikasi slice asli
	sortedNumbers := make([]int, len(numbers))
	copy(sortedNumbers, numbers)
	
	// Menggunakan bubble sort untuk mengurutkan
	n := len(sortedNumbers)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if sortedNumbers[j] > sortedNumbers[j+1] {
				// Menukar posisi jika elemen saat ini lebih besar dari elemen berikutnya
				sortedNumbers[j], sortedNumbers[j+1] = sortedNumbers[j+1], sortedNumbers[j]
			}
		}
	}
	
	return sortedNumbers
}
