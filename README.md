# Number Sorter

A simple yet powerful command-line tool written in Go that reads numbers from a text file, sorts them in ascending order, and outputs the sorted result to a new file.

## Features

- **File Input Processing**: Reads numbers from text files line by line
- **Intelligent Parsing**: Automatically skips empty lines and handles invalid entries gracefully
- **Bubble Sort Algorithm**: Implements efficient bubble sort for number sorting
- **Automatic Output Management**: Creates output directory and generates appropriately named output files
- **Comprehensive Error Handling**: Robust error handling with clear user feedback
- **CLI Interface**: Simple command-line interface with helpful usage instructions
- **Unit Testing**: Complete test suite covering various scenarios including edge cases

## Project Structure

```
STS/
├── input/                    # Input directory
│   └── angka.txt            # Sample input file with numbers
├── output/                  # Output directory (auto-created)
│   └── angka_sorted.txt     # Generated sorted output file
├── logic.go                 # Core sorting logic and SortNumbers function
├── logic_test.go           # Comprehensive unit tests
├── main.go                 # CLI application entry point
├── go.mod                  # Go module definition
└── README.md               # Project documentation
```

### File Descriptions

- **`main.go`**: Main application entry point with CLI interface and file I/O operations
- **`logic.go`**: Contains the `SortNumbers` function that implements bubble sort algorithm
- **`logic_test.go`**: Unit tests covering various test cases including random numbers, duplicates, empty slices, and edge cases
- **`input/angka.txt`**: Sample input file containing numbers to be sorted
- **`output/`**: Directory where sorted results are saved (automatically created)

## Usage

### Prerequisites

- Go 1.19 or later installed on your system
- Basic understanding of command-line operations

### Running the Program

#### Method 1: Using `go run` (Recommended for development)

```bash
# Navigate to the project directory
cd STS

# Run the program with input file
go run main.go input/angka.txt
```

#### Method 2: Building and Running Executable

```bash
# Build the executable
go build -o number-sorter main.go

# Run the executable
./number-sorter input/angka.txt

# On Windows
number-sorter.exe input/angka.txt
```

### Input File Format

The input file should contain one number per line:

```
11000
2420
123
45
245
245
334
```

### Expected Output

The program will:
1. Read numbers from the input file
2. Sort them in ascending order
3. Save the sorted result to `output/` directory
4. Display progress and results in the console

Example console output:
```
🚀 Memulai program sorting dengan file input: input/angka.txt
📊 Berhasil membaca 7 angka dari file input
📋 Angka sebelum diurutkan: [11000 2420 123 45 245 245 334]
🔄 Mengurutkan angka...
✅ Angka setelah diurutkan: [45 123 245 245 334 2420 11000]
📁 Membuat folder output: output
📝 Menulis 7 angka ke file output
💾 Hasil berhasil disimpan ke: output/angka_sorted.txt
🎉 Program selesai!
```

### Running Tests

```bash
# Run all tests
go test

# Run tests with verbose output
go test -v

# Run tests with coverage
go test -cover
```

## Error Handling

The application implements comprehensive error handling to ensure robust operation:

### Input Validation
- **Missing Arguments**: Displays usage instructions if no input file path is provided
- **File Not Found**: Clear error message when input file doesn't exist
- **Permission Issues**: Handles file access permission errors gracefully

### Data Processing
- **Invalid Numbers**: Skips non-numeric lines with warning messages
- **Empty Lines**: Automatically ignores empty lines in input file
- **Empty Files**: Handles cases where input file contains no valid numbers

### File Operations
- **Output Directory**: Automatically creates output directory if it doesn't exist
- **File Writing**: Handles write permission and disk space issues
- **Path Resolution**: Properly handles file path resolution across different operating systems

### Error Messages
The application provides user-friendly error messages with emojis for better visibility:

- ❌ **Critical Errors**: File not found, permission denied, write failures
- ⚠️ **Warnings**: Invalid data lines, empty input files
- 📁 **Info Messages**: Directory creation, file operations
- ✅ **Success Messages**: Successful operations and completion

### Example Error Scenarios

```bash
# Missing argument
$ go run main.go
❌ Error: Path ke file input tidak diberikan!
📖 Penggunaan: go run main.go <path_to_input_file>

# File not found
$ go run main.go nonexistent.txt
❌ Error membaca file input: tidak dapat membuka file nonexistent.txt: open nonexistent.txt: no such file or directory

# Invalid data in file
⚠️  Peringatan: Baris 3 bukan angka yang valid: 'abc' - diabaikan
```

## Development

### Adding New Features
1. Implement core logic in `logic.go`
2. Add corresponding tests in `logic_test.go`
3. Update CLI interface in `main.go` if needed
4. Run tests to ensure functionality: `go test -v`

### Testing
The project includes comprehensive unit tests covering:
- Random number sorting
- Already sorted arrays
- Reverse sorted arrays
- Duplicate numbers
- Empty slices
- Single elements
- Negative numbers
- Original array immutability

## License

This project is open source and available under the MIT License.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
