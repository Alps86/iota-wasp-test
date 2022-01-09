package main

import (
	"fmt"
	"github.com/iotaledger/goshimmer/client/wallet/packages/seed"
	"github.com/iotaledger/wasp/client"
	"github.com/iotaledger/wasp/packages/iscp"
	"github.com/iotaledger/wasp/packages/iscp/request"
	"github.com/iotaledger/wasp/packages/iscp/requestargs"
	"github.com/iotaledger/wasp/packages/kv"
	"github.com/iotaledger/wasp/packages/kv/codec"
	"github.com/mr-tron/base58"
	"net/http"
	"time"
)

func main() {
	fmt.Println("start")

	c := client.NewWaspClient("http://localhost:9090", http.Client{})
	info, err := c.Info()
	checkError(err)

	fmt.Println(info)

	chain, err := iscp.ChainIDFromString("nAstjWkRwdUSA197pMQTYUzLnL4RpaLmzWnoedeCCL8Y")
	checkError(err)

	seedBytes, err := base58.Decode("3z3xSNwQLbkRd7NbhEBGgvnhgEToaRMwpbxSuzdMZzfT")
	checkError(err)

	s := seed.NewSeed(seedBytes)
	args := requestargs.New()
	req := request.NewOffLedger(chain, iscp.Hn("test"), iscp.Hn("increment"), args)
	req.Sign(s.KeyPair(0))

	err = c.PostOffLedgerRequest(chain, req)
	checkError(err)

	dic, err := c.CallView(chain, iscp.Hn("test"), "getCounter", nil, time.Second*30)
	checkError(err)

	counter, err := dic.Get("counter")
	checkError(err)

	fmt.Println(kv.Key(ValueToString("int", counter)))
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func ValueToString(vtype string, v []byte) string {
	switch vtype {
	case "address":
		addr, err := codec.DecodeAddress(v)
		checkError(err)
		return addr.Base58()
	case "agentid":
		aid, err := codec.DecodeAgentID(v)
		checkError(err)
		return aid.String()
	case "bool":
		b, err := codec.DecodeBool(v)
		checkError(err)
		if b {
			return "true"
		}
		return "false"
	case "bytes", "base58":
		return base58.Encode(v)
	case "chainid":
		cid, err := codec.DecodeChainID(v)
		checkError(err)
		return cid.String()
	case "color":
		col, err := codec.DecodeColor(v)
		checkError(err)
		return col.String()
	case "hash":
		hash, err := codec.DecodeHashValue(v)
		checkError(err)
		return hash.String()
	case "hname":
		hn, err := codec.DecodeHname(v)
		checkError(err)
		return hn.String()
	case "int8":
		n, err := codec.DecodeInt8(v)
		checkError(err)
		return fmt.Sprintf("%d", n)
	case "int16":
		n, err := codec.DecodeInt16(v)
		checkError(err)
		return fmt.Sprintf("%d", n)
	case "int32":
		n, err := codec.DecodeInt32(v)
		checkError(err)
		return fmt.Sprintf("%d", n)
	case "int64", "int":
		n, err := codec.DecodeInt64(v)
		checkError(err)
		return fmt.Sprintf("%d", n)
	case "requestid":
		rid, err := codec.DecodeRequestID(v)
		checkError(err)
		return rid.String()
	case "string":
		return fmt.Sprintf("%q", string(v))
	case "uint8":
		n, err := codec.DecodeUint8(v)
		checkError(err)
		return fmt.Sprintf("%d", n)
	case "uint16":
		n, err := codec.DecodeUint16(v)
		checkError(err)
		return fmt.Sprintf("%d", n)
	case "uint32":
		n, err := codec.DecodeUint32(v)
		checkError(err)
		return fmt.Sprintf("%d", n)
	case "uint64":
		n, err := codec.DecodeUint64(v)
		checkError(err)
		return fmt.Sprintf("%d", n)
	}

	return ""
}
