package sku_test

func ExampleSkuAddSku() {
	var ctx *ministore.Ministore
	payload := []byte("{}")
	resp, err := sku.SkuAddSku(ctx, payload)
	fmt.Println(resp, err)
}

func ExampleSkuBatchAddSku() {
	var ctx *ministore.Ministore
	payload := []byte("{}")
	resp, err := sku.SkuBatchAddSku(ctx, payload)
	fmt.Println(resp, err)
}

func ExampleSkuDelSku() {
	var ctx *ministore.Ministore
	payload := []byte("{}")
	resp, err := sku.SkuDelSku(ctx, payload)
	fmt.Println(resp, err)
}

func ExampleSkuGetSku() {
	var ctx *ministore.Ministore
	payload := []byte("{}")
	resp, err := sku.SkuGetSku(ctx, payload)
	fmt.Println(resp, err)
}

func ExampleSkuUpSku() {
	var ctx *ministore.Ministore
	payload := []byte("{}")
	resp, err := sku.SkuUpSku(ctx, payload)
	fmt.Println(resp, err)
}

func ExampleSkuUpSkuPrice() {
	var ctx *ministore.Ministore
	payload := []byte("{}")
	resp, err := sku.SkuUpSkuPrice(ctx, payload)
	fmt.Println(resp, err)
}

func ExampleSkuUpStock() {
	var ctx *ministore.Ministore
	payload := []byte("{}")
	resp, err := sku.SkuUpStock(ctx, payload)
	fmt.Println(resp, err)
}

func ExampleSkuGetStock() {
	var ctx *ministore.Ministore
	payload := []byte("{}")
	resp, err := sku.SkuGetStock(ctx, payload)
	fmt.Println(resp, err)
}
