package main

import (
	"LLD-PRACTICE/file_storage/internal/service"
	"fmt"
)

func main() {
	// first create a folder
	// then create files here
	rootFolder := service.NewFolder("/root")

	familyFolder := service.NewFolder("/children")

	rootFolder.AddChildren(familyFolder)

	str1 := "Hello, Pavan"
	str2 := "Hello, Teja"

	str1ByteContent := []byte(str1)
	str2ByteContent := []byte(str2)

	pavanFile := service.NewFile("Pavan", str1ByteContent)
	tejaFile := service.NewFile("Teja", str2ByteContent)

	familyFolder.AddChildren(pavanFile)
	familyFolder.AddChildren(tejaFile)

	folderContents := rootFolder.GetChildNames()

	fmt.Println("the folder contents is : ", folderContents)

	// delete a file from the folder

	familyFolder.DeleteFileSystem("Pavan")

	folderContents = rootFolder.GetChildNames()

	fmt.Println("the folder contents is : ", folderContents)

}
