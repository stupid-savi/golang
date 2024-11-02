package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
)

func main() {
	file, err := os.Open("example.txt")

	if err != nil {
		panic(err)
	}

	defer file.Close()
	fileInfo, e := file.Stat()
	if e != nil {
		panic(e)
	}

	fmt.Println("filename: ", fileInfo.Name())
	fmt.Println("size: ", fileInfo.Size())
	fmt.Println("folder: ", fileInfo.IsDir())
	fmt.Println("modified_at", fileInfo.ModTime())
	fmt.Println("permissions", fileInfo.Mode())
	fileSize := fileInfo.Size()

	buffer := make([]byte, fileSize)

	fileLength, err := file.Read(buffer)

	if err != nil {
		panic(err)
	}

	fmt.Println("File", buffer, fileLength)

	for _, char := range buffer {
		fmt.Println("Character : ", string(char))
	}

	// Easy way

	fileData, err := os.ReadFile("example.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(fileData)) // Easy Peasy But only use for small files becaise it try to load entire data in the memory at once so for large files in GBs This lead to resource limitation issue

	// Read folders and files

	dir, err := os.Open("../")
	if err != nil {
		panic(err)
	}

	fileInfos, err := dir.ReadDir(-1) // >= 0  TO VIEW ALL FILES AND FOLDERS

	for _, file := range fileInfos {
		if file.IsDir() {
			fmt.Println("Folder: ", file)

		}
	}

	// Create a File

	f, err := os.Create("example2.txt")
	if err != nil {
		panic(err)
	}

	f.WriteString("Hi Go ")
	f.WriteString("Hi Golang ")

	// not appending but replacing
	fileName, err := os.Open("example2.txt")
	if err != nil {
		panic(err)
	}

	fileNameData, err := fileName.Stat()
	fileMode := fs.FileMode(0644)

	if err == nil {
		fileMode = fileNameData.Mode()
	}
	if err != nil {

		panic(err)
	}
	os.WriteFile("example2.txt", []byte("Bye Bye Go"), fileMode)
	defer f.Close()

	bytes := []byte("Bytes Data")
	f.Write(bytes)

	// Streaming Files

	sourceFile, err := os.Open("source-file.txt")
	if err != nil {
		panic(err)
	}

	defer sourceFile.Close()

	destFile, err := os.Create("destination-file.txt")
	if err != nil {
		panic(err)
	}

	defer destFile.Close()

	reader := bufio.NewReader(sourceFile)
	writer := bufio.NewWriter(destFile)

	for {

		b, err := reader.ReadByte()

		if err != nil {
			if err.Error() != "EOF" {
				panic(err)
			}

			break
		}

		e := writer.WriteByte(b)

		if e != nil {
			panic(e)
		}

	}

	writer.Flush()

	fmt.Println("Writeing data from one file to other via streaming Done")

	// deleting a file

	removeErr := os.Remove("test.txt")

	if removeErr != nil {
		panic(removeErr)
	}

	fmt.Println("File Deleted Successfully")

}
