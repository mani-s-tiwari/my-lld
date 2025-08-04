package services

import (
	"sync"
	"sync/atomic"
	"time"
)

type Issue struct {
	ID        int64
	BookID    int64
	StudentID int64
	IssueDate time.Time
	DueDate   time.Time
	Returned  bool
}

// IssueService manages issued books
type IssueService struct {
	mu     sync.RWMutex
	issues map[int64]*Issue
}

var currentIssueID int64 = 0

func getNextIssueID() int64 {
	return atomic.AddInt64(&currentIssueID, 1)
}

// NewIssueService intializes the issue service
func NewIssueService() *IssueService {
	return &IssueService{
		issues: make(map[int64]*Issue),
	}
}

// IssueBook creates a issue record
func (is *IssueService) IssueBook(bookID, studentID int64, durationDays int) *Issue {
	is.mu.Lock()
	defer is.mu.Unlock()

	now := time.Now().UTC()
	issue := &Issue{
		ID:        getNextIssueID(),
		BookID:    bookID,
		StudentID: studentID,
		IssueDate: now,
		DueDate:   now.AddDate(0, 0, durationDays),
		Returned:  false,
	}
	is.issues[issue.ID] = issue
	return issue
}

// ReturnBook marks an issue as returned
func (is *IssueService) ReturnBook(issueID int64) bool {
	is.mu.Lock()
	defer is.mu.Unlock()

	issue, exists := is.issues[issueID]
	if !exists || issue.Returned {
		return false
	}

	issue.Returned = true
	return true
}
