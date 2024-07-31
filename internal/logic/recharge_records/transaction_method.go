package recharge_records

import (
	"context"
	"encoding/json"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"gov2panel/internal/type/transaction"
	"gov2panel/internal/utils"
	"math"
	"net/url"
	"time"
)

// 交易信息数据请求

// Trc20ApiGet Trc20的请求，contractAddress 合约地址，walletAddress 钱包地址
func Trc20ApiGet(ctx context.Context, contractAddress string, walletAddress string) (trc20Res transaction.Trc20Res, err error) {
	var trc20Result []byte
	if trc20Result, err = utils.ApiGet("https://nile.trongrid.io/v1/accounts/"+walletAddress+"/transactions/trc20", url.Values{
		"only_confirmed":   {"true"},
		"only_to":          {"true"},
		"limit":            {"200"},
		"contract_address": {contractAddress},
	}, g.MapStrStr{"TRON-PRO-API-KEY": "88aa8eeb-ccd2-48a6-86b7-2ac9c649f432"}); err != nil {
		g.Log().Error(ctx, "Trc20：交易列表请求失败！")
		return
	}

	if err = json.Unmarshal(trc20Result, &trc20Res); err != nil {
		g.Log().Error(ctx, "Trc20：请求转换json失败！")
		return
	}

	return
}

// TrxApiGet Trc20的请求，walletAddress 钱包地址
func TrxApiGet(ctx context.Context, walletAddress string) (trxRes transaction.TrxRes, err error) {

	var trxResult []byte
	if trxResult, err = utils.ApiGet("https://nile.trongrid.io/v1/accounts/"+walletAddress+"/transactions", url.Values{
		"only_confirmed": {"true"},
		"only_to":        {"true"},
		"limit":          {"200"},
	}, g.MapStrStr{"TRON-PRO-API-KEY": "88aa8eeb-ccd2-48a6-86b7-2ac9c649f432"}); err != nil {
		g.Log().Error(ctx, "Trx：交易列表请求失败！")
		return
	}

	if err = json.Unmarshal(trxResult, &trxRes); err != nil {
		g.Log().Error(ctx, "Trx：请求转换json失败！")
		return
	}

	return
}

// Erc20ApiGet Trc20的请求，contractAddress 合约地址，walletAddress 钱包地址
func Erc20ApiGet(ctx context.Context, contractAddress string, walletAddress string) (erc20Res transaction.Erc20Res, err error) {

	var erc20Result []byte
	if erc20Result, err = utils.ApiGet("https://api.etherscan.io/api", url.Values{
		"module":          {"account"},
		"action":          {"tokentx"},
		"contractaddress": {contractAddress},
		"address":         {walletAddress},
		"page":            {"1"},
		"offset":          {"1000"},
		"startblock":      {"0"},
		"endblock":        {"99999999"},
		"sort":            {"desc"},
		"apikey":          {"DEIYQVETP9XBYYG52PA4NBPJGXXNUNWGWC"},
	}, g.MapStrStr{}); err != nil {
		g.Log().Error(ctx, "Erc20：交易列表请求失败！")
		return
	}

	if err = json.Unmarshal(erc20Result, &erc20Res); err != nil {
		g.Log().Error(ctx, "Erc20：请求转换json失败！")
		return
	}

	return
}

// EthApiGet Trc20的请求，walletAddress 钱包地址
func EthApiGet(ctx context.Context, walletAddress string) (ethRes transaction.EthRes, err error) {
	var ethResult []byte
	if ethResult, err = utils.ApiGet("https://api-sepolia.etherscan.io/api", url.Values{
		"module":     {"account"},
		"action":     {"txlist"},
		"address":    {walletAddress},
		"startblock": {"0"},
		"endblock":   {"99999999"},
		"page":       {"1"},
		"offset":     {"1000"},
		"sort":       {"desc"},
		"apikey":     {"DEIYQVETP9XBYYG52PA4NBPJGXXNUNWGWC"},
	}, g.MapStrStr{}); err != nil {
		g.Log().Error(ctx, "Eth：交易列表请求失败！")
		return
	}

	if err = json.Unmarshal(ethResult, &ethRes); err != nil {
		g.Log().Error(ctx, "Eth：请求转换json失败！")
		return
	}

	return
}

// 交易验证函数

func EthVerify(ctx context.Context, ethRes transaction.EthRes, _order gdb.Record, deadline int64) {
	var err error

	for _, _row := range ethRes.Result {
		// 转换金额
		_rowAmount := gconv.Float64(_row.Value) * math.Pow(10, -18)
		_orderAmount := gconv.Float64(_order["amount"]) + gconv.Float64(_order["code"])*math.Pow(10, -6)

		// 转化时间戳
		var loc *time.Location
		var t time.Time
		loc, err = time.LoadLocation("Asia/Shanghai")
		t, err = time.ParseInLocation("2006-01-02 15:04:05", gconv.String(_order["created_at"]), loc)
		if err != nil {
			g.Log().Debug(ctx, "时间本地化失败")
			continue
		}
		beginTimestamp := t.Unix()

		// 验证订单是否超时
		if beginTimestamp+deadline < time.Now().Unix() {
			_, err = g.Model("v2_recharge_records").Where("id=", _order["id"]).Update(g.Map{
				"status": 0,
			})
			g.Log().Debug(ctx, "ERP20 USDT：这笔交易已经超时了！")
			break
		}

		// 验证交易时间
		if beginTimestamp > gconv.Int64(_row.TimeStamp) {
			g.Log().Debug(ctx, "ERP20 USDT：充值时间比该交易时间晚！")
			g.Dump(_row)
			break
		}

		// 验证交易金额
		if _orderAmount != _rowAmount {
			g.Log().Debug(ctx, "ERP20 USDT：验证交易金额失败！")
			continue
		}

		// 修改订单交易状态并将金额充值进入到用户账户中
		_, err = g.Model("v2_recharge_records").Where("id=", _order["id"]).Update(g.Map{
			"status": 2,
		})
		if err != nil {
			g.Log().Error(ctx, "ERP20 USDT：请求转换json失败！")
			continue
		}

		_, err = g.Model("v2_user").Where("id=", _order["user_id"]).Increment("balance", _orderAmount)
		g.Dump(_order["user_id"], _orderAmount)
		if err != nil {
			g.Log().Error(ctx, "ERP20 USDT：请求转换json失败！")
			continue
		}
		g.Log().Debug(ctx, "ERP20 USDT：这笔订单交易成功了！")
		break
	}
}

func TrxVerify(ctx context.Context, trc20Res transaction.Trc20Res, _order gdb.Record, deadline int64) {
	var err error

	for _, _row := range trc20Res.Data {
		// 转换金额
		_rowAmount := gconv.Float64(_row.Value) * math.Pow(10, -18)
		_orderAmount := gconv.Float64(_order["amount"]) + gconv.Float64(_order["code"])*math.Pow(10, -6)

		// 转化时间戳
		var loc *time.Location
		var t time.Time
		loc, err = time.LoadLocation("Asia/Shanghai")
		t, err = time.ParseInLocation("2006-01-02 15:04:05", gconv.String(_order["created_at"]), loc)
		if err != nil {
			g.Log().Debug(ctx, "时间本地化失败")
			continue
		}
		beginTimestamp := t.Unix()

		// 验证订单是否超时
		if beginTimestamp+deadline < time.Now().Unix() {
			_, err = g.Model("v2_recharge_records").Where("id=", _order["id"]).Update(g.Map{
				"status": 0,
			})
			g.Log().Debug(ctx, "Trc20：这笔交易已经超时了！")
			break
		}

		// 验证交易时间
		if beginTimestamp > gconv.Int64(_row.BlockTimestamp/1000) {
			g.Log().Debug(ctx, "Trc20：充值时间比该交易时间晚！")
			g.Dump(_row)
			break
		}

		// 验证交易金额
		if _orderAmount != _rowAmount {
			g.Log().Debug(ctx, "Trc20：验证交易金额失败！")
			continue
		}

		// 修改订单交易状态并将金额充值进入到用户账户中
		_, err = g.Model("v2_recharge_records").Where("id=", _order["id"]).Update(g.Map{
			"status": 2,
		})
		if err != nil {
			g.Log().Error(ctx, "Trc20：请求转换json失败！")
			continue
		}

		_, err = g.Model("v2_user").Where("id=", _order["user_id"]).Increment("balance", _orderAmount)
		g.Dump(_order["user_id"], _orderAmount)
		if err != nil {
			g.Log().Error(ctx, "Trc20：请求转换json失败！")
			continue
		}
		g.Log().Debug(ctx, "Trc20：这笔订单交易成功了！")
		break
	}
}

func Trc20Verify(ctx context.Context, trc20Res transaction.Trc20Res, _order gdb.Record, deadline int64) {
	var err error

	for _, _row := range trc20Res.Data {
		// 转换金额
		_rowAmount := gconv.Float64(_row.Value) * math.Pow(10, -18)
		_orderAmount := gconv.Float64(_order["amount"]) + gconv.Float64(_order["code"])*math.Pow(10, -6)

		// 转化时间戳
		var loc *time.Location
		var t time.Time
		loc, err = time.LoadLocation("Asia/Shanghai")
		t, err = time.ParseInLocation("2006-01-02 15:04:05", gconv.String(_order["created_at"]), loc)
		if err != nil {
			g.Log().Debug(ctx, "时间本地化失败")
			continue
		}
		beginTimestamp := t.Unix()

		// 验证订单是否超时
		if beginTimestamp+deadline < time.Now().Unix() {
			_, err = g.Model("v2_recharge_records").Where("id=", _order["id"]).Update(g.Map{
				"status": 0,
			})
			g.Log().Debug(ctx, "Trc20：这笔交易已经超时了！")
			break
		}

		// 验证交易时间
		if beginTimestamp > gconv.Int64(_row.BlockTimestamp/1000) {
			g.Log().Debug(ctx, "Trc20：充值时间比该交易时间晚！")
			g.Dump(_row)
			break
		}

		// 验证交易金额
		if _orderAmount != _rowAmount {
			g.Log().Debug(ctx, "Trc20：验证交易金额失败！")
			continue
		}

		// 修改订单交易状态并将金额充值进入到用户账户中
		_, err = g.Model("v2_recharge_records").Where("id=", _order["id"]).Update(g.Map{
			"status": 2,
		})
		if err != nil {
			g.Log().Error(ctx, "Trc20：请求转换json失败！")
			continue
		}

		_, err = g.Model("v2_user").Where("id=", _order["user_id"]).Increment("balance", _orderAmount)
		g.Dump(_order["user_id"], _orderAmount)
		if err != nil {
			g.Log().Error(ctx, "Trc20：请求转换json失败！")
			continue
		}
		g.Log().Debug(ctx, "Trc20：这笔订单交易成功了！")
		break
	}
}

func Erc20Verify(ctx context.Context, erc20Res transaction.Erc20Res, _order gdb.Record, deadline int64) {
	var err error

	for _, _row := range erc20Res.Result {
		// 转换金额
		_rowAmount := gconv.Float64(_row.Value) * math.Pow(10, -gconv.Float64(_row.TokenDecimal))
		_orderAmount := gconv.Float64(_order["amount"]) + gconv.Float64(_order["code"])*math.Pow(10, -6)

		// 转化时间戳
		var loc *time.Location
		var t time.Time
		loc, err = time.LoadLocation("Asia/Shanghai")
		t, err = time.ParseInLocation("2006-01-02 15:04:05", gconv.String(_order["created_at"]), loc)
		if err != nil {
			g.Log().Debug(ctx, "Erc20：时间本地化失败")
			continue
		}
		beginTimestamp := t.Unix()

		// 验证订单是否超时
		if beginTimestamp+deadline < time.Now().Unix() {
			_, err = g.Model("v2_recharge_records").Where("id=", _order["id"]).Update(g.Map{
				"status": 0,
			})
			g.Log().Debug(ctx, "Erc20：这笔交易已经超时了！")
			break
		}

		// 验证交易时间
		if beginTimestamp > gconv.Int64(_row.TimeStamp) {
			g.Log().Debug(ctx, "Erc20：充值时间比该交易时间晚！")
			g.Dump(_row)
			break
		}

		// 验证交易金额
		if _orderAmount != _rowAmount {
			g.Log().Debug(ctx, "Erc20：验证交易金额失败！")
			continue
		}

		// 修改订单交易状态并将金额充值进入到用户账户中
		_, err = g.Model("v2_recharge_records").Where("id=", _order["id"]).Update(g.Map{
			"status": 2,
		})
		if err != nil {
			g.Log().Error(ctx, "Erc20：请求转换json失败！")
			continue
		}

		_, err = g.Model("v2_user").Where("id=", _order["user_id"]).Increment("balance", _orderAmount)
		g.Dump(_order["user_id"], _orderAmount)
		if err != nil {
			g.Log().Error(ctx, "Erc20：请求转换json失败！")
			continue
		}
		g.Log().Debug(ctx, "Erc20：这笔订单交易成功了！")
		break
	}
}
