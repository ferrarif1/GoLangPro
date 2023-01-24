package go_cowswap

import (
	"context"
	"github.com/itsahedge/go-cowswap"
	"github.com/itsahedge/go-cowswap/util/signature-scheme/eip712"
	"testing"
)

func Test_VerifySignCancelOrder(t *testing.T) {
	client, err := go_cowswap.NewClient(go_cowswap.Options)
	if err != nil {
		t.Fatal(err)
	}
	uid := ""
	sig, typedData, err := client.SignCancelOrder(uid)
	if err != nil {
		t.Fatal(err)
	}
	////// CHECK SIGNATURE TO VERIFY OWNER
	hash, err := eip712.EncodeForSigning(*typedData)
	if err != nil {
		t.Logf("encode for signing err: %v", err)
	}
	checkAddress := client.TransactionSigner.SignerPubKey
	isOwner := eip712.VerifySig(checkAddress.Hex(), sig, hash.Bytes())
	t.Logf("order signature: %v", sig)
	t.Logf("typed data: %v", typedData)
	t.Logf("signature owner is verified: %v \n", isOwner)
}

func Test_CancelOrder(t *testing.T) {
	client, err := go_cowswap.NewClient(go_cowswap.Options)
	if err != nil {
		t.Fatal(err)
	}
	uid := ""
	res, statusCode, err := client.CancelOrder(context.Background(), uid)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("status code: %v", statusCode)
	t.Logf("res: %v", *res)
}

func Test_CancelOrders(t *testing.T) {
	client, err := go_cowswap.NewClient(go_cowswap.Options)
	if err != nil {
		t.Fatal(err)
	}
	uid_1 := ""
	uid_2 := ""
	uids := []string{uid_1, uid_2}
	res, statusCode, err := client.CancelOrders(context.Background(), uids)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("status code: %v", statusCode)
	t.Logf("res: %v", *res)
}
