package contest

import "time"

// ResultType represents enum of the submission result
type ResultType uint8

const (
	// Success indicates the result of submission is good to go!
	Success ResultType = 0
	// CompileTimeError indicates the result of compilation error
	CompileTimeError ResultType = 1
	// RuntimeError indicates the runtime error while running submission
	RuntimeError ResultType = 2
	// WrongAnswer indicates the program runs but not getting correct answer
	WrongAnswer ResultType = 3
)

// User represent contest participants, judge, admins
type User struct {
	ID        string `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Type      string `json:"type"`
}

// Team contains a list of parcipants together
type Team struct {
	ID      string `json:"id"`
	Name    string `json:"id"`
	Members []User `json:"members"`
	Score   uint32 `json:"score"`
}

// Attachment represent a single file upload entry
type Attachment struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Mimetype string `json:"mimetype"`
	Path     string `json:"-"`
}

// Question represent a single question
type Question struct {
	ID          string       `json:"id"`
	Description string       `json:"description"`
	Attachments []Attachment `json:"attachments"`
}

// Submission represent a source code submission from the partcipants
type Submission struct {
	ID         string       `json:"id"`
	SourceCode []Attachment `json:"sourcecode"`
	Result     ResultType   `json:"result"`
	Feedback   string       `json:"feedback"`
	Question   Question     `json:"question"`
	Team       Team         `json:"team"`
}

// Contest wraps all contest related stuff together
type Contest struct {
	ID          uint32     `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	HostedDate  time.Time  `json:"hostedDate"`
	Teams       []Team     `json:"team"`
	Questions   []Question `json:"questions"`
}

type Client interface {
	ContestService() ContestService
}

type ContestService interface {
}
