package order

import (
	"0voiceGateway/config"
	"0voiceGateway/services/goods/proto"
	"context"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
)

type Order struct {
	config *config.Config
}

func NewOrderInstance(config *config.Config) *Order {
	return &Order{
		config: config,
	}
}

type OrderDetail struct {
	ID            int64   `json:"id"`
	UserID        int64   `json:"user_id"`
	UserName      string  `json:"user_name"`
	GoodsID       int64   `json:"goods_id"`
	GoodsName     string  `json:"goods_name"`
	GoodsCategory string  `json:"goods_category"`
	Amount        float64 `json:"amount"`
	Status        int     `json:"status"`
}

// 获取订单详情
func (o *Order) GetOrderDetail(c *gin.Context) {
	detail := o.getOrderDetail()
	c.JSON(http.StatusOK, detail)
}
func (o *Order) getOrderDetail() *OrderDetail {
	userID := o.config.Order.UserID
	user := o.getUser(userID)
	goodsID := o.config.Order.GoodsID
	goods := o.getGoods(goodsID)
	detail := &OrderDetail{
		ID:            o.config.Order.ID,
		UserID:        userID,
		UserName:      user.Name,
		GoodsID:       goodsID,
		GoodsName:     goods.Name,
		GoodsCategory: goods.Category,
		Amount:        o.config.Order.Amount,
		Status:        o.config.Order.Status,
	}
	return detail
}

type Goods struct {
	Id       int64
	Name     string
	Category string
}

func (o *Order) getGoods(goodsId int64) *Goods {
	goodsAddr := o.config.Goods.ServiceAddr
	conn, err := grpc.Dial(goodsAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println(err)
		return nil
	}
	defer conn.Close()

	client := proto.NewGoodsClient(conn)
	in := &proto.GetGoodsReq{
		GoodsId: o.config.Order.GoodsID,
	}
	res, err := client.Get(context.Background(), in)
	if err != nil {
		log.Println(err)
		return nil
	}
	return &Goods{
		Id:       goodsId,
		Name:     res.GoodsName,
		Category: res.GoodsCategory,
	}
}

type User struct {
	Id   int64
	Name string
}

func (o *Order) getUser(userId int64) *User {
	return &User{
		Id:   userId,
		Name: "nick",
	}
}
