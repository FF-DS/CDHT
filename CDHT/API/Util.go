package API

import (
    "crypto/sha1"
	"hash/crc32"
	"math/big"
    "fmt"
)




func (api *ApiCommunication) NameToNodeIDTranslator(name string) *big.Int {	
    hashFunction := sha1.New()
    hashFunction.Write([]byte(name))
    sha := hashFunction.Sum(nil)

	intBase := int64(api.NodeRoutingTable.NodeInfo().JumpSpacing)
    base, m, hashedID := big.NewInt(intBase), api.NodeRoutingTable.NodeInfo().M,  (&big.Int{}).SetBytes(sha)

    modulo := base.Exp( base, m, nil)
    return hashedID.Mod(hashedID, modulo)
}



func (api *ApiCommunication) DataToNodeIDTranslator(data []byte) *big.Int {	
    hashFunction := sha1.New()
    hashFunction.Write(data)
    sha := hashFunction.Sum(nil)

	intBase := int64(api.NodeRoutingTable.NodeInfo().JumpSpacing)
    base, m, hashedID := big.NewInt(intBase), api.NodeRoutingTable.NodeInfo().M,  (&big.Int{}).SetBytes(sha) 

    modulo := base.Exp( base, m, nil)
    return hashedID.Mod(hashedID, modulo)
}



func (api *ApiCommunication) CalculateCRCchecksum(data interface{}) uint32 {	
    crc32q := crc32.MakeTable(0xD5828281)
	return crc32.Checksum([]byte(fmt.Sprintf("%v", data)), crc32q)
}
