# 微信小程序（Mini Program）

```go
"github.com/shenghui0779/gochat"
"github.com/shenghui0779/gochat/mp"
```

### 初始化小程序实例

```go
wxmp := gochat.NewMP(appid, appsecret)

// 如果启用了服务器配置，需要设置配置项
wxmp.SetServerConfig(token, encodingAESKey)
```

### 授权

```go
// 获取小程序授权的session_key
wxmp.Code2Session(ctx, code)

// 解密授权信息
wxmp.DecryptAuthInfo(session_key, iv, encrypted_data, dest)
```

### 接口调用凭据

```go
// 获取小程序的access_token
wxmp.AccessToken(ctx)
```

### 用户信息

```go
// 用户支付完成后，获取该用户的 UnionId，无需用户授权
// 微信支付订单号
wxmp.Do(ctx, access_token, mp.GetPaidUnionIDByTransactionID(openid, transactionID, dest))
// 微信支付商户订单号和微信支付商户号
wxmp.Do(ctx, access_token, mp.GetPaidUnionIDByOutTradeNO(openid, mchid, outTradeNO, dest))
```

### 消息

```go
// 发送统一服务消息
wxmp.Do(ctx, access_token, mp.SendUniformMessage(openid, msg))

// 发送订阅消息
wxmp.Do(ctx, access_token, mp.SendSubscribeMessage(openid, msg))

// 发送模板消息（已废弃，请使用订阅消息）
wxmp.Do(ctx, access_token, mp.SendTemplateMessage(openid, msg))

// 发送客服文本消息
wxmp.Do(ctx, access_token, mp.SendKFTextMessage(openid, msg))

// 客服图文链接消息
wxmp.Do(ctx, access_token, mp.SendKFImageMessage(openid, msg))

// 发送客服图文链接消息
wxmp.Do(ctx, access_token, mp.SendKFLinkMessage(openid, msg))

// 发送客服小程序卡片消息
wxmp.Do(ctx, access_token, mp.SendKFMinipMessage(openid, msg))

// 下发当前输入状态（仅支持客服消息）
wxmp.Do(ctx, access_token, mp.SetTyping(openid, msg))
```

### 插件管理

```go
// 向插件开发者发起使用插件的申请
wxmp.Do(ctx, access_token, mp.ApplyPlugin(pluginAppID, reason))

// 获取当前所有插件使用方（供插件开发者调用）
wxmp.Do(ctx, access_token, mp.GetPluginDevApplyList(page, num, dest))

// 查询已添加的插件
wxmp.Do(ctx, access_token, mp.GetPluginList(dest))

// 修改插件使用申请的状态（供插件开发者调用）
wxmp.Do(ctx, access_token, mp.SetDevPluginApplyStatus(action, appid, reason))

// 删除已添加的插件
wxmp.Do(ctx, access_token, mp.UnbindPlugin(pluginAppID))
```

### 小程序码

```go
// 创建小程序二维码（数量有限）
wxmp.Do(ctx, access_token, mp.CreateQRCode(path, dest, options...))

// 获取小程序二维码（数量有限）
wxmp.Do(ctx, access_token, mp.GetQRCode(path, dest, options...))

// 获取小程序二维码（数量不限）
wxmp.Do(ctx, access_token, mp.GetUnlimitQRCode(scene, dest, options...))
```

### 内容安全

```go
// 校验一张图片是否含有违法违规内容
wxmp.Do(ctx, access_token, mp.ImageSecCheck(filename))

// 异步校验图片/音频是否含有违法违规内容
wxmp.Do(ctx, access_token, mp.MediaSecCheckAsync(mediaType, mediaURL, dest))

// 检查一段文本是否含有违法违规内容
wxmp.Do(ctx, access_token, mp.MsgSecCheck(content))
```

### 图像处理

```go
// 图片智能裁切
wxmp.Do(ctx, access_token, mp.AICrop(filename, dest))
wxmp.Do(ctx, access_token, mp.AICropByURL(imgURL, dest))

// 条码/二维码识别
wxmp.Do(ctx, access_token, mp.ScanQRCode(filename, dest))
wxmp.Do(ctx, access_token, mp.ScanQRCodeByURL(imgURL, dest))

// 图片高清化
wxmp.Do(ctx, access_token, mp.SuperreSolution(filename, dest))
wxmp.Do(ctx, access_token, mp.SuperreSolutionByURL(imgURL, dest))
```

### OCR

```go
// 银行卡识别
wxmp.Do(ctx, access_token, mp.OCRBankCard(mode, filename, dest))
wxmp.Do(ctx, access_token, mp.OCRBankCardByURL(mode, imgURL, dest))

// 营业执照识别
wxmp.Do(ctx, access_token, mp.OCRBusinessLicense(mode, filename, dest))
wxmp.Do(ctx, access_token, mp.OCRBusinessLicenseByURL(mode, imgURL, dest))

// 身份证前面识别
wxmp.Do(ctx, access_token, mp.OCRIDCardFront(mode, filename, dest))
wxmp.Do(ctx, access_token, mp.OCRIDCardFrontByURL(mode, imgURL, dest))

// 身份证背面识别
wxmp.Do(ctx, access_token, mp.OCRIDCardBack(mode, filename, dest))
wxmp.Do(ctx, access_token, mp.OCRIDCardBackByURL(mode, imgURL, dest))

// 通用印刷体识别
wxmp.Do(ctx, access_token, mp.OCRPrintedText(mode, filename, dest))
wxmp.Do(ctx, access_token, mp.OCRPrintedTextByURL(mode, imgURL, dest))

// 行驶证识别
wxmp.Do(ctx, access_token, mp.OCRVehicleLicense(mode, filename, dest))
wxmp.Do(ctx, access_token, mp.OCRVehicleLicenseByURL(mode, imgURL, dest))
```

### 临时素材

```go
// 上传临时素材到微信服务器
wxmp.Do(ctx, access_token, mp.UploadMedia(mediaType, filename, dest))

// 获取客服消息内的临时素材
wxmp.Do(ctx, access_token, mp.GetMedia(mediaID, dest))
```

### 消息事件

```go
// 验证消息事件签名
wxmp.VerifyEventSign(signature, items...)

// 事件消息解密
wxmp.DecryptEventMessage(msg_encrypt)
```

### 其它

```go
// 调用服务平台提供的服务
wxmp.Do(ctx, access_token, mp.InvokeService(data, dest))

// 生物认证秘钥签名验证
wxmp.Do(ctx, access_token, mp.SoterVerify(sign, dest))

// 获取用户的安全等级（无需用户授权）
wxmp.Do(ctx, access_token, mp.GetUserRiskRank(data, dest))
```