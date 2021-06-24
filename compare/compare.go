package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// Mencari array value yang berbeda
func Difference(a, b []string) (diff []string) {
	m := make(map[string]bool)

	for _, item := range b {
			m[item] = true
	}

	for _, item := range a {
			if _, ok := m[item]; !ok {
					diff = append(diff, item)
			}
	}
	return
}

// Mencari array value yang sama
func Intersection(a, b []string) (c []string) {
	m := make(map[string]bool)

	for _, item := range a {
			m[item] = true
	}

	for _, item := range b {
			if _, ok := m[item]; ok {
					c = append(c, item)
			}
	}
	return
}


// Function mendapatkan list directory dan file pada target folder yang diinputkan
func getDir(dir string) ([]string, error){
	var directory []string 

	// digunakan untuk mendapatkan list folder dan file
	err := filepath.Walk(dir,
	func (path string, info os.FileInfo, err error) error {
			if err != nil {
				return  err
			}

			pathreplace := strings.Replace(dir,"./","",-1)
			// membuat list directory dan file
			directory = append(directory, strings.Replace(path,pathreplace,"",-1))
			return nil
	})


	if err != nil {
		return nil, err
	}

	return directory, nil
}


// TASK #1. Implementasikan sebuah *program* yang membandingkan isi dari dua direktori melalui *parameter*. 
func compare(path1 string, path2 string)  {

	var finalResult[]string;

	dir1, err := getDir(path1)	

	if err!=nil{
		fmt.Println(err)
	}

	dir2, err := getDir(path2)	
	if err!=nil{
		fmt.Println(err)
	}

	
	new := Difference(dir1, dir2);
	deleted := Difference(dir2, dir1);


	for _, item := range new {
		finalResult = append(finalResult, item+" NEW")
	}

	for _, item := range deleted {
		finalResult = append(finalResult, item+" DELETED")
	}
		


	sort.Strings(finalResult)

	for _, result := range finalResult {
		fmt.Println(result)
	}
	

}


// TASK #2. Modifikasi program #1 untuk compare file content untuk rule (1), jika ada perbedaan beri keterangan MODIFIED
func compare_2(path1 string, path2 string)  {

	var finalResult[]string;

	dir1, err := getDir(path1)	

	if err!=nil{
		fmt.Println(err)
	}

	dir2, err := getDir(path2)	
	if err!=nil{
		fmt.Println(err)
	}

	
	modified := Intersection(dir1, dir2);


	for _, item := range modified {
		finalResult = append(finalResult, item+" MODIFIED")
	}

		
	sort.Strings(finalResult)

	for _, result := range finalResult {
		fmt.Println(result)
	}
	

}



func main() {
	/**
	TASK #1. Implementasikan sebuah *program* yang membandingkan isi dari dua direktori melalui *parameter*. 
	1. Jika file ada di source dan target, abaikan
	2. Jika file ada di source tapi tidak ada di target berikan keterangan **NEW**
	3. Jika file tidak ada di source tapi ada di target berikan keterangan **DELETED**
	*/
	
	/**
	TASK #2. Modifikasi program #1 untuk compare file content untuk rule (1), jika ada perbedaan beri keterangan MODIFIED
	*/
	
	
	var menu int;

	fmt.Println("MENU")
	fmt.Println("TASK #1. Implementasikan sebuah *program* yang membandingkan isi dari dua direktori melalui *parameter*.")
	fmt.Println("	- Jika file ada di source dan target, abaikan")
	fmt.Println("	- Jika file ada di source tapi tidak ada di target berikan keterangan **NEW**")
	fmt.Println("	- Jika file tidak ada di source tapi ada di target berikan keterangan **DELETED**")
	fmt.Println("TASK #2. Modifikasi program #1 untuk compare file content untuk rule (1), jika ada perbedaan beri keterangan MODIFIED")
	
	fmt.Print("Pilih Menu : ")
	fmt.Scan(&menu)

	if menu == 1 {

		compare("./source/","./target/")	

	} else if menu == 2 {

		compare_2("./source/","./target/")

	} else {

		fmt.Println("Menu tidak ditemukan")

	}

}