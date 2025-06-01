package utils

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type ErrType map[string]interface{}

func GetSlug(s string) string {
	s = strings.ToLower(s)
	s = strings.ReplaceAll(s, " ", "-")
	s = strings.ReplaceAll(s, ":", "")
	s = strings.ReplaceAll(s, ";", "")
	s = strings.ReplaceAll(s, ",", "")
	s = strings.ReplaceAll(s, ".", "")

	var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

	// Generate a random integer between 0 and 9
	randomInt := rnd.Intn(10)
	s = s + "-" + strconv.Itoa(randomInt)
	return s
}

// Hashing Password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

// Comparing Password
func IsValidPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

type Platform string

const (
	Web     Platform = "web"
	Android Platform = "android"
	IOS     Platform = "ios"
	All     Platform = "all"
)

type Timer struct {
	Logs map[string]time.Time
}

func (t *Timer) Start(name string) {
	if t.Logs == nil {
		t.Logs = make(map[string]time.Time)
	}
	t.Logs[name] = time.Now()
}

func (t *Timer) End(name string) {
	start, is := t.Logs[name]
	if !is {
		fmt.Println("Start time not declared")
		return
	}
	delete(t.Logs, name)
	duration := time.Since(start)
	fmt.Println(name, "took", duration)
}

func ParseDate(date string) (time.Time, error) {
	layouts := []string{
		"2006-01-02",
		"2006/01/02",
		"2006-01-02 15:04:05",
		"2006-01-02T15:04:05",
		"2006-01-02T15:04:05Z",
		"2006-01-02T15:04:05.000Z",
	}

	for _, layout := range layouts {
		if d, err := time.Parse(layout, date); err == nil {
			return d, nil
		}
	}

	return time.Time{}, fmt.Errorf("invalid date format: %s", date)
}
