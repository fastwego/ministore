package cat_test

func ExampleCatGetCatList() {
	var ctx *ministore.Ministore
	payload := []byte("{}")
	resp, err := cat.CatGetCatList(ctx, payload)
	fmt.Println(resp, err)
}

func ExampleCatGetBrand() {
	var ctx *ministore.Ministore
	payload := []byte("{}")
	resp, err := cat.CatGetBrand(ctx, payload)
	fmt.Println(resp, err)
}

func ExampleCatGetFreightTemplate() {
	var ctx *ministore.Ministore
	payload := []byte("{}")
	resp, err := cat.CatGetFreightTemplate(ctx, payload)
	fmt.Println(resp, err)
}
