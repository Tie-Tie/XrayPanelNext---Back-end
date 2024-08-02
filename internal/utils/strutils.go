package utils

import (
	"context"
	"crypto/md5"
	"database/sql"
	"fmt"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/google/uuid"
	"io"
	"io/ioutil"
	"math"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// md5 盐加密
func MD5V(password, salt string) string {
	combined := password + salt
	hasher := md5.New()
	io.WriteString(hasher, combined)
	return fmt.Sprintf("%x", hasher.Sum(nil))
}

// 消费订单生成 时间戳-套餐id-原价-优惠码-用户ID
func UseOrderNo(planId int, price float64, code string, userID int) string {
	return uuid.New().String()
	//fmt.Sprintf("%v-%v-%v-%v-%d", time.Now().Unix(), planId, price, code, userID)
}

// 充值订单生成 时间戳-充值金额(实际支付的)-payID-用户ID
func RechargeOrderNo(price float64, payId, userID int) string {
	return uuid.New().String()
	//fmt.Sprintf("%v-%v-%v-%d", time.Now().Unix(), price, payId, userID)
}

// bytes 转 GB
func BytesToGB(bytes int64) float64 {
	gigabytes := Decimal(float64(bytes) / 1073741824)
	return gigabytes
}

// GB 转 bytes
func GBToBytes(gigabytes float64) int64 {
	bytes := int64(gigabytes * 1073741824)
	return bytes
}

// 2个字符后所有显示*号
func MaskString(input string) string {
	if len(input) <= 2 {
		return input
	}
	// 使用 strings.Repeat 函数来生成星号(*)的部分
	masked := input[:2] + strings.Repeat("*", len(input)-2)
	return masked
}

// float64 只保留两位小数
func Decimal(value float64) float64 {
	value, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", value), 64)
	return value
}

// no rows in result set 错误判单
func IgnoreErrNoRows(err error) error {
	if err == sql.ErrNoRows {
		return nil
	} else {
		return err
	}
}

// 获取当前日期字符串 2023922
func GetDateNowStr() string {
	timeNow := time.Now()
	return fmt.Sprintf("%s%s%s", strconv.Itoa(timeNow.Year()), strconv.Itoa(int(timeNow.Month())), strconv.Itoa(timeNow.Day())) // = 2023922
}

// 获取当前日期字符串 2023922 - day
func GetDateNowMinusDayStr(day int) string {
	timeNow := time.Now()
	timeNow = timeNow.Add(-time.Duration(day) * 24 * time.Hour)

	return fmt.Sprintf("%s%s%s", strconv.Itoa(timeNow.Year()), strconv.Itoa(int(timeNow.Month())), strconv.Itoa(timeNow.Day())) // = 2023922
}

// 生成加密密码
func BcryptGeneratePassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// 密码效验
func BcryptCheckPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// 检查是否有特殊字符
func CheckStr(str string) bool {
	if strings.Contains(str, "'") ||
		strings.Contains(str, "\"") ||
		strings.Contains(str, "$") ||
		strings.Contains(str, "%") ||
		strings.Contains(str, "<") ||
		strings.Contains(str, ">") ||
		strings.Contains(str, "/") ||
		strings.Contains(str, "\\") ||
		strings.Contains(str, "#") ||
		strings.Contains(str, "&") {
		return true
	}
	return false
}

// ApiGet 自己封装的 GET 请求函数
func ApiGet(baseURL string, params url.Values, headers map[string]string) ([]byte, error) {
	Url, _ := url.Parse(baseURL)

	// 设置 URL 查询参数
	Url.RawQuery = params.Encode()

	// 创建请求
	req, err := http.NewRequest(http.MethodGet, Url.String(), nil)
	if err != nil {
		return nil, err
	}

	// 添加请求头
	for key, value := range headers {
		req.Header.Add(key, value)
	}

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			// 处理关闭连接时的错误
		}
	}(resp.Body)

	// 读取响应体
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// RoundToFixed rounds a float64 to the specified number of decimal places.
func RoundToFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return math.Round(num*output) / output
}

// ConvertToTimestamp 时间戳转化、验证时间是否超时
func ConvertToTimestamp(ctx context.Context, _order gdb.Record, deadline int64) int64 {
	// 转化时间戳
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		g.Log().Debug(ctx, "Trc20：加载时区失败")
		return -1
	}

	t, err := time.ParseInLocation("2006-01-02 15:04:05", gconv.String(_order["created_at"]), loc)
	if err != nil {
		g.Log().Debug(ctx, "Trc20：时间本地化失败")
		return -1
	}
	beginTimestamp := t.Unix()

	// 验证订单是否超时
	if beginTimestamp+deadline < time.Now().Unix() {
		_, err = g.Model("v2_recharge_records").Where("id=", _order["id"]).Update(g.Map{
			"status": 0,
		})
		g.Log().Debug(ctx, "Trc20：这笔交易已经超时了！")
		return 0
	}

	return beginTimestamp
}
