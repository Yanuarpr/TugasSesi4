package main

import (
	"fmt"
	"sort"
)

type Student struct {
	Name  string
	Score []int
}

func (s Student) AverageScore() float64 {
	sum := 0
	for _, score := range s.Score {
		sum += score
	}
	return float64(sum) / float64(len(s.Score))
}

func (s Student) MinMaxScore() (int, int) {
	sort.Ints(s.Score)
	min := s.Score[0]
	max := s.Score[len(s.Score)-1]
	return min, max
}

func main() {
	// Membuat slice untuk menyimpan data siswa
	students := make([]Student, 0)

	// Memasukkan data 5 siswa
	students = append(students, Student{Name: "Rizky", Score: []int{80}})
	students = append(students, Student{Name: "Kobar", Score: []int{75}})
	students = append(students, Student{Name: "Ismail", Score: []int{70}})
	students = append(students, Student{Name: "Umam", Score: []int{75}})
	students = append(students, Student{Name: "Yopan", Score: []int{60}})

	// Menghitung skor rata-rata dan mencari skor minimum dan maksimum dari seluruh siswa
	totalScore := 0
	minScore := 100
	maxScore := 0

	for _, student := range students {
		averageScore := student.AverageScore()
		totalScore += int(averageScore)

		min, max := student.MinMaxScore()
		if min < minScore {
			minScore = min
		}
		if max > maxScore {
			maxScore = max
		}
	}

	// Menampilkan hasil
	fmt.Printf("Skor rata-rata dari seluruh siswa adalah: %.2f\n", float64(totalScore)/float64(len(students)))
	fmt.Printf("Skor minimum dari seluruh siswa adalah: %d\n", minScore)
	fmt.Printf("Skor maksimum dari seluruh siswa adalah: %d\n", maxScore)
}
