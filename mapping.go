package main

import (
	"fmt"
)

type queryResponse struct {
	id          int
	originalUrl string
	singleUse   int
	expired     int
}

func (s *StoreStruct) CreateMapping(url, code string, singleUse bool) (bool, error) {
	_, err := s.stmts[QueryCreateMapping].Exec(url, code, singleUse, false)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (s *StoreStruct) GetUrlForCode(code string) (string, error) {
	var res queryResponse
	err := s.stmts[QueryFindMapping].QueryRow(code).Scan(&res.id, &res.originalUrl, &res.singleUse, &res.expired)

	if err != nil {
		return "", fmt.Errorf("%s was not found", code)
	}

	if res.singleUse == 1 {
		if res.expired == 1 {
			fmt.Println("expired")
			return "", fmt.Errorf("%s has already expired", code)
		} else {
			fmt.Println("will expire")

			_, err = s.stmts[QueryExpireMapping].Exec(1, res.id)
			return res.originalUrl, nil
		}
	}

	return res.originalUrl, nil
}
