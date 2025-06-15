package repository

type LaporanBarang struct {
	ID         string
	Nama       string
	UmurHari   int
	Investasi  float64
	Depresiasi float64
}

func GetBarangPerluDiganti() []LaporanBarang {
	return []LaporanBarang{
		{ID: "B001", Nama: "Laptop A", UmurHari: 120},
		{ID: "B002", Nama: "Printer B", UmurHari: 150},
	}
}

func HitungTotalInvestasiDanDepresiasi() (totalInvestasi, totalDepresiasi float64) {
	return 10000000, 2500000
}

func GetLaporanBarangByID(id string) LaporanBarang {
	if id == "B001" {
		return LaporanBarang{ID: "B001", Nama: "Laptop A", Investasi: 5000000, Depresiasi: 1000000}
	}
	return LaporanBarang{}
}
