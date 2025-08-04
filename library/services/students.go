package services

import (
	"sync"
	"sync/atomic"
)

// Student struct represents a student
type Student struct {
	ID         int64
	Roll       int64
	Name       string
	Department string
}

// StudentService manages students in memory
type StudentService struct {
	mu       sync.RWMutex
	students map[int64]*Student
}

var currentStudentID int64 = 0

func getNextStudentID() int64 {
	return atomic.AddInt64(&currentStudentID, 1)
}

// NewStudentService initalizes the service
func NewStudentService() *StudentService {
	return &StudentService{
		students: make(map[int64]*Student),
	}
}

// Addd Student adds a student with a generated ID
func (ss *StudentService) AddStudent(name, dept string, roll int64) *Student {
	ss.mu.Lock()
	defer ss.mu.Unlock()

	student := &Student{
		ID:         getNextStudentID(),
		Name:       name,
		Department: dept,
		Roll:       roll,
	}

	ss.students[student.ID] = student
	return student
}
