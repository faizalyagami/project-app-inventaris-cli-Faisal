package repository_test

import (
	"project-app-inventaris-cli-faisal/repository"
	"testing"
)

func TestHitungTotalInvestasiDanDepresiasi(t *testing.T) {
	totalInvestasi, totalDepresiasi := repository.HitungTotalInvestasiDanDepresiasi()

	if totalInvestasi <= 0 {
		t.Errorf("Expected totalInvestasi > 0, got %.2f", totalInvestasi)
	}

	if totalDepresiasi < 0 {
		t.Errorf("Expected totalDepresiasi >= 0, got %.2f", totalDepresiasi)
	}
}

func TestGetBarangPerluDiganti(t *testing.T) {
	list := repository.GetBarangPerluDiganti()

	if len(list) == 0 {
		t.Error("Expected at least one item in barang perlu diganti")
	}

	for _, item := range list {
		if item.UmurHari <= 100 {
			t.Errorf("Expected UmurHari > 100, got %d", item.UmurHari)
		}
	}
}

