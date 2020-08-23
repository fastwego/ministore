package service_test

func ExampleServiceCheckAuth() {
	var ctx *ministore.Ministore
	payload := []byte("{}")
	resp, err := service.ServiceCheckAuth(ctx, payload)
	fmt.Println(resp, err)
}

func ExampleServiceGetServiceList() {
	var ctx *ministore.Ministore
	payload := []byte("{}")
	resp, err := service.ServiceGetServiceList(ctx, payload)
	fmt.Println(resp, err)
}

func ExampleServiceGetServiceOrderList() {
	var ctx *ministore.Ministore
	payload := []byte("{}")
	resp, err := service.ServiceGetServiceOrderList(ctx, payload)
	fmt.Println(resp, err)
}
