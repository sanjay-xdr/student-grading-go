package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type Grade string

const (
	A Grade = "A"
	B Grade = "B"
	C Grade = "C"
	F Grade = "F"
)

type student struct {
	firstName, lastName, university                string
	test1Score, test2Score, test3Score, test4Score int
}

type studentStat struct {
	student
	finalScore float32
	grade      Grade
}

func parseCSV(filePath string) []student {
	fd, error := os.Open(filePath)

	if error != nil {
		fmt.Println(error)
	}

	defer fd.Close()

	fileReader := csv.NewReader(fd)
	records, error := fileReader.ReadAll()

	if error != nil {
		fmt.Println(error)
	}

	data := make([]student, 0)

	// fmt.Printf("Records:  %T\n", records)
	for i, record := range records {

		if i == 0 {
			continue
		}

		test1, err1 := strconv.Atoi(record[3])
		test2, err2 := strconv.Atoi(record[4])
		test3, err3 := strconv.Atoi(record[5])
		test4, err4 := strconv.Atoi(record[6])

		if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
			fmt.Print("Something went wrong")
		}
		// console.log(record);
		data = append(data, student{

			record[0],
			record[1],
			record[2],
			test1,
			test2,
			test3,
			test4,
		})

	}
	return data
}

func calculateGrade(students []student) []studentStat {

	result := make([]studentStat, 0)

	for _, data := range students {

		finalScore := float32(data.test1Score+data.test2Score+data.test3Score+data.test4Score) / 4

		var grade Grade

		if finalScore >= 70 {
			grade = "A"
		} else if finalScore >= 50 && finalScore < 70 {
			grade = "B"
		} else if finalScore >= 35 && finalScore < 50 {
			grade = "C"
		} else {
			grade = "F"
		}

		result = append(result, studentStat{
			data, finalScore, grade,
		})

	}
	return result
}

func findOverallTopper(gradedStudents []studentStat) studentStat {

	sort.Slice(gradedStudents, func(i, j int) bool {
		return gradedStudents[i].finalScore > gradedStudents[j].finalScore
	})

	fmt.Print(gradedStudents[0], " Over all topper Ig")
	return gradedStudents[0]
}

func findTopperPerUniversity(gs []studentStat) map[string]studentStat {




	


	
}

func main() {

	fmt.Println(parseCSV("grades.csv"))

	finalData := parseCSV("grades.csv")

	score := calculateGrade(finalData)

	test := findOverallTopper(score).student

	fmt.Print(test)

}
