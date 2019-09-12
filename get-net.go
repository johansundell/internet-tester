package main

import (
	"errors"
	"net/http"
	"strconv"
	"time"
)

func checkInternet(url string) error {
	timeout := time.Duration(5 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	resp, err := client.Get(url)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return errors.New("Wrong status code " + strconv.Itoa(resp.StatusCode))
	}
	return nil
}
