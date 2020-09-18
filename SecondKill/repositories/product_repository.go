package repositories

import (
	"SecondKill/datamodels"
	"errors"
	"github.com/jinzhu/gorm"
)

type IProduct interface {
	//连接数据
	Insert(*datamodels.Product) (int64, error)
	Delete(int64) bool
	Update(*datamodels.Product) error
	SelectByKey(int64) (*datamodels.Product, error)
	SelectAll() ([]*datamodels.Product, error)
	SubProductNum(productID int64,userId int64) error
}

type product struct {
	mysqlConn *gorm.DB
}

func NewProduct(sqlDb *gorm.DB) *product {
	return &product{mysqlConn: sqlDb}
}

func (p *product) Insert(d *datamodels.Product) (int64, error) {
	result := p.mysqlConn.Create(d)
	return result.RowsAffected, result.Error
}

func (p *product) Delete(i int64) bool {
	var product datamodels.Product
	product.ID = uint(i)
	err := p.mysqlConn.Delete(&product)
	return err != nil
}

func (p *product) Update(d *datamodels.Product) error {
	return p.mysqlConn.Save(d).Error
}

func (p *product) SelectByKey(i int64) (*datamodels.Product, error) {
	var product datamodels.Product
	// 先查询
	e := p.mysqlConn.First(&product, i).Error
	if e != nil {
		e = errors.New("商品查询不到：" + e.Error())
	}

	return &product, e
}

func (p *product) SelectAll() ([]*datamodels.Product, error) {
	var products []*datamodels.Product
	e := p.mysqlConn.Find(&products).Error

	return products, e
}

func (p *product) SubProductNum(productID int64, userId int64) error {
	// 创建事物
	begin := p.mysqlConn.Begin()
	// 根据商品ID查询商品
	var product datamodels.Product
	if err := begin.First(&product, productID).Error; err != nil {
		begin.Rollback()
		return errors.New("查询订单错误：" + err.Error())
	}
	if product.ProductNum > 0 {
		// 扣除商品数量
		product.ProductNum -= 1
		if err := begin.Save(product).Error; err != nil {
			begin.Rollback()
			return errors.New("扣除商品数量错误：" + err.Error())
		}

		// 创建订单
		order := &datamodels.Order{
			UserId:      userId,
			ProductId:   productID,
			OrderStatus: datamodels.OrderSuccess,
		}
		if err := begin.Create(order).Error; err != nil {
			begin.Rollback()
			return errors.New("创建订单错误：" + err.Error())
		}
		// 无错误则提交事物
		begin.Commit()
		return nil
	} else {
		return errors.New("商品数量不足")
	}
}
