#### Problem 1
##### Sorting and visualization

1. *Design* dan implementasikan sebuah *program* atau *subprogram* yang akan menampilan visualisasi *data array* sederhana dalam bentuk *vertical barcharts*, dan sebagai tambahan tampilkan setiap nilai data di sumbu *horizontal*.
    
    ```
    INPUT: Numerical array
    [1, 4, 5, 6, 8, 2]

    OUTPUT: Vertical Barcharts

            |   
            |   
          | |  
        | | |   
      | | | |  
      | | | |  
      | | | | |
    | | | | | | 
    1 4 5 6 8 2 

    ```
2. Implementasikan algoritma *insertion sort*, dan gunakan *subprogram* (1) untuk memvisualisasikan setiap langkah/*steps* *sorting* 

    ```
    INPUT: Numerical array

    [1, 4, 5, 6, 8, 2]

    OUTPUT:

    - Sorted array (ascending)
    - Steps visualization

            |   
            |   
          | |  
        | | |   
      | | | |   
      | | | |   
      | | | | | 
    | | | | | | 
    1 4 5 6 8 2 

              | 
              | 
          |   | 
        | |   | 
      | | |   | 
      | | |   | 
      | | | | | 
    | | | | | | 
    1 4 5 6 2 8 

    ... dan seterusnya ...

    ```

3. Modifikasi *subprogram* (2) untuk *reverse sorting* dan lakukan juga visualisasi dengan *subprogram* (1)






#### Penjelasan

1.  Penjelasan Code

```go
// Function Membuat grapic
func verticalBar(numbers []int){

    // Mencari Maximal pada value slice untuk menentukan tinggi pertama dalam perulangan
	max := maxIntSlice(numbers)

    // Membuat perulangan mencari tinggi
	for i := 0; i < max; i++ {
		fmt.Println(" ")
        // Membuat perulangan mencari lebar
		for j := 0; j < len(numbers); j++ {
            // menentukan Jumlah nilai yang akan di kasi ruang kosong
			space := max - numbers[j]
			if space > i {
				fmt.Print(" ")
			}else{
				fmt.Print("|")
			}
			
		}
	}

	fmt.Println(" ")
    // Menampilkan slice number
	for _, nums := range numbers {
		fmt.Print(nums)
	}
	fmt.Println(" ")
}
```

2.  Pada task #2 menggunakan metode Bubble sort
3.  Pada task #3 menggunakan metode inverse Bubble sort
4.  Untuk running program menggunakan perintah 

``` bash
go run sort.go
```
