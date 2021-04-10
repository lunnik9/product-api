package merch_repo

import (
	"fmt"
	"testing"

	"github.com/product-api/src/db"
)

func TestMerchMock_GetMerchByNameAndPassword(t *testing.T) {
	con, err := db.Connect("ec2-34-252-251-16.eu-west-1.compute.amazonaws.com:5432",
		"pnumlsyvxztrfm",
		"ee24c557c61258df433cfc825ea7e389ef53c907cb43195366c78f73d3c2acf4",
		"d1dlpo67q6hl95",
	)
	if err != nil {
		t.Fatal(err)
	}

	var mr = MerchPostgres{con}

	merch, err := mr.GetMerchByNameAndPassword("88005553535", "asd")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(merch)

}
