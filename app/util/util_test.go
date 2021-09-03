package util

import (
	"testing"
)

func TestExcelLie(t *testing.T) {
	t.Log(ExcelLie(100))
}

func TestMain(m *testing.M)  {
	m.Run()
}
