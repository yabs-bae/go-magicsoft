package main

import (
	"fmt"
)

// Function Mendapatkan maximal value pada array
func maxIntSlice(v []int) int {
	var max int
    for i := 1; i < len(v); i++ {
        if max < v[i] {
            max = v[i]
        }
    }
	return max
}


// Function Membuat grapic
func verticalBar(numbers []int){

	max := maxIntSlice(numbers)

	for i := 0; i < max; i++ {
		fmt.Println(" ")
		for j := 0; j < len(numbers); j++ {
			space := max - numbers[j]
			if space > i {
				fmt.Print(" ")
			}else{
				fmt.Print("|")
			}
			
		}
	}

	fmt.Println(" ")

	for _, nums := range numbers {
		fmt.Print(nums)
	}
	fmt.Println(" ")


}

// Sort menggunakan metode Bubble sort
func sortAsc(numbers []int) {
	verticalBar(numbers)
	fmt.Println(" ")
	n := len(numbers);  
	temp := 0;  
	for i :=0; i < n; i++{  
		for j :=1; j < (n-i); j++{  
			if numbers[j-1] > numbers[j] {  
				temp = numbers[j-1];  
				numbers[j-1] = numbers[j];  
				numbers[j] = temp;  
			}  

		}  
		verticalBar(numbers)
		fmt.Println(" ")
	}  

}


// Sort menggunakan metode inverse Bubble sort
func sortDesc(numbers []int){
	verticalBar(numbers)
	fmt.Println(" ")
	
	n := len(numbers);  
	temp := 0;

	for i := 0; i < ( n - 1  ); i++ {
		for j := 0	; j < n - i -1  ; j++ {
			if numbers[j] < numbers[j+1] {
				temp = numbers[j];
				numbers[j] = numbers[j+1];
				numbers[j+1] = temp;
			}
		}

		verticalBar(numbers)
		fmt.Println(" ")

	}
}

func main(){
	var numbers = []int{
		1, 4, 5, 6, 8, 2,
	}

	var menu int;

	fmt.Println("MENU")
	fmt.Println("1. Design dan implementasikan sebuah program atau subprogram yang akan menampilan visualisasi data array sederhana dalam bentuk vertical barcharts, dan sebagai tambahan tampilkan setiap nilai data di sumbu horizontal.")
	fmt.Println("2. Implementasikan algoritma insertion sort, dan gunakan subprogram (1) untuk memvisualisasikan setiap langkah/steps sorting")
	fmt.Println("3. Modifikasi subprogram (2) untuk reverse sorting dan lakukan juga visualisasi dengan subprogram (1)")
	fmt.Print("Pilih Menu : ")
	fmt.Scan(&menu)

	if menu == 1 {

		verticalBar(numbers)

	} else if menu == 2 {

		sortDesc(numbers)

	} else if menu == 3{

		sortAsc(numbers)

	} else {

		fmt.Println("Menu tidak ditemukan")

	}
	
	


}