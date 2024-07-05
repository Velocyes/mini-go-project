Run app:
NGROK_AUTHTOKEN="2ipR5KiW7yYp6fcu9ECJze4degg_4qyw7odQvAigQ9FzBqXQc" go run cmd/app.go

Endpoints:

# Products #
- GetAllProducts:\n
  [GET] http://localhost:{port}/products
- GetProductsByIDs:\n
  [GET] http://localhost:{port}/products/{ids}
- InsertNewProducts:\n
  [POST] http://localhost:{port}/products
- AlterProducts:\n
  [PUT] http://localhost:{port}/products
- DeleteProductsByIDs:\n
  [DELETE] http://localhost:{port}/products/{ids}

# Orders #
- GetAllOrders:\n
  [GET] http://localhost:{port}/orders
- GetOrdersByIDs\n
  [GET] http://localhost:{port}/orders/{ids}
- InsertNewOrders:\n
  [POST] http://localhost:{port}/orders
- DeleteOrdersByIDs:\n
  [DELETE] http://localhost:{port}/orders/{ids}
  
Notes:
- Using ngrok to expose local to public

- Payment functionalities implemented using 'Midtrans Sandbox' but havn't completed yet (Only making transaction, the webhook for completing transaction is not finished yet)
- CI/CD is not implemented
