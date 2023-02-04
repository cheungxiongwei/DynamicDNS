package main

import (
	"fmt"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	dnspod "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/dnspod/v20210323"
	"log"
)

func UpdateTencentIp(param CMDParam) {

	ipv4, err := GetLocalHostAddress()
	if err != nil {
		log.Printf("http get remote ipv4 failed\n")
		return
	}

	// 实例化一个认证对象，入参需要传入腾讯云账户 SecretId 和 SecretKey，此处还需注意密钥对的保密
	// 代码泄露可能会导致 SecretId 和 SecretKey 泄露，并威胁账号下所有资源的安全性。以下代码示例仅供参考，建议采用更安全的方式来使用密钥，请参见：https://cloud.tencent.com/document/product/1278/85305
	// 密钥可前往官网控制台 https://console.cloud.tencent.com/cam/capi 进行获取
	credential := common.NewCredential(
		"AKIDFRTO4YptjNm01nqqb7QXbgZEk13pQp1b",
		param.Password,
	)
	// 实例化一个client选项，可选的，没有特殊需求可以跳过
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "dnspod.tencentcloudapi.com"
	// 实例化要请求产品的client对象,clientProfile是可选的
	client, _ := dnspod.NewClient(credential, "", cpf)

	// 实例化一个请求对象,每个接口都会对应一个request对象
	request := dnspod.NewModifyDynamicDNSRequest()

	request.Domain = common.StringPtr(cmd.Domain)
	request.DomainId = common.Uint64Ptr(94001397)
	request.RecordId = common.Uint64Ptr(1336183831)
	request.RecordLine = common.StringPtr("默认")
	request.Value = common.StringPtr(ipv4)

	// 返回的resp是一个ModifyDynamicDNSResponse的实例，与请求对象对应
	response, err := client.ModifyDynamicDNS(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	// 输出json格式的字符串回包
	fmt.Printf("updating remote dns ip address: %s %s\n", ipv4, response.ToJsonString())
}
