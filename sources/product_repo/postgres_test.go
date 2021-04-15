package product_repo

import (
	"fmt"
	"testing"

	"github.com/lunnik9/product-api/domain"
	"github.com/lunnik9/product-api/sources/db"
)

func TestProductPostgres_Get(t *testing.T) {
	con, err := db.Connect("postgres://pnumlsyvxztrfm:ee24c557c61258df433cfc825ea7e389ef53c907cb43195366c78f73d3c2acf4@ec2-34-252-251-16.eu-west-1.compute.amazonaws.com:5432/d1dlpo67q6hl95")
	if err != nil {
		t.Fatal(err)
	}

	pr := ProductPostgres{con}

	product, err := pr.Get(1)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%+v", product)
}

func TestProductPostgres_Delete(t *testing.T) {
	con, err := db.Connect("postgres://pnumlsyvxztrfm:ee24c557c61258df433cfc825ea7e389ef53c907cb43195366c78f73d3c2acf4@ec2-34-252-251-16.eu-west-1.compute.amazonaws.com:5432/d1dlpo67q6hl95")
	if err != nil {
		t.Fatal(err)
	}

	pr := ProductPostgres{con}

	err = pr.Delete(1)
	if err != nil {
		t.Fatal(err)
	}
}

func TestProductPostgres_Create(t *testing.T) {
	con, err := db.Connect("postgres://pnumlsyvxztrfm:ee24c557c61258df433cfc825ea7e389ef53c907cb43195366c78f73d3c2acf4@ec2-34-252-251-16.eu-west-1.compute.amazonaws.com:5432/d1dlpo67q6hl95")
	if err != nil {
		t.Fatal(err)
	}

	pr := ProductPostgres{con}

	product := domain.Product{
		Barcode:       "3228024150199",
		Name:          "СЫР PRESIDENT RONDELE",
		StockId:       "123",
		Amount:        50,
		Unit:          "piece",
		PurchasePrice: 500,
		SellingPrice:  1000,
	}

	id, err := pr.Create(product)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(id)
}

func TestProductPostgres_Update(t *testing.T) {
	con, err := db.Connect("postgres://pnumlsyvxztrfm:ee24c557c61258df433cfc825ea7e389ef53c907cb43195366c78f73d3c2acf4@ec2-34-252-251-16.eu-west-1.compute.amazonaws.com:5432/d1dlpo67q6hl95")
	if err != nil {
		t.Fatal(err)
	}

	pr := ProductPostgres{con}

	product := domain.Product{
		Id:            3,
		Barcode:       "3228024150199",
		Name:          "СЫР PRESIDENT RONDELE123",
		StockId:       "123",
		Amount:        50,
		Unit:          "piece",
		PurchasePrice: 500,
		SellingPrice:  1000,
	}

	newProduct, err := pr.Update(product)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%+v", newProduct)
}
