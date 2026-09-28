package ticket
import("crypto/rand";"fmt";"math/big")
func ID()string{n,_:=rand.Int(rand.Reader,big.NewInt(900000));return fmt.Sprintf("PASS-%06d",n.Int64()+100000)}
