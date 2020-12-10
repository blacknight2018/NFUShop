```mermaid
graph TB
    User--->Order
    User--->Address
    User--->Cart
    Order--->Sub_Goods
    Cart--->Sub_Goods
    Sub_Goods--->Result
    Result--->Utils

    Gin--->Service

    
    Service--->GetHotSubGoods--->SearchSubGoods
    Service--->GetNewestSubGoods--->SearchSubGoods
    Service--->GetSubGoods--->SearchSubGoods

    Service--->GetOrder--->GetSubGoods
    Service--->CreateOrder--->GetSubGoods
    Service--->DelCart
    Service--->GetCart--->GetSubGoods
    SearchSubGoods--->User
    

```