package product_repo

import (
	"fmt"
	"reflect"
	"time"

	"github.com/go-pg/pg/v10"
	"github.com/lunnik9/product-api/domain"
	pe "github.com/lunnik9/product-api/product_errors"
)

type ProductPostgres struct {
	db *pg.DB
}

func NewProductPostgres(db *pg.DB) ProductPostgres {
	return ProductPostgres{db: db}
}

func (pr *ProductPostgres) Get(id int64) (*domain.Product, error) {
	var view = domain.ProductView{Id: id}

	err := pr.db.Model(&view).WherePK().Select()
	if err != nil {
		return nil, pe.New(409, err.Error())
	}

	product := domain.ProductViewToDomain(view)

	categories, err := pr.getCategoriesMapByIds(view.CategoryId)
	if err != nil {
		return nil, err
	}

	product.Category = categories[view.CategoryId].Name

	return &product, nil
}
func (pr *ProductPostgres) Delete(id int64) error {

	var view = domain.ProductView{Id: id}

	_, err := pr.db.Model(&view).WherePK().Delete()
	if err != nil {
		return pe.New(409, err.Error())
	}

	return nil
}
func (pr *ProductPostgres) Create(product domain.Product) (int64, error) {
	product.CreatedOn = time.Now().UTC()
	product.UpdatedOn = time.Now().UTC()

	view := domain.ProductDomainToView(product)

	_, err := pr.db.Model(&view).Returning("id").Insert()
	if err != nil {
		return 0, pe.New(409, err.Error())
	}

	return view.Id, nil
}
func (pr *ProductPostgres) Update(product domain.Product) (*domain.Product, error) {
	product.UpdatedOn = time.Now().UTC()

	view := domain.ProductDomainToView(product)

	_, err := pr.db.Model(&view).WherePK().Update()
	if err != nil {
		return nil, pe.New(409, err.Error())
	}

	return &product, nil
}
func (pr *ProductPostgres) Filter(limit, offset int, merchantId, stockId, name string) ([]domain.Product, error) {
	var (
		views          []domain.ProductView
		resp           []domain.Product
		query          string
		categoryIdsMap = make(map[int64]struct{})
		categoryIds    []int64
		categoriesMap  = make(map[int64]domain.Category)
	)

	query = "select * from product  where merchant_id = ? and stock_id = ?"
	if name != "" {
		name = "%" + name + "%"
		query += fmt.Sprintf(" and (name like '%v' or barcode like '%v')", name, name)
	}

	query = query + " order by updated_on desc limit ? offset ?"

	_, err := pr.db.Query(&views, query, merchantId, stockId, limit, offset)
	if err != nil {
		return nil, pe.New(409, err.Error())
	}

	if len(views) == 0 {
		return resp, nil
	}

	for _, view := range views {
		categoryIdsMap[view.CategoryId] = struct{}{}
	}

	for category := range categoryIdsMap {
		categoryIds = append(categoryIds, category)
	}

	categoriesMap, err = pr.getCategoriesMapByIds(categoryIds...)
	if err != nil {
		return nil, err
	}

	for _, view := range views {
		product := domain.ProductViewToDomain(view)
		product.Category = categoriesMap[view.CategoryId].Name
		resp = append(resp, product)
	}

	return resp, nil
}

func (pr *ProductPostgres) GetProductByBarcode(merchantId, stockId, barcode string) (*domain.Product, error) {
	var (
		view    domain.ProductView
		query   string
		product domain.Product
	)

	query = "select * from product  where merchant_id = ? and stock_id = ? and barcode = ?"

	_, err := pr.db.Query(&view, query, merchantId, stockId, barcode)
	if err != nil {
		return nil, pe.New(409, err.Error())
	}

	if reflect.ValueOf(view).IsZero() {
		return nil, nil
	}

	product = domain.ProductViewToDomain(view)

	return &product, nil

}

func (pr *ProductPostgres) getCategoriesMapByIds(ids ...int64) (map[int64]domain.Category, error) {
	var categoriesMap = make(map[int64]domain.Category)

	categories, err := pr.getCategoriesByIds(ids...)
	if err != nil {
		return nil, err
	}

	for _, category := range categories {
		categoriesMap[category.Id] = category
	}

	return categoriesMap, nil
}

func (pr *ProductPostgres) getCategoriesByIds(ids ...int64) ([]domain.Category, error) {
	var (
		views []domain.CategoryView
		resp  []domain.Category
	)

	err := pr.db.Model(&views).Where("id in (?)", pg.In(ids)).Select()
	if err != nil {
		return nil, pe.New(409, err.Error())
	}

	for _, view := range views {
		resp = append(resp, domain.CategoryViewToDomain(view))
	}

	return resp, nil
}

func (pr *ProductPostgres) GetCategory(id int64) (*domain.Category, error) {
	var view = domain.CategoryView{Id: id}

	err := pr.db.Model(&view).WherePK().Select()
	if err != nil {
		return nil, pe.New(409, err.Error())
	}

	category := domain.CategoryViewToDomain(view)

	return &category, nil
}

func (pr *ProductPostgres) DeleteCategory(id int64) error {
	var view = domain.CategoryView{Id: id}

	_, err := pr.db.Model(&view).WherePK().Delete()
	if err != nil {
		return pe.New(409, err.Error())
	}

	return nil
}

func (pr *ProductPostgres) CreateCategory(category domain.Category) (int64, error) {
	view := domain.CategoryDomainToView(category)

	_, err := pr.db.Model(&view).Returning("id").Insert()
	if err != nil {
		return 0, pe.New(409, err.Error())
	}

	return view.Id, nil
}

func (pr *ProductPostgres) UpdateCategory(category domain.Category) (*domain.Category, error) {

	view := domain.CategoryDomainToView(category)

	_, err := pr.db.Model(&view).WherePK().Update()
	if err != nil {
		return &category, pe.New(409, err.Error())
	}

	return &category, nil
}

func (pr *ProductPostgres) FilterCategory(limit, offset int, merchantId, stockId string) ([]domain.Category, error) {
	var (
		views      []domain.CategoryView
		categories []domain.Category
	)

	_, err := pr.db.Query(&views, "select * from category where merchant_id= ? and stock_id = ? limit ? offset ?",
		merchantId, stockId, limit, offset)
	if err != nil {
		return nil, pe.New(409, err.Error())
	}

	for _, view := range views {
		categories = append(categories, domain.CategoryViewToDomain(view))
	}

	return categories, nil
}

func (pr *ProductPostgres) MDelete(ids []int64) error {
	var (
		views []domain.ProductView
	)

	_, err := pr.db.Model(&views).Where("id in (?)", pg.In(ids)).Delete()
	if err != nil {
		return pe.New(409, err.Error())
	}

	return nil
}

func (pr *ProductPostgres) SaveTransfer(transfer domain.Transfer) error {
	product, err := pr.Get(transfer.ProductId)
	if err != nil {
		return err
	}

	switch transfer.Reason {
	case "received":
		product.Amount += transfer.Amount
	case "pulled":
		product.Amount -= transfer.Amount
	case "sold":
		product.Amount -= transfer.Amount
	default:
		return pe.New(409, fmt.Sprintf("invalid reason %v", transfer.Reason))
	}

	product.UpdatedOn = time.Now().UTC()
	transfer.Timestamp = time.Now().UTC()

	_, err = pr.Update(*product)
	if err != nil {
		return pe.New(409, err.Error())
	}

	err = pr.InsertTransfer(transfer)
	if err != nil {
		return pe.New(409, err.Error())
	}

	return nil
}

func (pr *ProductPostgres) InsertTransfer(transfer domain.Transfer) error {
	view := domain.TransferDomainToView(transfer)

	view.Timestamp = time.Now().UTC()

	_, err := pr.db.Model(&view).Insert()
	if err != nil {
		return pe.New(409, err.Error())
	}

	return nil
}

func (pr *ProductPostgres) GetTransfers(productId int64, limit, offset int) ([]domain.Transfer, error) {
	var (
		views     []domain.TransferView
		transfers []domain.Transfer
	)

	_, err := pr.db.Query(&views, "select * from transfer where product_id = ? limit ? offset ?",
		productId, limit, offset)
	if err != nil {
		return nil, pe.New(409, err.Error())
	}

	for _, view := range views {
		transfers = append(transfers, domain.TransferViewToDomain(view))
	}

	return transfers, nil
}

func (pr *ProductPostgres) Sync(merchantId, stockId string, lastUpdate time.Time) ([]domain.Product, error) {
	var (
		views []domain.ProductView
		resp  []domain.Product
		query string
	)

	query = "select * from product where updated_on > timestamp ? and merchant_id = ? and stock_id = ?"

	_, err := pr.db.Query(&views, query, lastUpdate.Format("2006-01-02 15:04:05.999"), merchantId, stockId)
	if err != nil {
		return nil, pe.New(409, err.Error())
	}

	for _, view := range views {
		resp = append(resp, domain.ProductViewToDomain(view))
	}

	return resp, nil

}
