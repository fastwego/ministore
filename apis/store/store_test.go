package store

func TestMain(m *testing.M) {
	test.Setup()
	os.Exit(m.Run())
}

func TestStoreGetStoreInfo(t *testing.T) {
	mockResp := map[string][]byte{
		"case1": []byte("{\"errcode\":0,\"errmsg\":\"ok\"}"),
	}
	var resp []byte
	test.MockSvrHandler.HandleFunc(apiStoreGetStoreInfo, func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(resp))
	})
	type args struct {
		ctx     *ministore.Ministore
		payload []byte
	}
	tests := []struct {
		name     string
		args     args
		wantResp []byte
		wantErr  bool
	}{
		{name: "case1", args: args{ctx: test.MockMinistore}, wantResp: mockResp["case1"], wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp = mockResp[tt.name]
			gotResp, err := StoreGetStoreInfo(tt.args.ctx, tt.args.payload)
			//fmt.Println(string(gotResp), err)
			if (err != nil) != tt.wantErr {
				t.Errorf("StoreGetStoreInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("StoreGetStoreInfo() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
