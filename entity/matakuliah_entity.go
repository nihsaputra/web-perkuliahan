package entity

type Matakuliah struct {
	KodeMk string `gorm:"column:kode_mk"`
	NamaMk string `gorm:"column:nama_mk"`
	Sks    int    `gorm:"column:sks"`
}
