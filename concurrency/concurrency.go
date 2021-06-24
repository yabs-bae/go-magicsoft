package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
)
 
var (
	// wg is used to force the application wait for all goroutines to finish before exiting.
	wg sync.WaitGroup
	// jobChan is a buffered channel that has the capacity of maximum 11 resource slot.
	jobChan = make(chan int, 11)
	//
	url string = "https://data.gov.sg/api/action/datastore_search?resource_id=eb8b932c-503c-41e7-b513-114cffbe2338&limit=500"

	path = "./export/";

)
 

// Struct untuk export Data JSON
type Json struct {
	Help string `json:"help"`
	Success bool `json:"success"`
	Result struct {
		Resource_id string `json:"resource_id"`
		Records []struct {
			ID int16 `json:"_id"`
			Sex string `json:"sex"`
			NoOfGraduates string `json:"no_of_graduates"`
			TypeOfCourse string `json:"type_of_course"`
			Year string `json:"year"`

		}
	}
}




// Function untuk export CSV
func exportCSV(year string , part [][]string,path string){
	os.MkdirAll(path, os.ModePerm)
	csvFile, err := os.Create(path+year+".csv")

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	csvwriter := csv.NewWriter(csvFile)
	
	for _, empRow := range part {
		_ = csvwriter.Write(empRow)
	}
	csvwriter.Flush()
	csvFile.Close()
}

// FUnction untuk mengambil hasil data JSON dari URL
func MakeRequest() (map[string][][]string){
	resp, _ := http.Get(url)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
        log.Println("err handle it")
    }
	
	var data Json
    json.Unmarshal(body, &data)
	dataJson := data.Result.Records


	empData := make(map[string][][]string)

	for _, res := range dataJson {

		temp := []string{
			strconv.Itoa(int(res.ID)),
			res.Sex,
			res.NoOfGraduates,
			res.TypeOfCourse,
			res.Year,
		}
		empData[res.Year] = append(empData[res.Year], temp)
	}

	return empData


  }

func main() {
	rand.Seed(time.Now().UnixNano())
	length := len(MakeRequest())
	fmt.Println("BEGIN")
 
	// Tell how many worker you will be running.
	wg.Add(1)
 
	// Run 1 worker to handle jobs.
	go worker(jobChan, &wg)
 
	// Send 10 jobs to job channel.
	for i := 1; i <= length; i++ {
		if !queueJob(i, jobChan) {
			fmt.Println("Channel is full... Service unavailable...")
		}
	}
 
	// Close the job channel.
	close(jobChan)
 
	// Block exiting until all the goroutines are finished.
	wg.Wait()
 
	fmt.Println("END")
}
 
// queueJob puts job into channel. If channel buffer is full, return false.
func queueJob(job int, jobChan chan<- int) bool {
	select {
	case jobChan <- job:
		return true
	default:
		return false
	}
}
 
// worker processes jobs.
func worker(jobChan <-chan int, wg *sync.WaitGroup) {
	// As soon as the current goroutine finishes (job done!), notify back WaitGroup.
	defer wg.Done()
 
	fmt.Println("Worker is waiting for jobs")
	result := MakeRequest()

	var procces []string 

	for year , _ := range result {
		procces = append(procces, year)
	}

	


	for job := range jobChan {
			temp := job - 1
			log.Println("Worker picked job", job)
			log.Println(procces[temp]) 

			wait := time.Duration(rand.Intn(len(result)))
			time.Sleep(wait * time.Second)
		if len(procces) > job {
			data := result[procces[temp]]
			exportCSV(procces[temp],data,path)
			// Process the job ...
	
			log.Println("Worker completed job", job, "in", int(wait), "second(s)")
		}else{
			log.Println("Worker out of length ", job, "in", int(wait), "second(s)")

		}
		
	}
}