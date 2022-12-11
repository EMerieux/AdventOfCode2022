package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		log.Fatalf("Gros naze: %v", e)
	}
}

type File struct {
	Name string
	Size int
}

type Folder struct {
	ParentFullPath string
	FullPath       string
	Name           string
	Folders        []*Folder
	Files          []File
	TotalSize      int
}

func getFullPath(currentFolderFullPath string, folderName string) string {
	if currentFolderFullPath == "/" {
		return currentFolderFullPath + folderName
	}
	return currentFolderFullPath + "/" + folderName
}

func checkAndCreateFolder(currentFolderFullPath string, folderName string, folders map[string]*Folder) {
	fullPath := getFullPath(currentFolderFullPath, folderName)
	if _, ok := folders[fullPath]; !ok {
		folders[fullPath] = &Folder{
			ParentFullPath: currentFolderFullPath,
			FullPath:       fullPath,
			Name:           folderName,
		}
	}
}

func printFolder(folder *Folder, tabLevel int) int {
	tabs := ""
	totalSize := 0
	for i := 0; i < tabLevel; i++ {
		tabs += "\t"
	}
	fmt.Println(tabs + "Folder Name: " + folder.Name)
	for _, f := range folder.Folders {
		totalSize += printFolder(f, tabLevel+1)
	}
	for _, file := range folder.Files {
		fmt.Println(tabs + "\t" + strconv.Itoa(file.Size) + " " + file.Name)
		totalSize += file.Size
	}
	folder.TotalSize = totalSize
	fmt.Println(tabs + "\tTOTAL SIZE: " + strconv.Itoa(totalSize))
	return totalSize
}

func main() {
	f, err := os.Open("input")
	check(err)
	defer f.Close()

	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)

	folders := make(map[string]*Folder, 0)
	folders["/"] = &Folder{
		FullPath: "/",
		Name:     "/",
	}
	currentFolder := folders["/"]

	for fileScanner.Scan() {
		line := fileScanner.Text()
		if string(line[0]) == "$" {
			if line[2:4] == "cd" {
				folderName := line[5:]
				if folderName == ".." {
					currentFolder = folders[currentFolder.ParentFullPath]
				} else if folderName != "/" {
					checkAndCreateFolder(currentFolder.FullPath, folderName, folders)
					fullPath := getFullPath(currentFolder.FullPath, folderName)
					currentFolder.Folders = append(currentFolder.Folders, folders[fullPath])
					currentFolder = folders[fullPath]
				}
			}
		} else if line[0:3] == "dir" {
			folderName := line[4:]
			checkAndCreateFolder(currentFolder.FullPath, folderName, folders)
		} else {
			parts := strings.Split(line, " ")
			fileSize, err := strconv.Atoi(parts[0])
			check(err)
			currentFolder.Files = append(currentFolder.Files, File{
				Name: parts[1],
				Size: fileSize,
			})
		}
	}

	if root, ok := folders["/"]; ok {
		printFolder(root, 0)
	}

	systemSize := 70000000
	sum := 0
	totalUsedMemory := 0
	dirSizes := make([]int, 0)
	for k, folder := range folders {
		if folder.TotalSize <= 100000 {
			sum += folder.TotalSize
		}
		if k == "/" {
			totalUsedMemory = folder.TotalSize
		}
		dirSizes = append(dirSizes, folder.TotalSize)
	}
	println("SUM: " + strconv.Itoa(sum))
	println("Total used memory: " + strconv.Itoa(totalUsedMemory))
	freeMemory := systemSize - totalUsedMemory
	println("Free memory: " + strconv.Itoa(freeMemory))
	requiredFreeMemory := 30000000
	missingFreeMemory := requiredFreeMemory - freeMemory
	sort.Ints(dirSizes)
	for _, size := range dirSizes {
		if size > missingFreeMemory {
			fmt.Println("Folder size to delete: " + strconv.Itoa(size))
			break
		}
	}
}
