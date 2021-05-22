package waybill_repo

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/lunnik9/product-api/domain"
	"github.com/lunnik9/product-api/sources/db"
)

func TestWaybillPostgres_Create(t *testing.T) {
	con, err := db.Connect("postgres://pnumlsyvxztrfm:ee24c557c61258df433cfc825ea7e389ef53c907cb43195366c78f73d3c2acf4@ec2-34-252-251-16.eu-west-1.compute.amazonaws.com:5432/d1dlpo67q6hl95")
	if err != nil {
		t.Fatal(err)
	}

	wr := WaybillPostgres{con}

	waybill := domain.Waybill{
		MerchantId: "45",
		StockId:    "123",
		TotalCost:  100,
		Status:     "draft",
		Type:       "inwaybill",
		Number:     "228",
	}

	_, err = wr.Create(waybill)
	if err != nil {
		t.Fatal(err)
	}

	waybillJson, err := json.Marshal(waybill)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(string(waybillJson))

}

func TestWaybillPostgres_Filter(t *testing.T) {
	con, err := db.Connect("postgres://pnumlsyvxztrfm:ee24c557c61258df433cfc825ea7e389ef53c907cb43195366c78f73d3c2acf4@ec2-34-252-251-16.eu-west-1.compute.amazonaws.com:5432/d1dlpo67q6hl95")
	if err != nil {
		t.Fatal(err)
	}

	wr := WaybillPostgres{con}

	waybills, err := wr.Filter(10, 0, "inwaybill", "228", "45", "123")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%+v", waybills)
}

func TestWaybillPostgres_CreateProduct(t *testing.T) {
	con, err := db.Connect("postgres://pnumlsyvxztrfm:ee24c557c61258df433cfc825ea7e389ef53c907cb43195366c78f73d3c2acf4@ec2-34-252-251-16.eu-west-1.compute.amazonaws.com:5432/d1dlpo67q6hl95")
	if err != nil {
		t.Fatal(err)
	}

	wr := WaybillPostgres{con}

	product := domain.WaybillProduct{
		Name:          "opa",
		Barcode:       "1231231",
		PurchasePrice: 100,
		SellingPrice:  200,
		Amount:        100,
		WaybillId:     1,
		ProductId:     4,
	}

	_, err = wr.CreateProduct(product)
	if err != nil {
		t.Fatal(err)
	}

	waybillJson, err := json.Marshal(product)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(string(waybillJson))
}

func TestWaybillPostgres_GetList(t *testing.T) {
	con, err := db.Connect("postgres://pnumlsyvxztrfm:ee24c557c61258df433cfc825ea7e389ef53c907cb43195366c78f73d3c2acf4@ec2-34-252-251-16.eu-west-1.compute.amazonaws.com:5432/d1dlpo67q6hl95")
	if err != nil {
		t.Fatal(err)
	}

	wr := WaybillPostgres{con}

	products, err := wr.GetList(10, 0, 1)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%+v", products)

}
