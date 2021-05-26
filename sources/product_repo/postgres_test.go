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

	product, err := pr.Get(4)
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

func TestProductPostgres_Filter(t *testing.T) {
	con, err := db.Connect("postgres://pnumlsyvxztrfm:ee24c557c61258df433cfc825ea7e389ef53c907cb43195366c78f73d3c2acf4@ec2-34-252-251-16.eu-west-1.compute.amazonaws.com:5432/d1dlpo67q6hl95")
	if err != nil {
		t.Fatal(err)
	}

	pr := ProductPostgres{con}

	products, err := pr.Filter(10, 0, "45", "1233", "СЫР")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%+v", products)

}

func TestProductPostgres_GetProductByBarcode(t *testing.T) {
	con, err := db.Connect("postgres://pnumlsyvxztrfm:ee24c557c61258df433cfc825ea7e389ef53c907cb43195366c78f73d3c2acf4@ec2-34-252-251-16.eu-west-1.compute.amazonaws.com:5432/d1dlpo67q6hl95")
	if err != nil {
		t.Fatal(err)
	}

	pr := ProductPostgres{con}

	product, err := pr.GetProductByBarcode("45", "123", "125654323456323")

	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%+v", product)
}

func TestProductPostgres_CreateCategory(t *testing.T) {
	con, err := db.Connect("postgres://pnumlsyvxztrfm:ee24c557c61258df433cfc825ea7e389ef53c907cb43195366c78f73d3c2acf4@ec2-34-252-251-16.eu-west-1.compute.amazonaws.com:5432/d1dlpo67q6hl95")
	if err != nil {
		t.Fatal(err)
	}

	pr := ProductPostgres{con}

	id, err := pr.CreateCategory(domain.Category{
		MerchantId: "45",
		StockId:    "123",
		Name:       "Eco",
	})
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(id)
}

func TestProductPostgres_DeleteCategory(t *testing.T) {

}

func TestProductPostgres_FilterCategory(t *testing.T) {
	con, err := db.Connect("postgres://pnumlsyvxztrfm:ee24c557c61258df433cfc825ea7e389ef53c907cb43195366c78f73d3c2acf4@ec2-34-252-251-16.eu-west-1.compute.amazonaws.com:5432/d1dlpo67q6hl95")
	if err != nil {
		t.Fatal(err)
	}

	pr := ProductPostgres{con}

	categories, err := pr.FilterCategory(10, 0, "45", "123")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%+v", categories)
}

func TestProductPostgres_UpdateCategory(t *testing.T) {

}

func TestProductPostgres_GetCategory(t *testing.T) {
	con, err := db.Connect("postgres://pnumlsyvxztrfm:ee24c557c61258df433cfc825ea7e389ef53c907cb43195366c78f73d3c2acf4@ec2-34-252-251-16.eu-west-1.compute.amazonaws.com:5432/d1dlpo67q6hl95")
	if err != nil {
		t.Fatal(err)
	}

	pr := ProductPostgres{con}

	category, err := pr.GetCategory(1)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%+v", category)

}

func TestProductPostgres_MDelete(t *testing.T) {
	con, err := db.Connect("postgres://pnumlsyvxztrfm:ee24c557c61258df433cfc825ea7e389ef53c907cb43195366c78f73d3c2acf4@ec2-34-252-251-16.eu-west-1.compute.amazonaws.com:5432/d1dlpo67q6hl95")
	if err != nil {
		t.Fatal(err)
	}

	pr := ProductPostgres{con}

	err = pr.MDelete([]int64{7, 8})
	if err != nil {
		t.Fatal(err)
	}
}

func TestProductPostgres_SaveTransfer(t *testing.T) {
	con, err := db.Connect("postgres://pnumlsyvxztrfm:ee24c557c61258df433cfc825ea7e389ef53c907cb43195366c78f73d3c2acf4@ec2-34-252-251-16.eu-west-1.compute.amazonaws.com:5432/d1dlpo67q6hl95")
	if err != nil {
		t.Fatal(err)
	}

	pr := ProductPostgres{con}

	err = pr.SaveTransfer(domain.Transfer{
		ProductId:     4,
		SellingPrice:  100,
		PurchasePrice: 500,
		Amount:        10,
		Reason:        "sold",
		Source:        "waybill",
		SourceId:      "100",
	})

	if err != nil {
		t.Fatal(err)
	}
}

func TestProductPostgres_GetTransfers(t *testing.T) {
	con, err := db.Connect("postgres://pnumlsyvxztrfm:ee24c557c61258df433cfc825ea7e389ef53c907cb43195366c78f73d3c2acf4@ec2-34-252-251-16.eu-west-1.compute.amazonaws.com:5432/d1dlpo67q6hl95")
	if err != nil {
		t.Fatal(err)
	}

	pr := ProductPostgres{con}

	transfers, err := pr.GetTransfers(4, 10, 0)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%+v", transfers)

}
