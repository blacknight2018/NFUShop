```mermaid
graph TB
    Gin--->User
    Gin--->Notify
    Notify--->Sub_Goods
    User--->Order
    User--->Address
    User--->Cart
    Order--->Sub_Goods
    Cart--->Sub_Goods
    Sub_Goods--->Result
    Result--->Utils
```