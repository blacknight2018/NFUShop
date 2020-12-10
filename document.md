```mermaid
graph TB
    User--->Order
    User--->Address
    User--->Cart
    Order--->Sub_Goods
    Cart--->Sub_Goods
    Sub_Goods--->Result
    Result--->Utils

    Gin--->UID--->CID--->Api--->Service
    Gin--->UID--->GoodsID--->Api--->Service
    Gin--->UID--->AID--->Api--->Service
   
    
    Service--->GetHotSubGoods--->SearchSubGoods
    Service--->GetNewestSubGoods--->SearchSubGoods
    Service--->GetSubGoods--->SearchSubGoods

    Service--->GetOrder--->GetSubGoods
    Service--->CreateOrder--->GetSubGoods
    Service--->DelCart
    Service--->GetCart--->GetSubGoods
    SearchSubGoods--->User
    

```