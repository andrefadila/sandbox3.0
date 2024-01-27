package task

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

type Gender int64

const (
	Akhwat Gender = iota
	Ikhwan
)

func (g Gender) String() (string, error) {
	switch g {
	case Akhwat:
		return "T", nil
	case Ikhwan:
		return "N", nil
	default:
		return "", errors.New("unsupported gender")
	}
}

const maxId = 99999

func CreateNIP(gender Gender, year, month, id int) (string, error) {
	// Gender
	genderStr, errGender := gender.String()
	if errGender != nil {
		return "", errGender
	}

	// Year
	yearStr := strconv.Itoa(year % 100)

	// Month
	if month < 1 || month > 12 {
		return "", errors.New("unsupported month")
	}
	monthStr := "2"
	if month <= 6 {
		monthStr = "1"
	}

	// ID
	if id < 1 || id > maxId {
		return "", errors.New("unsupported id")
	}
	idStr := fmt.Sprintf("%05d", id)

	// Compose NIP
	nip := "AR" + genderStr + yearStr + monthStr + "-" + idStr

	return nip, nil
}

func GenerateNIPs(gender Gender, year, month, count, start int) ([]string, error) {
	nips := []string{}
	for i := start; i < start+count; i++ {
		nip, err := CreateNIP(gender, year, month, i)
		if err != nil {
			return nil, err
		}
		nips = append(nips, nip)
	}
	return nips, nil
}

func isValidNip(nip string) bool {
	regex := regexp.MustCompile(`^(ARN|ART)\d{2}[1-2]-\d{5}$`)
	return regex.MatchString(nip)
}

func CreateNextNIP(nip string) (string, error) {
	// Validation
	if !isValidNip(nip) {
		return "", errors.New("unsupported nip")
	}

	// Get string before id
	nipPrefix := nip[0:7]

	// Get id
	currentId, err := strconv.Atoi(nip[len(nip)-5:])
	nextId := currentId + 1
	if nextId > maxId || err != nil {
		return "", errors.New("unsupported id")
	}

	// Compose next NIP
	nextNip := nipPrefix + fmt.Sprintf("%05d", nextId)

	return nextNip, nil
}

func GenerateNextNIPs(nip string, count int) ([]string, error) {
	nips := []string{}
	for i := 0; i < count; i++ {
		nextNip, err := CreateNextNIP(nip)
		if err != nil {
			return nil, err
		}
		nips = append(nips, nextNip)
		nip = nextNip
	}
	return nips, nil
}
