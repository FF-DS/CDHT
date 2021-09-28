package Database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

const tableSchema string = `
	CREATE TABLE dns_record (
		RecordType text,
		RecordKey text,
		RecordValue text
	);
`


type SqlDB struct {
	DatabaseName   string
	dbConn   *sql.DB
	insertStatement   *sql.Stmt
	readStatement   *sql.Stmt
	updateStatement  *sql.Stmt
	deleteStatement  *sql.Stmt
	tx *sql.Tx
	tried bool
}


// ## ------------- Connect ------------------ ##
func (sqlDB *SqlDB) connectToDB(){
	db, err := sql.Open("sqlite3", sqlDB.DatabaseName)
	if err != nil {
		return
	}
	sqlDB.dbConn = db
}



// ## ------------- Init ------------------ ##
func (sqlDB *SqlDB) Init() *SqlDB{
	sqlDB.connectToDB()
	sqlDB.createTable()
	sqlDB.initPS()

	return sqlDB
}



func (sqlDB *SqlDB) createTable(){
	_, err := sqlDB.dbConn.Exec(tableSchema)
	if err != nil {
		log.Printf("%q\n", err)
		return
	}
}


func (sqlDB *SqlDB) initPS(){
	tx, err := sqlDB.dbConn.Begin()
	if err != nil {
		log.Println(err)
		return
	}
	sqlDB.insertStatement, err = tx.Prepare("INSERT INTO dns_record(RecordType, RecordKey, RecordValue) values(?, ?, ?)")
	if err != nil {
		log.Println(err)
		return
	}

	sqlDB.readStatement, err = sqlDB.dbConn.Prepare("SELECT RecordType, RecordKey, RecordValue FROM dns_record where RecordKey = ? and RecordType = ? ")
	if err != nil {
		log.Println(err)
		return
	}

	sqlDB.updateStatement, err = sqlDB.dbConn.Prepare("UPDATE dns_record SET RecordValue = ? WHERE RecordKey = ? and RecordType = ? ")
	if err != nil {
		log.Println(err)
		return
	}

	sqlDB.deleteStatement, err = sqlDB.dbConn.Prepare("DELETE FROM dns_record WHERE RecordKey = ? and RecordType = ? ")
	if err != nil {
		log.Println(err)
		return
	}

	sqlDB.tx = tx
}




// ## ------------- Operations ------------------ ##
func (sqlDB *SqlDB) GetAllRecords() ([]DnsRecord, bool) {
	dnsRecords := []DnsRecord{}

	rows, err := sqlDB.dbConn.Query("select RecordType, RecordKey, RecordValue from dns_record")
	if err != nil {
		log.Println(err)
		return dnsRecords, false
	}
	defer rows.Close()


	for rows.Next() {
		record := DnsRecord{}
		err = rows.Scan(&record.RecordType, &record.RecordKey, &record.RecordValue)
		if err != nil {
			log.Println(err)
			return dnsRecords, false
		}
		dnsRecords = append(dnsRecords, record)
	}
	err = rows.Err()
	if err != nil {
		log.Println(err)
		return dnsRecords, false
	}

	return dnsRecords, true
}

func (sqlDB *SqlDB) InsertRecord(record DnsRecord) bool {
	_, err := sqlDB.insertStatement.Exec(record.RecordType, record.RecordKey, record.RecordValue)
	if err != nil {
		log.Println(err)
		sqlDB.initPS()
		sqlDB.InsertRecord(record)
		return false
	}
	sqlDB.tx.Commit()
	return true
}


func (sqlDB *SqlDB) FindRecord(recordKey string, recordType string) (DnsRecord, bool) {
	record := DnsRecord{}
	err := sqlDB.readStatement.QueryRow(recordKey, recordType).Scan(&record.RecordType, &record.RecordKey, &record.RecordValue)
	if err != nil {
		log.Println(err)
		return DnsRecord{}, false
	}

	return record,true
}

func (sqlDB *SqlDB) UpdateRecord(recordValue string, recordKey string, recordType string) bool {
	_, err := sqlDB.updateStatement.Exec(recordValue, recordKey, recordType)
	if err != nil {
		log.Println(err)
		return false
	}
	sqlDB.tx.Commit()
	return true
}


func (sqlDB *SqlDB) DeleteRecord(recordKey string, recordType string) bool {
	_, err := sqlDB.deleteStatement.Exec(recordKey, recordType)
	if err != nil {
		log.Println(err)
		return false
	}
	sqlDB.tx.Commit()
	return true
}


func (sqlDB *SqlDB) Close() {
	sqlDB.insertStatement.Close()
	sqlDB.readStatement.Close()
	sqlDB.dbConn.Close()
}


