package merch_repo

import (
	"fmt"
	"testing"
	"time"

	"github.com/lunnik9/product-api/domain"
	"github.com/lunnik9/product-api/sources/db"
)

func TestMerchPostgres_GetMerchByNameAndPassword(t *testing.T) {
	con, err := db.Connect("postgres://pnumlsyvxztrfm:ee24c557c61258df433cfc825ea7e389ef53c907cb43195366c78f73d3c2acf4@ec2-34-252-251-16.eu-west-1.compute.amazonaws.com:5432/d1dlpo67q6hl95")
	if err != nil {
		t.Fatal(err)
	}

	var mr = MerchPostgres{con}

	merch, err := mr.GetMerchByNameAndPassword("7273765595", "zxc")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v", merch)

}

func TestMerchPostgres_GetMerchByToken(t *testing.T) {
	con, err := db.Connect("postgres://pnumlsyvxztrfm:ee24c557c61258df433cfc825ea7e389ef53c907cb43195366c78f73d3c2acf4@ec2-34-252-251-16.eu-west-1.compute.amazonaws.com:5432/d1dlpo67q6hl95")
	if err != nil {
		t.Fatal(err)
	}

	var mr = MerchPostgres{con}

	merch, err := mr.GetMerchByToken("123")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%+v", merch)
}

func TestMerchPostgres_UpdateMerch(t *testing.T) {
	con, err := db.Connect("postgres://pnumlsyvxztrfm:ee24c557c61258df433cfc825ea7e389ef53c907cb43195366c78f73d3c2acf4@ec2-34-252-251-16.eu-west-1.compute.amazonaws.com:5432/d1dlpo67q6hl95")
	if err != nil {
		t.Fatal(err)
	}

	var mr = MerchPostgres{con}

	merch := domain.Merchant{
		MerchantId:   "45",
		MerchantName: "arnur",
		Password:     "zxc",
		Token:        "456",
		UpdateTime:   time.Now(),
		TokenTTL:     10000,
		LastCheck:    time.Now(),
		Mobile:       "7273765595",
	}

	err = mr.UpdateMerch(merch)
	if err != nil {
		t.Fatal(err)
	}
}

func TestMerchPostgres_CheckRights(t *testing.T) {
	con, err := db.Connect("postgres://pnumlsyvxztrfm:ee24c557c61258df433cfc825ea7e389ef53c907cb43195366c78f73d3c2acf4@ec2-34-252-251-16.eu-west-1.compute.amazonaws.com:5432/d1dlpo67q6hl95")
	if err != nil {
		t.Fatal(err)
	}

	var mr = MerchPostgres{con}

	fmt.Println(mr.CheckRights("456"))
}

func TestMerchPostgres_GetListOfCashBoxes(t *testing.T) {
	con, err := db.Connect("postgres://pnumlsyvxztrfm:ee24c557c61258df433cfc825ea7e389ef53c907cb43195366c78f73d3c2acf4@ec2-34-252-251-16.eu-west-1.compute.amazonaws.com:5432/d1dlpo67q6hl95")
	if err != nil {
		t.Fatal(err)
	}

	var mr = MerchPostgres{con}

	cashBoxes, err := mr.GetListOfCashBoxes("45", "123")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%+v", cashBoxes)

}
