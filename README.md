# trade

Trade is a sample money transfer program that demonstrates how users log in, how to transfer funds to other accounts, and risk control for large transfers.

### Assumptions

* 假定一个客户有只有一个转账账户，且账户已经通过风控审计
* 转账中涉及汇率计算是非常复杂的，假定没有汇率影响，都是同一币种(CNY)
* 转账中手续费计算比较复杂，假定收费金额为100，并且在转账金额中扣除
* 假定用户转账时必须要输入交易密码增加安全性
* 假定转账金额3000,000触发风险预警，必须复核风控信息才能继续转换
* 假定客户转账金额没有当日累计金额限制
* 假定所有密码校验都通过，并且复核员ID为5555
* 假定所有上账和下帐都不用调用第三方API

###  API

* Login : 登录接口


```shell script
curl -H "Content-type: application/json" -X POST -d '{"name":"yulei"}'  http://127.0.0.1:9000/login
{"code":0,"data":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1OTg5Mjk0ODIsIm5hbWUiOiJ5dWxlaSIsInVpZCI6IjEifQ.Pf5gmKXZPQePYUwUcnze
BeiCgxoF7Ru7sPFJabfCNXk"}
```

* Add-Transfer : 转账接口

小额资金直接做账务处理

```shell script
 curl -X POST -H "Content-Type: application/json" -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.ey
JleHAiOjE1OTg5MDczNDMsIm5hbWUiOiJ5dWxlaSIsInVpZCI6IjEifQ.IHQgBBad17ih0EKBDkHrcsmMmy9djwd48dTsfUXzxj4"  http://127.0.0.1:9000/v1/add-transfer -d '{"uuid":"bbacb7ca-9dad-12d1-80b4-88c04fd782","to_account":"6cbc0bf671915aa36970ade080f4c474","to_name":"toone","from_account":"6c84c4b8c2f1c1a15feb6967578049bb","money_type":"CNY","money_amt":1200,"password":"xxxxxx","comment":"test-transfer","is_realtime":true,"op_datetime":"2020-06-04T19:25:16","op_timezone":"Asia/Shanghai","notification_type":"email","postscript":"hello!"}'{"code":0,"data":{"trans_date":20200831,"trans_time":140400,"uuid":"09400c10-f8b4-46b9-bcce-acadcffa0e1d","from_uuid":"bbacb7ca-9dad-1
2d1-80b4-88c04fd782","to_uuid":"09400c10-f8b4-46b9-bcce-acadcffa0e1d","from_bid":1,"from_cid":1,"to_bid":2,"to_cid":2,"is_delay":true,"is_large":false,"is_reject":false,"amt":1200,"fee":100,"remark":"test-transfer","money_type":"CNY","errcode":9999,"confirm_date":0,"confirm_time":0,"confirm_amt":0,"confirm_opid":0}}

```

大额转账接口, 资金先冻结，调用转账确认接口才会真正资金划转

```shell script

curl -X POST -H "Content-Type: application/json" -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.ey
JleHAiOjE1OTg5MDczNDMsIm5hbWUiOiJ5dWxlaSIsInVpZCI6IjEifQ.IHQgBBad17ih0EKBDkHrcsmMmy9djwd48dTsfUXzxj4"  http://127.0.0.1:9000/v1/add-transfer -d '{"uuid":"2299b789-7a11-12d1-73b4-4cc84fd2678","to_account":"6cbc0bf671915aa36970ade080f4c474","to_name":"toone","from_account":"6c84c4b8c2f1c1a15feb6967578049bb","money_type":"CNY","money_amt":9000000,"password":"xxxxxx","comment":"test-transfer","is_realtime":true,"op_datetime":"2020-06-04T19:25:16","op_timezone":"Asia/Shanghai","notification_type":"email","postscript":"hello!"}'{"code":0,"data":{"trans_date":20200831,"trans_time":140514,"uuid":"41d670d2-d430-4640-98d0-22de2b17682a","from_uuid":"2299b789-7a11-1
2d1-73b4-4cc84fd2678","to_uuid":"41d670d2-d430-4640-98d0-22de2b17682a","from_bid":1,"from_cid":1,"to_bid":2,"to_cid":2,"is_delay":true,"is_large":false,"is_reject":false,"amt":9000000,"fee":100,"remark":"test-transfer","money_type":"CNY","errcode":6666,"confirm_date":0,"confirm_time":0,"confirm_amt":0,"confirm_opid":0}}

```

* confirm-transfer : 转账确认接口

```shell script
curl -X POST -H "Content-Type: application/json" -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.ey
JleHAiOjE1OTg5MDczNDMsIm5hbWUiOiJ5dWxlaSIsInVpZCI6IjEifQ.IHQgBBad17ih0EKBDkHrcsmMmy9djwd48dTsfUXzxj4"  http://127.0.0.1:9000/v1/confirm-transfer -d '{"uuid":"41d670d2-d430-4640-98d0-22de2b17682a","money_type":"CNY","money_amt":9000000,"op_name":"qq111","risk_level":3,"comment":"me333sl"}'{"code":0,"data":{"trans_date":20200831,"trans_time":140514,"uuid":"41d670d2-d430-4640-98d0-22de2b17682a","from_uuid":"2299b789-7a11-1
2d1-73b4-4cc84fd2678","to_uuid":"41d670d2-d430-4640-98d0-22de2b17682a","from_bid":1,"from_cid":1,"to_bid":2,"to_cid":2,"is_delay":true,"is_large":false,"is_reject":false,"amt":9000000,"fee":100,"remark":"test-transfer","money_type":"CNY","errcode":9999,"confirm_date":20200831,"confirm_time":140922,"confirm_amt":9000000,"confirm_opid":5555}}
```

* transfers : 转账明细查询

```shell script
curl http://127.0.0.1:9000/v1/transfers -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1OTg5MDczNDMsIm5hbWUiOiJ5dWxlaSIsInVpZCI6IjEifQ.IHQgBBad17ih0EKBDkHrcsmMmy9djwd48dTsfUXzxj4"
{"code":0,"data":[{"trans_date":20200831,"trans_time":140514,"uuid":"41d670d2-d430-4640-98d0-22de2b17682a","from_uuid":"2299b789-7a11-12d1-73b4-4cc84fd2678","to_uuid":"41d670d2-d430-4640-98d0-22de2b17682a","from_bid":1,"from_cid":1,"to_bid":2,"to_cid":2,"is_delay":true,"is_large":false,"is_reject":false,"amt":9000000,"fee":100,"remark":"test-transfer","money_type":"CNY","errcode":9999,"confirm_date":20200831,"confirm_time":140922,"confirm_amt":9000000,"confirm_opid":5555,"fees":[{"trans_date":0,"trans_time":0,"uuid":"41d670d2-d430-4640-98d0-22de2b17682a","money_type":"CNY","amt":9000000,"fee":100,"remark":"test-transfer"}]}]}
```



###  license

trade is licensed under the Apache License 2.0
 