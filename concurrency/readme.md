#### Problem 4
##### Concurrency Task Worker

Implementasi sebuah program untuk merangkum data graduates di singapore berdasarkan tahun lulus dan jurusan, simpan dalam file csv per tahun.

```
2020.csv
2019.csv
``` 

* Untuk sumber informasi gunakan API yang disediakan oleh **Singapore Goverment Data** khusus data [graduates](https://data.gov.sg/dataset/graduates-from-university-first-degree-courses-by-type-of-course?view_id=fa0e401c-6251-4a15-aebc-a5f3d2c85752&resource_id=eb8b932c-503c-41e7-b513-114cffbe2338)
* Gunakan *net/http* package untuk mengambil data dari API yang disediakan.
* Olah data mentah dari API jika diperlukan.
* Implementasi *concurrent* process untuk process mengambilan data dari API menggunakan *goroutine*, untuk [ilustrasi](https://talks.golang.org/2012/concurrency.slide)
* Batasi jumlah *concurrent* process, dengan mengaplikasikan *Queue* dan *Worker*.

ilustrasi:
```
grab -concurrent_limit=2 -output=/home/yourname/museum 
```

#### Untuk menjalankan program

```bash
go run concurrency.go
```
