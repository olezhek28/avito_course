package hw10programoptimization

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/buger/jsonparser"
)

type User struct {
	ID       int64
	Name     string
	Username string
	Email    string
	Phone    string
	Password string
	Address  string
}

type DomainStat map[string]int

func GetDomainStat(r io.Reader, domain string) (DomainStat, error) {
	return countDomains(r, domain)
}

func countDomains(r io.Reader, domain string) (DomainStat, error) {
	scanner := bufio.NewScanner(r)
	res := make(DomainStat)
	for scanner.Scan() {
		user, err := getUserInfo(scanner.Bytes())
		if err != nil {
			return nil, err
		}

		if !strings.Contains(user.Email, "@") {
			return nil, fmt.Errorf("email does not contain @")
		}

		if strings.Contains(user.Email, domain) {
			tail := strings.ToLower(strings.SplitN(user.Email, "@", 2)[1])
			res[tail]++
		}
	}

	return res, nil
}

func getUserInfo(line []byte) (*User, error) {
	user := new(User)
	var err error
	user.ID, err = jsonparser.GetInt([]byte(line), "Id")
	if err != nil {
		return nil, err
	}

	user.Username, err = jsonparser.GetString(line, "Username")
	if err != nil {
		return nil, err
	}

	user.Email, err = jsonparser.GetString(line, "Email")
	if err != nil {
		return nil, err
	}

	user.Phone, err = jsonparser.GetString(line, "Phone")
	if err != nil {
		return nil, err
	}

	user.Password, err = jsonparser.GetString(line, "Password")
	if err != nil {
		return nil, err
	}

	user.Address, err = jsonparser.GetString(line, "Address")
	if err != nil {
		return nil, err
	}

	return user, nil
}
