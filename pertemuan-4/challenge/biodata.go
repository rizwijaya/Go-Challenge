package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

type student struct {
	no         int
	name       string
	address    string
	occupation string
	reason     string
}

type classroom []student

type biodata interface {
	getName(no int) string
	getAddress(no int) string
	getOccupation(no int) string
	getReason(no int) string
}

func (cr classroom) getName(no int) string {
	for _, student := range cr {
		if no == student.no {
			return student.name
		}
	}
	return "Nama tidak ditemukan"
}

func (cr classroom) getAddress(no int) string {
	for _, student := range cr {
		if no == student.no {
			return student.address
		}
	}
	return "Alamat tidak ditemukan"
}

func (cr classroom) getOccupation(no int) string {
	for _, student := range cr {
		if no == student.no {
			return student.occupation
		}
	}
	return "Pekerjaan tidak ditemukan"
}

func (cr classroom) getReason(no int) string {
	for _, student := range cr {
		if no == student.no {
			return student.reason
		}
	}
	return "Alasan tidak ditemukan"
}

func newClassroom() classroom {
	return classroom{
		student{
			no:         1,
			name:       "Rizqi Wijaya",
			address:    "Jawa Timur",
			occupation: "Backend Developer",
			reason:     "Untuk mencari ilmu baru.",
		},
		student{
			no:         2,
			name:       "Hafidz Ariq",
			address:    "Jawa Barat",
			occupation: "Software Engineer",
			reason:     "Untuk mencari Jodoh",
		},
	}
}

func printData(absen int) {
	cr := newClassroom()
	for _, student := range cr {
		if absen == student.no {
			fmt.Println("No Absen :", student.no)
			fmt.Println("Nama :", student.name)
			fmt.Println("Alamat :", student.address)
			fmt.Println("Pekerjaan :", student.occupation)
			fmt.Println("Alasan :", student.reason)
		}
	}
}

func main() {
	args := os.Args[1:]
	student_id := args[0]
	i, err := strconv.Atoi(student_id)
	if err != nil {
		log.Println(err)
	}
	printData(i)
}
