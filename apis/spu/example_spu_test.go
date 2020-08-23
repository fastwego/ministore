package spu_test

func ExampleSpuAddSpu() {
	var ctx *ministore.Ministore
	payload := []byte("{}")
	resp, err := spu.SpuAddSpu(ctx, payload)
	fmt.Println(resp, err)
}

func ExampleSpuDelSpu() {
	var ctx *ministore.Ministore
	payload := []byte("{}")
	resp, err := spu.SpuDelSpu(ctx, payload)
	fmt.Println(resp, err)
}

func ExampleSpuGetSpu() {
	var ctx *ministore.Ministore
	payload := []byte("{}")
	resp, err := spu.SpuGetSpu(ctx, payload)
	fmt.Println(resp, err)
}

func ExampleSpuGetSpuList() {
	var ctx *ministore.Ministore
	payload := []byte("{}")
	resp, err := spu.SpuGetSpuList(ctx, payload)
	fmt.Println(resp, err)
}

func ExampleSpuUpSpu() {
	var ctx *ministore.Ministore
	payload := []byte("{}")
	resp, err := spu.SpuUpSpu(ctx, payload)
	fmt.Println(resp, err)
}

func ExampleSpuUpSpuListing() {
	var ctx *ministore.Ministore
	payload := []byte("{}")
	resp, err := spu.SpuUpSpuListing(ctx, payload)
	fmt.Println(resp, err)
}

func ExampleSpuUpSpuDelisting() {
	var ctx *ministore.Ministore
	payload := []byte("{}")
	resp, err := spu.SpuUpSpuDelisting(ctx, payload)
	fmt.Println(resp, err)
}
