package order_test

func ExampleOrderGetOrderList() {
	var ctx *ministore.Ministore
	payload := []byte("{}")
	resp, err := order.OrderGetOrderList(ctx, payload)
	fmt.Println(resp, err)
}

func ExampleOrderGetOrderDetail() {
	var ctx *ministore.Ministore
	payload := []byte("{}")
	resp, err := order.OrderGetOrderDetail(ctx, payload)
	fmt.Println(resp, err)
}
