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

	verticalBar(numbers)
	sortDesc(numbers)
	sortAsc(numbers)


}