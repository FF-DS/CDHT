
package DNS

import (
	"poc/app/DNS/Database"
	"fmt"
)


type DBHandler struct {
	db *Database.SqlDB
	DatabaseName string
}


func (dns *DBHandler) InitApp() *DBHandler{
	db := Database.SqlDB{DatabaseName: dns.DatabaseName}
	dns.db = db.Init()
	return dns
}


func (dns *DBHandler) ListCurrentData() []Database.DnsRecord {
	recordsList, status := dns.db.GetAllRecords()
	if !status {
		fmt.Println("Cant't get Records")
	}
	return recordsList
}


func (dns *DBHandler) AddRecords(record Database.DnsRecord) bool {
	status := dns.db.InsertRecord(record)

	if !status {
		fmt.Println("Cant't insert Records")
	}

	return status
}


func (dns *DBHandler) FindRecord(recordKey string, recordType string) Database.DnsRecord {
	record, status := dns.db.FindRecord(recordKey, recordType)

	if !status {
		fmt.Println("Cant't get Records")
	}

	return record
}


func (dns *DBHandler) UpdateRecord(recordValue string, recordKey string, recordType string) bool {
	status := dns.db.UpdateRecord(recordValue, recordKey, recordType)

	if !status {
		fmt.Println("Cant't update Records")
	}

	return status
}

func (dns *DBHandler) RemoveRecord(recordKey string, recordType string) bool {
	status := dns.db.DeleteRecord(recordKey, recordType)

	if !status {
		fmt.Println("Cant't delete Records")
	}

	return status
}


func (dns *DBHandler) Close() {
	dns.db.Close()
}