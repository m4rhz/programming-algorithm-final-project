package main

import "fmt"

const (
	maxStudents int = 2023
	maxCourses  int = 2023
)

type Course struct {
	ID    int
	Name  string
	SKS   int
	UTS   float64
	UAS   float64
	Quiz  float64
	Total float64
	Grade string
}

type Student struct {
	ID      int
	Name    string
	Courses [maxCourses]Course
	total   int
}

func main() {
	var students [maxStudents]Student
	var courses [maxCourses]Course
	var studentCount int
	var courseCount int

	menu(&students, &courses, &studentCount, &courseCount)

}

func menu(students *[maxStudents]Student, courses *[maxCourses]Course, studentCount *int, courseCount *int) {
	running := true
	for running {
		fmt.Println("Aplikasi Nilai Mahasiswa")
		fmt.Println("1. Tambah Mahasiswa")
		fmt.Println("2. Edit Mahasiswa")
		fmt.Println("3. Hapus Mahasiswa")
		fmt.Println("4. Tambah Matakuliah")
		fmt.Println("5. Edit Matakuliah")
		fmt.Println("6. Hapus Matakuliah")
		fmt.Println("7. Tampilkan Daftar Matakuliah Berdasarkan Mahasiswa")
		fmt.Println("8. Tampilkan Daftar Mahasiswa Berdasarkan Matakuliah")
		fmt.Println("9. Tampilkan Mahasiswa Terurut Nilai")
		fmt.Println("10. Tampilkan Mahasiswa Terurut Matakuliah")
		fmt.Println("11. Tampilkan Transkrip Nilai Mahasiswa")
		fmt.Println("12. Enroll Mahasiswa ke Matakuliah Tertentu")
		fmt.Println("13. Masukkan Nilai Untuk Setiap Matakuliah")
		fmt.Println("0. Keluar")
		var choice int
		fmt.Print("Pilihan: ")
		fmt.Scan(&choice)
		if choice == 1 {
			addStudent(students, studentCount)
		} else if choice == 2 {
			editStudent(students, studentCount)
		} else if choice == 3 {
			deleteStudent(students, studentCount)
		} else if choice == 4 {
			addCourse(courses, courseCount)
		} else if choice == 5 {
			editCourse(courses, courseCount)
		} else if choice == 6 {
			deleteCourse(courses, courseCount)
		} else if choice == 7 {
			searchCoursesByStudent(students, courses, studentCount, courseCount)
		} else if choice == 8 {
			searchStudentsByCourse(students, courses, studentCount, courseCount)
		} else if choice == 9 {
			printStudentsByGrade(students, studentCount)
		} else if choice == 10 {
			printStudentsBySKS(students, studentCount, courses)
		} else if choice == 11 {
			for i := 0; i < *studentCount; i++ {
				printTranscript(&students[i])
			}
		} else if choice == 12 {
			enrollCourse(students, courses, studentCount, courseCount)
		} else if choice == 13 {
			inputCourseGrade(students, studentCount, courses, courseCount)
		} else if choice == 0 {
			fmt.Println("Terima kasih. Sampai jumpa!")
			running = false
		} else {
			fmt.Println("Pilihan tidak valid. Silakan pilih kembali.")
		}
	}
}

func addStudent(students *[maxStudents]Student, studentCount *int) {
	if *studentCount == maxStudents {
		fmt.Println("Kapasitas mahasiswa telah penuh.")
	} else {
		var student Student
		fmt.Print("ID Mahasiswa: ")
		fmt.Scan(&student.ID)
		fmt.Print("Nama Mahasiswa: ")
		fmt.Scan(&student.Name)
		exists := false
		for i := 0; i < *studentCount; i++ {
			if students[i].ID == student.ID {
				exists = true
			}
		}
		if exists {
			fmt.Println("Mahasiswa dengan ID tersebut sudah terdaftar.")
		} else {
			students[*studentCount] = student
			*studentCount++
			fmt.Println("Mahasiswa berhasil ditambahkan.")
		}
	}
}

func editStudent(students *[maxStudents]Student, studentCount *int) {
	if *studentCount == 0 {
		fmt.Println("Tidak ada mahasiswa yang tersedia.")
	} else {
		var id int
		fmt.Print("ID Mahasiswa: ")
		fmt.Scan(&id)
		found := false
		for i := 0; i < *studentCount; i++ {
			if students[i].ID == id {
				var name string
				fmt.Print("Nama Mahasiswa: ")
				fmt.Scan(&name)
				students[i].Name = name
				fmt.Println("Mahasiswa berhasil diubah.")
				found = true
			}
		}
		if !found {
			fmt.Println("Mahasiswa dengan ID tersebut tidak ditemukan.")
		}
	}
}

func deleteStudent(students *[maxStudents]Student, studentCount *int) {
	if *studentCount == 0 {
		fmt.Println("Tidak ada mahasiswa yang tersedia.")
	} else {
		var id int
		fmt.Print("ID Mahasiswa: ")
		fmt.Scan(&id)
		found := false
		for i := 0; i < *studentCount; i++ {
			if students[i].ID == id {
				found = true
				for j := i; j < *studentCount-1; j++ {
					students[j] = students[j+1]
				}
				students[*studentCount-1] = Student{}
				*studentCount--

				fmt.Println("Mahasiswa berhasil dihapus.")
			}
		}
		if !found {
			fmt.Println("Mahasiswa dengan ID tersebut tidak ditemukan.")
		}
	}
}

func addCourse(courses *[maxCourses]Course, courseCount *int) {
	if *courseCount == maxCourses {
		fmt.Println("Kapasitas matakuliah telah penuh.")
	} else {
		var course Course
		fmt.Print("ID Matakuliah: ")
		fmt.Scan(&course.ID)
		fmt.Print("Nama Matakuliah: ")
		fmt.Scan(&course.Name)
		fmt.Print("SKS Matakuliah: ")
		fmt.Scan(&course.SKS)
		duplicateID := false
		for i := 0; i < *courseCount; i++ {
			if courses[i].ID == course.ID {
				duplicateID = true
			}
		}
		if duplicateID {
			fmt.Println("Matakuliah dengan ID tersebut sudah terdaftar.")
		} else {
			courses[*courseCount] = course
			*courseCount++
			fmt.Println("Matakuliah berhasil ditambahkan.")
		}
	}
}

func editCourse(courses *[maxCourses]Course, courseCount *int) {
	if *courseCount == 0 {
		fmt.Println("Tidak ada matakuliah yang tersedia.")
	} else {
		var id int
		fmt.Print("ID Matakuliah: ")
		fmt.Scan(&id)
		found := false
		for i := 0; i < *courseCount; i++ {
			if courses[i].ID == id {
				var name string
				var sks int
				fmt.Print("Nama Matakuliah: ")
				fmt.Scan(&name)
				fmt.Print("SKS Matakuliah: ")
				fmt.Scan(&sks)
				courses[i].Name = name
				courses[i].SKS = sks
				fmt.Println("Matakuliah berhasil diubah.")
				found = true
			}
		}
		if !found {
			fmt.Println("Matakuliah dengan ID tersebut tidak ditemukan.")
		}
	}
}

func deleteCourse(courses *[maxCourses]Course, courseCount *int) {
	if *courseCount == 0 {
		fmt.Println("Tidak ada matakuliah yang tersedia.")
	} else {
		var id int
		fmt.Print("ID Matakuliah: ")
		fmt.Scan(&id)
		found := false
		for i := 0; i < *courseCount; i++ {
			if courses[i].ID == id {
				found = true
				for j := i; j < *courseCount-1; j++ {
					courses[j] = courses[j+1]
				}
				courses[*courseCount-1] = Course{}
				*courseCount--
			}
		}
		if !found {
			fmt.Println("Matakuliah dengan ID tersebut tidak ditemukan.")
		} else {
			fmt.Println("Matakuliah berhasil dihapus.")
		}
	}
}

func enrollCourse(students *[maxStudents]Student, courses *[maxCourses]Course, studentCount *int, courseCount *int) {
	success := false

	if *studentCount == 0 {
		fmt.Println("Tidak ada mahasiswa yang tersedia.")
	} else if *courseCount == 0 {
		fmt.Println("Tidak ada matakuliah yang tersedia.")
	} else {
		var studentID int
		fmt.Print("ID Mahasiswa: ")
		fmt.Scan(&studentID)

		var studentIndex int
		found := false
		for i := 0; i < *studentCount; i++ {
			if students[i].ID == studentID {
				studentIndex = i
				found = true
			}
		}

		if !found {
			fmt.Println("Mahasiswa dengan ID tersebut tidak ditemukan.")
		} else {
			var courseID int
			fmt.Print("ID Matakuliah: ")
			fmt.Scan(&courseID)

			var courseIndex int
			found = false
			for i := 0; i < *courseCount; i++ {
				if courses[i].ID == courseID {
					courseIndex = i
					found = true
				}
			}

			if !found {
				fmt.Println("Matakuliah dengan ID tersebut tidak ditemukan.")
			} else {
				var check bool = false
				student := &students[studentIndex]
				for i := 0; i < maxCourses && check != true; i++ {
					if student.Courses[i].ID == 0 {
						student.Courses[i] = courses[courseIndex]
						student.total += courses[courseIndex].SKS
						fmt.Println("Mahasiswa berhasil mendaftar matakuliah.")
						success = true
						check = true
					}
				}

				if !success {
					fmt.Println("Mahasiswa telah mencapai batas maksimum matakuliah yang dapat diambil.")
				}
			}
		}
	}

	if !success {
		fmt.Println("Gagal mendaftarkan mahasiswa ke matakuliah.")
	}
}

func searchCoursesByStudent(students *[maxStudents]Student, courses *[maxCourses]Course, studentCount *int, courseCount *int) {
	if *studentCount == 0 {
		fmt.Println("Tidak ada mahasiswa yang tersedia.")
	} else {
		var studentID int
		fmt.Print("ID Mahasiswa: ")
		fmt.Scan(&studentID)
		var studentIndex int
		found := false
		for i := 0; i < *studentCount; i++ {
			if students[i].ID == studentID {
				studentIndex = i
				found = true
			}
		}
		if !found {
			fmt.Println("Mahasiswa dengan ID tersebut tidak ditemukan.")
		} else {
			student := students[studentIndex]
			fmt.Printf("Matakuliah yang diambil oleh Mahasiswa %d - %s:\n", student.ID, student.Name)
			for i := 0; i < maxCourses; i++ {
				course := student.Courses[i]
				if course.ID != 0 {
					fmt.Printf("- %d. %s\n", course.ID, course.Name)
				}
			}
		}
	}
}

func searchStudentsByCourse(students *[maxStudents]Student, courses *[maxCourses]Course, studentCount *int, courseCount *int) {
	if *courseCount == 0 {
		fmt.Println("Tidak ada matakuliah yang tersedia.")
	} else {
		var courseID int
		fmt.Print("ID Matakuliah: ")
		fmt.Scan(&courseID)
		var courseIndex int
		found := false
		for i := 0; i < *courseCount; i++ {
			if courses[i].ID == courseID {
				courseIndex = i
				found = true
			}
		}
		if !found {
			fmt.Println("Matakuliah dengan ID tersebut tidak ditemukan.")
		} else {
			course := courses[courseIndex]
			fmt.Printf("Mahasiswa yang mengambil Matakuliah %d - %s:\n", course.ID, course.Name)
			for i := 0; i < *studentCount; i++ {
				student := students[i]
				for j := 0; j < maxCourses; j++ {
					if student.Courses[j].ID == course.ID {
						fmt.Printf("- %d. %s\n", student.ID, student.Name)
					}
				}
			}
		}
	}
}

func printStudentsBySKS(students *[maxStudents]Student, studentCount *int, courses *[maxCourses]Course) {
	var sortedStudents [maxStudents]Student
	sortedCount := 0
	for i := 0; i < *studentCount; i++ {
		sortedStudents[sortedCount] = students[i]
		sortedCount++
	}
	total := 0
	for i := 0; i < maxCourses; i++ {
		total += courses[i].SKS
	}
	// Insertion Sort (Descending)
	for i := 1; i < sortedCount; i++ {
		key := sortedStudents[i]
		j := i - 1
		for j >= 0 && sortedStudents[j].total < key.total {
			sortedStudents[j+1] = sortedStudents[j]
			j--
		}
		sortedStudents[j+1] = key
	}
	for i := 0; i < sortedCount; i++ {
		student := sortedStudents[i]
		fmt.Printf("- ID: %d, Nama: %s, Total SKS: %d\n", student.ID, student.Name, student.total)
	}
}

func printStudentsByGrade(students *[maxStudents]Student, studentCount *int) {
	// Selection Sort (Ascending)
	fmt.Println("Daftar Mahasiswa Terurut Nilai:")
	for i := 0; i < *studentCount-1; i++ {
		minIndex := i
		for j := i + 1; j < *studentCount; j++ {
			if students[j].ID != 0 && students[j].Courses[0].ID != 0 {
				total1 := (students[minIndex].Courses[0].UTS + students[minIndex].Courses[0].UAS + students[minIndex].Courses[0].Quiz) / 3
				total2 := (students[j].Courses[0].UTS + students[j].Courses[0].UAS + students[j].Courses[0].Quiz) / 3
				if total2 < total1 {
					minIndex = j
				}
			}
		}
		if minIndex != i {
			students[i], students[minIndex] = students[minIndex], students[i]
		}
	}
	for i := 0; i < *studentCount; i++ {
		student := &students[i]
		if student.ID != 0 {
			for j := 0; j < maxCourses; j++ {
				course := &student.Courses[j]
				if course.UTS != 0 && course.UAS != 0 && course.Quiz != 0 {
					total := (course.UTS + course.UAS + course.Quiz) / 3
					grade := calculateGrade(total)
					fmt.Printf("- ID: %d, Nama: %s, Total: %.2f, Grade: %s\n", student.ID, student.Name, total, grade)
				}
			}
		}
	}
}

func inputCourseGrade(students *[maxStudents]Student, studentCount *int, courses *[maxCourses]Course, courseCount *int) {
	var studentID int
	fmt.Print("ID Mahasiswa: ")
	fmt.Scan(&studentID)

	studentIndex := findStudentIndexByID(students, studentCount, studentID)
	if studentIndex == -1 {
		fmt.Println("Mahasiswa tidak ditemukan.")
	} else {
		var courseID int
		fmt.Print("ID Matakuliah: ")
		fmt.Scan(&courseID)
		courseIndex := findCourseIndexByID(courses, courseCount, courseID)
		if courseIndex == -1 {
			fmt.Println("Matakuliah tidak ditemukan.")
		} else {
			course := &students[studentIndex].Courses[courseIndex]

			fmt.Print("Nilai UTS: ")
			fmt.Scan(&course.UTS)

			fmt.Print("Nilai UAS: ")
			fmt.Scan(&course.UAS)

			fmt.Print("Nilai Quiz: ")
			fmt.Scan(&course.Quiz)

			course.Total = (course.UTS + course.UAS + course.Quiz) / 3
			course.Grade = calculateGrade(course.Total)

			fmt.Println("Nilai berhasil diinput.")
		}
	}
}

func findStudentIndexByID(students *[maxStudents]Student, studentCount *int, studentID int) int {
	for i := 0; i < *studentCount; i++ {
		if students[i].ID == studentID {
			return i
		}
	}
	return -1
}

func findCourseIndexByID(courses *[maxCourses]Course, courseCount *int, courseID int) int {
	for i := 0; i < *courseCount; i++ {
		if courses[i].ID == courseID {
			return i
		}
	}
	return -1
}

func printTranscript(student *Student) {
	fmt.Println("Transkrip Mahasiswa")
	fmt.Println("ID:", student.ID)
	fmt.Println("Nama:", student.Name)

	for i := 0; i < maxCourses; i++ {
		course := &student.Courses[i]
		if course.UTS != 0 && course.UAS != 0 && course.Quiz != 0 {
			fmt.Println(course.Name)
			total := (course.UTS + course.UAS + course.Quiz) / 3
			grade := calculateGrade(total)
			fmt.Printf("Nilai UTS: %.2f\n", course.UTS)
			fmt.Printf("Nilai UAS: %.2f\n", course.UAS)
			fmt.Printf("Nilai Quiz: %.2f\n", course.Quiz)
			fmt.Printf("Nilai Total: %.2f\n", total)
			fmt.Printf("Grade: %s\n", grade)
		}
	}
	fmt.Println()
}

func calculateGrade(total float64) string {
	if total >= 80 {
		return "A"
	} else if total >= 70 {
		return "B"
	} else if total >= 60 {
		return "C"
	} else if total >= 50 {
		return "D"
	} else {
		return "E"
	}
}
