package Database

import (
	"fmt"
)

const (
	DNS_RECORD_A      string  =  "A"
	DNS_RECORD_AAAA   string  =  "AAAA"
	DNS_RECORD_ALIAS  string  =  "ALIAS"
	DNS_RECORD_CNAME  string  =  "CNAME"
	DNS_RECORD_MX     string  =  "MX"
	DNS_RECORD_NS     string  =  "NS"
	DNS_RECORD_PTR    string  =  "PTR"
	DNS_RECORD_SOA    string  =  "SOA"
	DNS_RECORD_SRV    string  =  "SRV"
	DNS_RECORD_TXT    string  =  "TXT"
)


type DnsRecord struct {
	RecordType string
	RecordKey string
	RecordValue string
}


// #--------------------------------- string log ---------------------------------# //
func  (record *DnsRecord) ToString() string {
	str := "    Record Data\n"  
	str += fmt.Sprintf("        [+] Record Type : %s\n", record.RecordType )
	str += fmt.Sprintf("        [+] Record Key : %s\n", record.RecordKey )
	str += fmt.Sprintf("        [+] Record Value : %s\n", record.RecordValue )
	return str;
}


func (record *DnsRecord) ToMap() map[string]string {
	return map[string]string{
		"RecordType" : record.RecordType,
		"RecordKey" : record.RecordKey,
		"RecordValue" : record.RecordValue,
	}
}

func ToRecord(data map[string]string) DnsRecord {
	return DnsRecord {
		RecordType: data["RecordType"],
		RecordKey: data["RecordKey"],
		RecordValue: data["RecordValue"],
	}
}