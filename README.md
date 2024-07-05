Run app:
NGROK_AUTHTOKEN="2ipR5KiW7yYp6fcu9ECJze4degg_4qyw7odQvAigQ9FzBqXQc" go run cmd/app.go

Endpoints:

# Products #
- GetAllProducts:
  [GET] http://localhost:{port}/products
- GetProductsByIDs
  [GET] http://localhost:{port}/products/{ids}
- InsertNewProducts:
  [POST] http://localhost:{port}/products
- AlterProducts:
  [PUT] http://localhost:{port}/products
- DeleteProductsByIDs:
  [DELETE] http://localhost:{port}/products/{ids}

# Orders #
- GetAllOrders:
  [GET] http://localhost:{port}/orders
- GetOrdersByIDs
  [GET] http://localhost:{port}/orders/{ids}
- InsertNewOrders:
  [POST] http://localhost:{port}/orders
- DeleteOrdersByIDs:
  [DELETE] http://localhost:{port}/orders/{ids}
  
Notes:
- Using ngrok to expose local to public

- Payment functionalities implemented using 'Midtrans Sandbox' but havn't completed yet (Only making transaction, the webhook for completing transaction is not finished yet)
- CI/CD is not implemented
