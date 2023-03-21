package iqbal

import (
	"fmt"
	"testing"

	"github.com/aiteung/atdb"
)

var MongoInfo = atdb.DBInfo{
	DBString: MongoString,
	DBName:   "qobel",
}

var MongoConn = atdb.MongoConnect(MongoInfo)

// func TestInsertJalurPenerimaan(t *testing.T) {
// 	dbname := "qobel"
// 	jalurpenerimaan:= JalurPenerimaan{
// 		ID:      primitive.NewObjectID(),
// 		Jalurtanpates: "rapot",
// 		Jalurtes: "Mandiri",
// 		Beasiswa: "UTBK",
// 		Status: "Masih Berlaku",
// 	}
// 	insertedID := InsertJalurPenerimaan(dbname, jalurpenerimaan)
// 	if insertedID == nil {
// 		t.Error("Failed to insert user")
// 	}
// }

// func TestInsertInformasi(t *testing.T) {
// 	dbname := "qobel"
// 	informasi := Informasi{
// 		ID:       primitive.NewObjectID(),
// 		Alur: "Proses Pendaftaran",
// 		Catatan: "catatan",
// 	}
// 	insertedID := InsertInformasi(dbname, informasi)
// 	if insertedID == nil {
// 		t.Error("Failed to insert informasi")
// 	}
// }


// func TestInsertBiaya(t *testing.T) {
// 	dbname := "qobel"
// 	biaya := Biaya{
// 		ID:     primitive.NewObjectID(),
// 		Biayasemester: "7.700.000.00",
// 	}
// 	insertedID := InsertBiaya(dbname, biaya)
// 	if insertedID == nil {
// 		t.Error("Failed to insert biaya")
// 	}
// }

func TestGetBiaya(t *testing.T) {
	tahun := "7.700.000.00"
	hasil := GetBiaya(tahun, MongoConn, "biaya")
	fmt.Println(hasil)
}