// package main

// import (
// 	"strconv"

// 	"github.com/gin-gonic/gin"
// )

// type User struct {
// 	ID    string `json:"id"`
// 	Name  string `json:"name"`
// 	Email string `json:"email"`
// }

// var userList []User

// func main() {
// 	// Inisialisasi router
// 	router := gin.Default()

// 	// Definisikan route dan handler untuk GET /users
// 	router.GET("/users", func(c *gin.Context) {
// 		c.JSON(200, userList)
// 	})

// 	// Definisikan route dan handler untuk POST /users
// 	router.POST("/users", func(c *gin.Context) {
// 		var newUser User
// 		if err := c.ShouldBindJSON(&newUser); err != nil {
// 			c.JSON(400, gin.H{"error": "Data tidak valid"})
// 			return
// 		}
// 		newUser.ID = strconv.Itoa(len(userList) + 1)
// 		userList = append(userList, newUser)
// 		c.JSON(201, newUser)
// 	})

// 	// Definisikan route dan handler untuk PUT /users/:id
// 	router.PUT("/users/:id", func(c *gin.Context) {
// 		id := c.Param("id")
// 		var updatedUser User
// 		if err := c.ShouldBindJSON(&updatedUser); err != nil {
// 			c.JSON(400, gin.H{"error": "Data tidak valid"})
// 			return
// 		}
// 		for i, user := range userList {
// 			if user.ID == id {
// 				userList[i].Name = updatedUser.Name
// 				userList[i].Email = updatedUser.Email
// 				c.JSON(200, userList[i])
// 				return
// 			}
// 		}
// 		c.JSON(404, gin.H{"error": "Pengguna tidak ditemukan"})
// 	})

// 	// Definisikan route dan handler untuk DELETE /users/:id
// 	router.DELETE("/users/:id", func(c *gin.Context) {
// 		id := c.Param("id")
// 		for i, user := range userList {
// 			if user.ID == id {
// 				userList = append(userList[:i], userList[i+1:]...)
// 				c.JSON(200, gin.H{"message": "Pengguna berhasil dihapus"})
// 				return
// 			}
// 		}
// 		c.JSON(404, gin.H{"error": "Pengguna tidak ditemukan"})
// 	})

// 	// Jalankan server
// 	router.Run(":8080")
// }

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// URL endpoint API publik
	url := "https://apiweb.mile.app/api/v3/task"

	// Bearer token
	token := "Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJhdWQiOiI2MmNmZWY2NDQ4NGMxZTcwMTU0YjQ5YjIiLCJqdGkiOiJjMjc4MDQyMWFkNjk5YzQ0YmJkMmU2ZjUzZTk4NDM5ZmRhOTAzNjhmNTA3N2Y0MzNiZGQ0MDVkZmE4MTI3YTZiOGI0ZWIwYzI5ZTI2OWEzMCIsImlhdCI6MTY5Nzc2OTI2MywibmJmIjoxNjk3NzY5MjYzLCJleHAiOjQ4NTM0NDI4NjMsInN1YiI6IjY1MjZiNTE2Y2U3YWY4ZWMxZTA5NmY3YiIsInNjb3BlcyI6W119.oWi2UTbyfcTkE2iqE62wYy0qacuNMO1cwecDFvzXA9cUqwzVeTKDTgo8Zq1Arr9Txr4G4z9RylRsx5mTYroaDBIWuoWfcmnljF8VMe9HhWHP8EqC4w3Wv6Eoke8WkkUPqVgGoH3zvgV0DNfR12mCjpoRn8Tbn-5613VQGun8n7cs0UTl_snq5Q1uYzUFsG4ODNDPQ1E3NWO6dcvZ4y1V8ONNqayDRoQdddPl0yFSCCFToMYQo_ymd_kjIDqTkcqHzur_V04AjKpa66iWC4Lw9uhQ3Gdfa26DV9sqqnx8J2cMYsR9XlrUWzioSsbE82bsMAgX8oHCWOq8J2D1G5cBHq_hm7kZ5eE8daySoaIr6ldxSIsGep1BDlAEE2cV3ogMIz3RjM0X-DFxIfWaos1ZQSnI1wrA12QSan75sEpwflIpXposRfnJe6H4rZZrycHGUPv3-fpPF7JEgW8kJQ_bcyVAPAKrTr1G3bHWKSzuZYavXgiH2MLv1rwJ3W-wJpEv50BXEftE0dJjxynAVGoqJQr_osIUeumsE0ZKUm29B-0eIYMpTGgI6i9fanmBJxNXxcpmClI36Ac34ZueV53B16ZAJ3auwzpxOjpEEzolxKTuZjAEGEkc2iff2tNcLTaXMg2AulGDX6pHJhlG0qUo4hJ3PX3AGZYNNVCnf_KbgR8"

	// Buat HTTP client
	client := &http.Client{}

	// Buat HTTP request dengan method GET
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Gagal membuat request:", err)
		return
	}

	// Set header Authorization dengan bearer token
	req.Header.Set("Authorization", token)

	// Kirim request ke API publik
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Gagal mengirim request:", err)
		return
	}
	defer resp.Body.Close()

	// Baca respons dari API
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Gagal membaca respons:", err)
		return
	}

	// Tampilkan respons dari API
	fmt.Println(string(body))
}

// package main

// import (
// 	"database/sql"
// 	"fmt"
// 	"log"

// 	_ "github.com/go-sql-driver/mysql"
// )

// func main() {
// 	// Konfigurasi koneksi database
// 	db, err := sql.Open("mysql", "root:hacker149@tcp(localhost:3306)/pasien")
// 	if err != nil {
// 		log.Fatal("Gagal membuka koneksi database:", err)
// 	}
// 	defer db.Close()

// 	// Uji koneksi ke database
// 	err = db.Ping()
// 	if err != nil {
// 		log.Fatal("Gagal terhubung ke database:", err)
// 	}

// 	fmt.Println("Koneksi ke database berhasil!")

// 	// Lakukan operasi lain dengan database di sini
// 	// ...
// }

// package main

// import (
// 	"database/sql"
// 	"fmt"
// 	"log"
// 	"net/http"

// 	_ "github.com/go-sql-driver/mysql"
// )

// func main() {
// 	// Membuat handler untuk endpoint /data
// 	http.HandleFunc("/data", func(w http.ResponseWriter, r *http.Request) {
// 		// Konfigurasi koneksi database
// 		db, err := sql.Open("mysql", "root:hacker149@tcp(localhost:3306)/pasien")
// 		if err != nil {
// 			http.Error(w, "Gagal membuka koneksi database", http.StatusInternalServerError)
// 			return
// 		}
// 		defer db.Close()

// 		// Query data dari tabel pasien
// 		rows, err := db.Query("SELECT id, nama, alamat FROM pasien")
// 		if err != nil {
// 			http.Error(w, "Gagal melakukan query", http.StatusInternalServerError)
// 			return
// 		}
// 		defer rows.Close()

// 		// Looping hasil query dan tampilkan data
// 		for rows.Next() {
// 			var id int
// 			var nama, alamat string
// 			err = rows.Scan(&id, &nama, &alamat)
// 			if err != nil {
// 				http.Error(w, "Gagal membaca data", http.StatusInternalServerError)
// 				return
// 			}
// 			fmt.Fprintf(w, "ID: %d, Nama: %s, Alamat: %s\n", id, nama, alamat)
// 		}
// 		err = rows.Err()
// 		if err != nil {
// 			http.Error(w, "Gagal membaca hasil query", http.StatusInternalServerError)
// 			return
// 		}
// 	})

// 	log.Println("Server berjalan pada http://localhost:8080/")
// 	log.Fatal(http.ListenAndServe(":8080", nil))
// }

// test
