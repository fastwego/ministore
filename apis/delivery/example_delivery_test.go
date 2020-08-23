package delivery_test

func ExampleDeliveryGetDeliveryCompanyList() {
	var ctx *ministore.Ministore
	payload := []byte("{}")
	resp, err := delivery.DeliveryGetDeliveryCompanyList(ctx, payload)
	fmt.Println(resp, err)
}

func ExampleDeliverySendDelivery() {
	var ctx *ministore.Ministore
	payload := []byte("{}")
	resp, err := delivery.DeliverySendDelivery(ctx, payload)
	fmt.Println(resp, err)
}
