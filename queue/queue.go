package queue

type MessageQueue struct {
	Type     int // Type Message: 1 Jurnal, 2 Hitung Ulang
	Payload  string
	UserID   uint32 // Parent ID
	CabangID uint32 // Cabang ID
}

type Jurnal struct {
	NoTransaksi   string     `json:"no_transaksi"`
	NamaTransaksi string     `json:"nama_transaksi"`
	NoReferensi   string     `json:"no_referensi"`
	IDCabang      uint64     `json:"id_cabang"`
	TglTransaksi  string     `json:"tgl_transaksi"`
	Status        int        `json:"status"`
	IsViewed      int        `json:"is_viewed"`
	Prefix        string     `json:"prefix"`
	TipeTransaksi int        `json:"tipe_transaksi"`
	Keterangan    string     `json:"keterangan"`
	ListInfoAkun  []ListAkun `json:"list_info_akun"`
}

type ListAkun struct {
	IDAkunting string  `json:"id_akunting"`
	Info       string  `json:"info"`
	Debet      float64 `json:"debet"`
	Kredit     float64 `json:"kredit"`
	ProdNo     string  `json:"prod_no"`
}
