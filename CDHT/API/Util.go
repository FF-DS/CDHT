package API

import (
	"math/big"
    "crypto/sha1"
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