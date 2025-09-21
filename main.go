package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	// Memeriksa apakah argumen command line diberikan
	if len(os.Args) < 2 {
		fmt.Println("❌ Error: Path ke file input tidak diberikan!")
		fmt.Println("📖 Penggunaan: go run main.go <path_to_input_file>")
		fmt.Println("📝 Contoh: go run main.go input/angka.txt")
		os.Exit(1)
	}

	inputPath := os.Args[1]
	fmt.Printf("🚀 Memulai program sorting dengan file input: %s\n", inputPath)

	// Membaca file input
	numbers, err := readInputFile(inputPath)
	if err != nil {
		fmt.Printf("❌ Error membaca file input: %v\n", err)
		os.Exit(1)
	}

	if len(numbers) == 0 {
		fmt.Println("⚠️  Peringatan: File input kosong atau tidak ada angka yang valid!")
		return
	}

	fmt.Printf("📊 Berhasil membaca %d angka dari file input\n", len(numbers))
	fmt.Printf("📋 Angka sebelum diurutkan: %v\n", numbers)

	// Mengurutkan angka menggunakan fungsi SortNumbers
	fmt.Println("🔄 Mengurutkan angka...")
	sortedNumbers := SortNumbers(numbers)

	fmt.Printf("✅ Angka setelah diurutkan: %v\n", sortedNumbers)

	// Membuat folder output jika belum ada
	outputDir := "output"
	if err := createOutputDirectory(outputDir); err != nil {
		fmt.Printf("❌ Error membuat folder output: %v\n", err)
		os.Exit(1)
	}

	// Membuat nama file output
	outputFileName := generateOutputFileName(inputPath)
	outputPath := filepath.Join(outputDir, outputFileName)

	// Menulis hasil ke file output
	if err := writeOutputFile(outputPath, sortedNumbers); err != nil {
		fmt.Printf("❌ Error menulis file output: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("💾 Hasil berhasil disimpan ke: %s\n", outputPath)
	fmt.Println("🎉 Program selesai!")
}

// readInputFile membaca file input baris per baris dan mengkonversi ke integer
func readInputFile(filePath string) ([]int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("tidak dapat membuka file %s: %v", filePath, err)
	}
	defer file.Close()

	var numbers []int
	scanner := bufio.NewScanner(file)
	lineNumber := 0

	for scanner.Scan() {
		lineNumber++
		line := strings.TrimSpace(scanner.Text())

		// Abaikan baris kosong
		if line == "" {
			continue
		}

		// Konversi string ke integer
		num, err := strconv.Atoi(line)
		if err != nil {
			fmt.Printf("⚠️  Peringatan: Baris %d bukan angka yang valid: '%s' - diabaikan\n", lineNumber, line)
			continue
		}

		numbers = append(numbers, num)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error membaca file: %v", err)
	}

	return numbers, nil
}

// createOutputDirectory membuat folder output jika belum ada
func createOutputDirectory(dirPath string) error {
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		fmt.Printf("📁 Membuat folder output: %s\n", dirPath)
		return os.MkdirAll(dirPath, 0755)
	}
	return nil
}

// generateOutputFileName membuat nama file output berdasarkan nama file input
func generateOutputFileName(inputPath string) string {
	baseName := filepath.Base(inputPath)
	ext := filepath.Ext(baseName)
	nameWithoutExt := strings.TrimSuffix(baseName, ext)
	
	// Jika file sudah memiliki suffix _sorted, tambahkan _sorted lagi
	if strings.HasSuffix(nameWithoutExt, "_sorted") {
		return nameWithoutExt + ext
	}
	
	return nameWithoutExt + "_sorted" + ext
}

// writeOutputFile menulis slice of integers ke file output
func writeOutputFile(filePath string, numbers []int) error {
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("tidak dapat membuat file %s: %v", filePath, err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	// Menulis setiap angka ke file, satu angka per baris
	for _, num := range numbers {
		_, err := writer.WriteString(fmt.Sprintf("%d\n", num))
		if err != nil {
			return fmt.Errorf("error menulis ke file: %v", err)
		}
	}

	fmt.Printf("📝 Menulis %d angka ke file output\n", len(numbers))
	return nil
}
