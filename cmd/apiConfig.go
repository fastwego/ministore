// Copyright 2020 FastWeGo
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package main

type Param struct {
	Name string
	Type string
}

type Api struct {
	Name        string
	Description string
	Request     string
	See         string
	FuncName    string
	GetParams   []Param
}

type ApiGroup struct {
	Name    string
	Apis    []Api
	Package string
}

var apiConfig = []ApiGroup{
	{
		Name:    `服务商接口`,
		Package: `service`,
		Apis: []Api{
			{
				Name:        "登录验证",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/product/service/check_auth?component_access_token=xxxxxxxxx",
				See:         "https://developers.weixin.qq.com/doc/ministore/minishopopencomponent/API/service/check_auth.html",
				FuncName:    "ServiceCheckAuth",
			},
			{
				Name:        "获取用户购买的在有效期内的服务列表",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/product/service/get_list?access_token=xxxxxxxxx",
				See:         "https://developers.weixin.qq.com/doc/ministore/minishopopencomponent/API/service/get_service_list.html",
				FuncName:    "ServiceGetServiceList",
			},
			{
				Name:        "获取用户购买的服务列表",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/product/service/get_order_list?access_token=xxxxxxxxx",
				See:         "https://developers.weixin.qq.com/doc/ministore/minishopopencomponent/API/service/get_service_order_list.html",
				FuncName:    "ServiceGetServiceOrderList",
			},
		},
	},
	{
		Name:    `接入商品前必需接口`,
		Package: `cat`,
		Apis: []Api{
			{
				Name:        "获取类目详情",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/product/category/get?access_token=xxxxxxxxx",
				See:         "https://developers.weixin.qq.com/doc/ministore/minishopopencomponent/API/cat/get_cat_list.html",
				FuncName:    "CatGetCatList",
			},
			{
				Name:        "获取品牌列表",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/product/brand/get?access_token=xxxxxxxxx",
				See:         "https://developers.weixin.qq.com/doc/ministore/minishopopencomponent/API/cat/get_brand.html",
				FuncName:    "CatGetBrand",
			},
			{
				Name:        "获取运费模板",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/product/delivery/get_freight_template?access_token=xxxxxxxxx",
				See:         "https://developers.weixin.qq.com/doc/ministore/minishopopencomponent/API/cat/get_freight_template.html",
				FuncName:    "CatGetFreightTemplate",
			},
		},
	},
	{
		Name:    `SPU接口(修改需要重新上架商品)`,
		Package: `spu`,
		Apis: []Api{
			{
				Name:        "添加商品",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/product/spu/add?access_token=xxxxxxxxx",
				See:         "https://developers.weixin.qq.com/doc/ministore/minishopopencomponent/API/spu/add_spu.html",
				FuncName:    "SpuAddSpu",
			},
			{
				Name:        "删除商品",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/product/spu/del?access_token=xxxxxxxxx",
				See:         "https://developers.weixin.qq.com/doc/ministore/minishopopencomponent/API/spu/del_spu.html",
				FuncName:    "SpuDelSpu",
			},
			{
				Name:        "获取商品",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/product/spu/get?access_token=xxxxxxxxx",
				See:         "https://developers.weixin.qq.com/doc/ministore/minishopopencomponent/API/spu/get_spu.html",
				FuncName:    "SpuGetSpu",
			},
			{
				Name:        "获取商品列表",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/product/spu/get_list?access_token=xxxxxxxxx",
				See:         "https://developers.weixin.qq.com/doc/ministore/minishopopencomponent/API/spu/get_spu_list.html",
				FuncName:    "SpuGetSpuList",
			},
			{
				Name:        "更新商品",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/product/spu/update?access_token=xxxxxxxxx",
				See:         "https://developers.weixin.qq.com/doc/ministore/minishopopencomponent/API/spu/up_spu.html",
				FuncName:    "SpuUpSpu",
			},
			{
				Name:        "上架商品",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/product/spu/listing?access_token=xxxxxxxxx",
				See:         "https://developers.weixin.qq.com/doc/ministore/minishopopencomponent/API/spu/up_spu_listing.html",
				FuncName:    "SpuUpSpuListing",
			},
			{
				Name:        "下架商品",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/product/spu/delisting?access_token=xxxxxxxxx",
				See:         "https://developers.weixin.qq.com/doc/ministore/minishopopencomponent/API/spu/up_spu_delisting.html",
				FuncName:    "SpuUpSpuDelisting",
			},
		},
	},
	{
		Name:    `SKU接口(修改需要重新上架商品)`,
		Package: `sku`,
		Apis: []Api{
			{
				Name:        "添加SKU",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/product/sku/add?access_token=xxxxxxxxx",
				See:         "https://developers.weixin.qq.com/doc/ministore/minishopopencomponent/API/sku/add_sku.html",
				FuncName:    "SkuAddSku",
			},
			{
				Name:        "批量添加SKU",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/product/sku/batch_add?access_token=xxxxxxxxx",
				See:         "https://developers.weixin.qq.com/doc/ministore/minishopopencomponent/API/sku/batch_add_sku.html",
				FuncName:    "SkuBatchAddSku",
			},
			{
				Name:        "删除SKU",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/product/sku/del?access_token=xxxxxxxxx",
				See:         "https://developers.weixin.qq.com/doc/ministore/minishopopencomponent/API/sku/del_sku.html",
				FuncName:    "SkuDelSku",
			},
			{
				Name:        "获取SKU信息",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/product/sku/get?access_token=xxxxxxxxx",
				See:         "https://developers.weixin.qq.com/doc/ministore/minishopopencomponent/API/sku/get_sku.html",
				FuncName:    "SkuGetSku",
			},
			{
				Name:        "更新SKU",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/product/sku/update?access_token=xxxxxxxxx",
				See:         "https://developers.weixin.qq.com/doc/ministore/minishopopencomponent/API/sku/up_sku.html",
				FuncName:    "SkuUpSku",
			},
			{
				Name:        "更新SKU价格",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/product/sku/update_price?access_token=xxxxxxxxx",
				See:         "https://developers.weixin.qq.com/doc/ministore/minishopopencomponent/API/sku/up_sku_price.html",
				FuncName:    "SkuUpSkuPrice",
			},
			{
				Name:        "更新库存",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/product/sku/stock/update?access_token=xxxxxxxxx",
				See:         "https://developers.weixin.qq.com/doc/ministore/minishopopencomponent/API/sku/up_stock.html",
				FuncName:    "SkuUpStock",
			},
			{
				Name:        "获取库存",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/product/sku/stock/get?access_token=xxxxxxxxx",
				See:         "https://developers.weixin.qq.com/doc/ministore/minishopopencomponent/API/sku/get_stock.html",
				FuncName:    "SkuGetStock",
			},
		},
	},
	{
		Name:    `订单接口`,
		Package: `order`,
		Apis: []Api{
			{
				Name:        "获取订单列表",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/product/order/get_list?access_token=xxxxxxxxx",
				See:         "https://developers.weixin.qq.com/doc/ministore/minishopopencomponent/API/order/get_order_list.html",
				FuncName:    "OrderGetOrderList",
			},
			{
				Name:        "获取订单详情",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/product/order/get?access_token=xxxxxxxxx",
				See:         "https://developers.weixin.qq.com/doc/ministore/minishopopencomponent/API/order/get_order_detail.html",
				FuncName:    "OrderGetOrderDetail",
			},
		},
	},
	{
		Name:    `物流接口`,
		Package: `delivery`,
		Apis: []Api{
			{
				Name:        "获取快递公司列表",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/product/delivery/get_company_list?access_token=xxxxxxxxx",
				See:         "https://developers.weixin.qq.com/doc/ministore/minishopopencomponent/API/delivery/get_delivery_company_list.html",
				FuncName:    "DeliveryGetDeliveryCompanyList",
			},
			{
				Name:        "订单发货",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/product/delivery/send?access_token=xxxxxxxxx",
				See:         "https://developers.weixin.qq.com/doc/ministore/minishopopencomponent/API/delivery/send_delivery.html",
				FuncName:    "DeliverySendDelivery",
			},
		},
	},
	{
		Name:    `优惠券接口`,
		Package: `coupon`,
		Apis: []Api{
			{
				Name:        "获取优惠券列表",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/product/coupon/get_list?access_token=xxxxxxxxx",
				See:         "https://developers.weixin.qq.com/doc/ministore/minishopopencomponent/API/coupon/get_coupon.html",
				FuncName:    "CouponGetCoupon",
			},
			{
				Name:        "发放优惠券",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/product/coupon/push?access_token=xxxxxxxxx",
				See:         "https://developers.weixin.qq.com/doc/ministore/minishopopencomponent/API/coupon/push_coupon.html",
				FuncName:    "CouponPushCoupon",
			},
		},
	},
	{
		Name:    `店铺接口`,
		Package: `store`,
		Apis: []Api{
			{
				Name:        "获取基本信息",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/product/store/get_info?access_token=xxxxxxxxx",
				See:         "https://developers.weixin.qq.com/doc/ministore/minishopopencomponent/API/store/get_store_info.html",
				FuncName:    "StoreGetStoreInfo",
			},
		},
	},
}
