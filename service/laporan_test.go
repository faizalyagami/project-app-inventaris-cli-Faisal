package service_test

import (
	"project-app-inventaris-cli-faisal/service"
	"testing"
)

func TestTampilkanTotalInvest(t *testing.T) {
	defer func ()  {
		if r := recover(); r != nil {
			t.Errorf("TampilkanTotalInvestasi panic: %v", r)
		}
	}()
	service.TampilkanTotalInvestasi()
}