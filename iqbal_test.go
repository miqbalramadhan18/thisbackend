package iqbal

import (
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestInsertJalurPenerimaan(t *testing.T) {
	dbname := "qobel"
	jalurpenerimaan:= JalurPenerimaan{
		ID:      primitive.NewObjectID(),
		Jalurtanpates: "rapot",
		Jalurtes: "Mandiri",
		Beasiswa: "UTBK",
		Status: "Masih Berlaku",
	}
	insertedID := InsertJalurPenerimaan(dbname, jalurpenerimaan)
	if insertedID == nil {
		t.Error("Failed to insert user")
	}
}

func TestInsertInformasi(t *testing.T) {
	dbname := "qobel"
	informasi := Informasi{
		ID:       primitive.NewObjectID(),
		Alur: "Proses Pendaftaran",
		Catatan: "catatan",
	}
	insertedID := InsertInformasi(dbname, informasi)
	if insertedID == nil {
		t.Error("Failed to insert informasi")
	}
}


func TestInsertBiaya(t *testing.T) {
	dbname := "qobel"
	biaya := Biaya{
		ID:     primitive.NewObjectID(),
		Biayasemester: "7.700.000.00",
	}
	insertedID := InsertBiaya(dbname, biaya)
	if insertedID == nil {
		t.Error("Failed to insert biaya")
	}
}