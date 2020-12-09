```mermaid
graph TB
    User--->Order
    User--->Address
    User--->Cart
    Order--->Goods
    Cart--->Goods
    Goods--->Result
    Result--->Utils
```