package elog_test

import (
	"errors"
	"os"
	"testing"
	"time"

	"github.com/prinick96/elog"
)

var log_dir = "./logs/"

func TestNew(t *testing.T) {
	testCases := []struct {
		Name  string
		Str   string
		Err   error
		Level uint8
	}{
		{
			Name:  "it should not write a error with panic",
			Str:   "test-error",
			Err:   nil,
			Level: elog.PANIC,
		},
		{
			Name:  "it should not write a error without panic",
			Str:   "test-error",
			Err:   nil,
			Level: elog.ERROR,
		},
		{
			Name:  "it should write a error with panic",
			Str:   "test-error",
			Err:   errors.New("self generated error"),
			Level: elog.PANIC,
		},
		{
			Name:  "it should write a error without panic",
			Str:   "test-error",
			Err:   errors.New("self generated error"),
			Level: elog.ERROR,
		},
	}

	for i := range testCases {
		tc := testCases[i]

		defer func() {
			if err := recover(); tc.Level != elog.PANIC && err != nil {
				t.Error("panic not wanted")
			}
		}()

		elog.New(tc.Level, tc.Str, tc.Err)

		strTime := time.Now().Format("02-January-2006")
		f_name := log_dir + strTime + ".log"

		if _, err := os.Stat(log_dir + f_name); tc.Err != nil && os.IsNotExist(err) {
			t.Error(tc.Name, " Expected file "+f_name+" to exist")
		}
	}
}

func TestClean(t *testing.T) {
	os.RemoveAll(log_dir)
}
