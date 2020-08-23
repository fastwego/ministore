package store_test

func ExampleStoreGetStoreInfo() {
	var ctx *ministore.Ministore
	payload := []byte("{}")
	resp, err := store.StoreGetStoreInfo(ctx, payload)
	fmt.Println(resp, err)
}
