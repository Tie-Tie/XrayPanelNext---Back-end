package public

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	v1 "gov2panel/api/public/v1"
	"gov2panel/internal/model/entity"
	"gov2panel/internal/service"
	"gov2panel/internal/utils"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
)

func (c *ControllerV1) Subscribe(ctx context.Context, req *v1.SubscribeReq) (res *v1.SubscribeRes, err error) {
	res = &v1.SubscribeRes{}
	user, err := service.User().GetUserByTokenAndUDAndGTExpiredAt(req.Token)
	if err != nil {
		return
	}

	serviceArr, err := service.ProxyService().GetServiceListByPlanIdAndShow1(user.GroupId)
	if err != nil {
		return
	}

	result := ""

	switch req.Flag {

	case "v2rayN": //v2rayn
		result = V2rayNSub(serviceArr, user)
	case "v2rayNG": //v2rayng
		result = V2rayNGSub(serviceArr, user)
	case "Shadowrocket": //shadowrocket
		result = ShadowrocketSub(serviceArr, user)
	case "Clash": //clash
		result = ClashSub(serviceArr, user)
		ghttp.RequestFromCtx(ctx).Response.WriteExit([]byte(result))
	case "Shadowsocks": //shadowsocks
		result = ShadowsocksSub(serviceArr, user)
	case "NekoBox": //nekobox
		result = NekoBoxSub(serviceArr, user)
	default:
		result = V2rayNSub(serviceArr, user)
	}

	ghttp.RequestFromCtx(ctx).Response.WriteExit(base64.StdEncoding.EncodeToString([]byte(result)))

	return nil, nil
}

// base64编码   单个：协议://base64编码
func base64Sub(serviceArr []*entity.V2ProxyService, user *entity.V2User) (result string) {
	for _, service := range serviceArr {
		service.Host = strings.ReplaceAll(service.Host, "$uuid$", user.Uuid)
		serviceJson := make(map[string]interface{})
		json.Unmarshal([]byte(service.ServiceJson), &serviceJson)
		switch strings.Split(service.Agreement, "/")[1] {
		case "vmess":
			s := map[string]string{
				"v":    "2",
				"add":  service.Host, //链接地址
				"ps":   service.Name, //名字
				"port": service.Port, //端口
				"id":   user.Uuid,    //uuid
				"aid":  "0",
				"net":  gconv.String(serviceJson["net"]),
				"type": gconv.String(serviceJson["type"]),
				"tls":  gconv.String(serviceJson["tls"]),
				"sni":  gconv.String(serviceJson["sni"]),
				"alpn": gconv.String(serviceJson["alpn"]),
				"host": gconv.String(serviceJson["host"]),
				"path": gconv.String(serviceJson["path"]),
				"scy":  gconv.String(serviceJson["scy"]),
				"fp":   gconv.String(serviceJson["fp"]),
			}
			ds, err := json.Marshal(s)
			if err != nil {
				return err.Error()
			}

			result = result + fmt.Sprintf("%s://%s\n", "vmess", base64.StdEncoding.EncodeToString(ds))
		case "vless":
			// vless://uuid@127.0.0.1:8888?encryption=none&security=reality&sni=sni.com&fp=qq&pbk=PublicKey&sid=ShortId&spx=SpiderX&type=tcp&headerType=http&host=host.com#vless
			// vless://uuid@127.0.0.1:8888?encryption=none&security=tls&sni=sni.com&alpn=http%2F1.1&fp=qq&pbk=PublicKey&sid=ShortId&spx=SpiderX&type=tcp&headerType=http&host=host.com#vless
			// vless://78f10ea1-81a4-4bf5-876f-90e3001f37dc@127.0.0.1:8888?encryption=none&flow=xtls-rprx-vision&security=tls&sni=sni.com&alpn=http%2F1.1&fp=qq&pbk=PublicKey&sid=ShortId&spx=SpiderX&type=tcp&headerType=http&host=host.com#vless

			result = result + fmt.Sprintf(
				"%s://%s@%s:%s?encryption=%s&flow=%s&security=%s&sni=%s&alpn=%s&fp=%s&pbk=%s&sid=%s&spx=%s&type=%s&serviceName=%s&mode=%s&headerType=%s&quicSecurity=%s&key=%s&host=%s&path=%s&seed=%s#%s\n",
				"vless",
				user.Uuid,
				service.Host,
				service.Port,
				gconv.String(serviceJson["encryption"]),
				gconv.String(serviceJson["flow"]),
				gconv.String(serviceJson["security"]),
				gconv.String(serviceJson["sni"]),
				gconv.String(serviceJson["alpn"]),
				gconv.String(serviceJson["fp"]),
				gconv.String(serviceJson["pbk"]),
				gconv.String(serviceJson["sid"]),
				gconv.String(strings.ReplaceAll(gconv.String(serviceJson["spx"]), "$uuid$", user.Uuid)),
				gconv.String(serviceJson["type"]),
				gconv.String(serviceJson["serviceName"]),
				gconv.String(serviceJson["mode"]),
				gconv.String(serviceJson["headerType"]),
				gconv.String(serviceJson["quicSecurity"]),
				gconv.String(serviceJson["key"]),
				gconv.String(serviceJson["host"]),
				gconv.String(serviceJson["path"]),
				gconv.String(serviceJson["seed"]),
				service.Name,
			)

		case "ss2022":
			ssPasswd := user.Uuid
			if gconv.String(serviceJson["cypher_method"]) == "2022-blake3-aes-128-gcm" {
				ssPasswd = gconv.String(serviceJson["server_key"]) + ":" + base64.StdEncoding.EncodeToString(
					gconv.Bytes(user.Uuid[0:16]),
				)
			}

			if gconv.String(serviceJson["cypher_method"]) == "2022-blake3-aes-256-gcm" {
				ssPasswd = gconv.String(serviceJson["server_key"]) + ":" + base64.StdEncoding.EncodeToString(
					gconv.Bytes(user.Uuid[0:32]),
				)
			}

			str := base64.StdEncoding.EncodeToString(
				gconv.Bytes(gconv.String(serviceJson["cypher_method"]) + ":" + ssPasswd),
			)
			str = strings.ReplaceAll(str, "+", "-")
			str = strings.ReplaceAll(str, "/", "_")
			str = strings.ReplaceAll(str, "=", "")

			// ss://base64(加密方式:密码)@地址:端口#别名
			// ss://OjY4ZDJjNTFmLTUzMTEtNDc2MS1hYTNhLTllNDg1MmYzMGYyNQ==@127.0.0.1:9996#ss2022
			result = result + fmt.Sprintf(
				"%s://%s@%s:%s#%s\n",
				"ss",
				str,
				service.Host,
				service.Port,
				service.Name,
			)

		case "trojan":
			//trojan://密码@地址:端口?security=tls&sni=sni.com&alpn=http%2F1.1&fp=chrome&type=tcp&headerType=none&host=host.com#名字

			result = result + fmt.Sprintf(
				"%s://%s@%s:%s?security=%s&sni=%s&alpn=%s&fp=%s&type=%s&headerType=%s&host=%s#%s\n",
				"trojan",
				user.Uuid,
				service.Host,
				service.Port,
				gconv.String(serviceJson["security"]),
				gconv.String(serviceJson["sni"]),
				gconv.String(serviceJson["alpn"]),
				gconv.String(serviceJson["fp"]),
				gconv.String(serviceJson["type"]),
				gconv.String(serviceJson["headerType"]),
				gconv.String(serviceJson["host"]),
				service.Name,
			)

		}

	}

	return
}

// ShadowrocketSub订阅
func ShadowrocketSub(serviceArr []*entity.V2ProxyService, user *entity.V2User) (result string) {
	result = result + fmt.Sprintf("STATUS=↑:%.2fGB,↓:%.2fGB,TOT:%.2fGBExpires:%s\n", utils.BytesToGB(user.U), utils.BytesToGB(user.D), utils.BytesToGB(user.TransferEnable), user.ExpiredAt)

	result = result + base64Sub(serviceArr, user)

	return
}

// v2rayNG订阅
func V2rayNGSub(serviceArr []*entity.V2ProxyService, user *entity.V2User) (result string) {
	s1 := map[string]string{
		"v":    "2",
		"add":  "127.0.0.1",                                  //链接地址
		"ps":   "套餐到期：" + user.ExpiredAt.Format("Y-m-d H:i"), //名字
		"net":  "tcp",
		"port": "80",      //端口
		"id":   user.Uuid, //uuid
		"aid":  "0",
	}
	ds1, err := json.Marshal(s1)
	if err != nil {
		return err.Error()
	}

	result = result + fmt.Sprintf("%s://%s\n", "vmess", base64.StdEncoding.EncodeToString(ds1))

	s1 = map[string]string{
		"v":    "2",
		"add":  "127.0.0.1",                                                                          //链接地址
		"ps":   "剩余流量：" + fmt.Sprintf("%.2f GB", utils.BytesToGB(user.TransferEnable-user.U-user.D)), //名字
		"port": "80",                                                                                 //端口
		"id":   user.Uuid,
		"aid":  "0",
		"net":  "tcp",
		"type": "none",
		"tls":  "",
		"sni":  "",
		"alpn": "",
		"host": "",
		"path": "",
		"scy":  "",
		"fp":   "", //uuid

	}
	ds1, err = json.Marshal(s1)
	if err != nil {
		return err.Error()
	}

	result = result + fmt.Sprintf("%s://%s\n", "vmess", base64.StdEncoding.EncodeToString(ds1))

	result = result + base64Sub(serviceArr, user)

	return
}

// v2rayN订阅
func V2rayNSub(serviceArr []*entity.V2ProxyService, user *entity.V2User) (result string) {
	s1 := map[string]string{
		"v":    "2",
		"add":  "127.0.0.1",                                  //链接地址
		"ps":   "套餐到期：" + user.ExpiredAt.Format("Y-m-d H:i"), //名字
		"net":  "tcp",
		"port": "80",      //端口
		"id":   user.Uuid, //uuid
		"aid":  "0",
	}
	ds1, err := json.Marshal(s1)
	if err != nil {
		return err.Error()
	}

	result = result + fmt.Sprintf("%s://%s\n", "vmess", base64.StdEncoding.EncodeToString(ds1))

	s1 = map[string]string{
		"v":    "2",
		"add":  "127.0.0.1",                                                                          //链接地址
		"ps":   "剩余流量：" + fmt.Sprintf("%.2f GB", utils.BytesToGB(user.TransferEnable-user.U-user.D)), //名字
		"port": "80",                                                                                 //端口
		"id":   user.Uuid,
		"aid":  "0",
		"net":  "tcp",
		"type": "none",
		"tls":  "",
		"sni":  "",
		"alpn": "",
		"host": "",
		"path": "",
		"scy":  "",
		"fp":   "", //uuid

	}
	ds1, err = json.Marshal(s1)
	if err != nil {
		return err.Error()
	}

	result = result + fmt.Sprintf("%s://%s\n", "vmess", base64.StdEncoding.EncodeToString(ds1))

	result = result + base64Sub(serviceArr, user)

	return
}

// NekoBox订阅
func NekoBoxSub(serviceArr []*entity.V2ProxyService, user *entity.V2User) (result string) {
	s1 := map[string]string{
		"v":    "2",
		"add":  "127.0.0.1",                                  //链接地址
		"ps":   "套餐到期：" + user.ExpiredAt.Format("Y-m-d H:i"), //名字
		"net":  "tcp",
		"port": "80",      //端口
		"id":   user.Uuid, //uuid
		"aid":  "0",
	}
	ds1, err := json.Marshal(s1)
	if err != nil {
		return err.Error()
	}

	result = result + fmt.Sprintf("%s://%s\n", "vmess", base64.StdEncoding.EncodeToString(ds1))

	s1 = map[string]string{
		"v":    "2",
		"add":  "127.0.0.1",                                                                          //链接地址
		"ps":   "剩余流量：" + fmt.Sprintf("%.2f GB", utils.BytesToGB(user.TransferEnable-user.U-user.D)), //名字
		"port": "80",                                                                                 //端口
		"id":   user.Uuid,
		"aid":  "0",
		"net":  "tcp",
		"type": "none",
		"tls":  "",
		"sni":  "",
		"alpn": "",
		"host": "",
		"path": "",
		"scy":  "",
		"fp":   "", //uuid

	}
	ds1, err = json.Marshal(s1)
	if err != nil {
		return err.Error()
	}

	result = result + fmt.Sprintf("%s://%s\n", "vmess", base64.StdEncoding.EncodeToString(ds1))

	result = result + base64Sub(serviceArr, user)

	return
}

// Shadowsocks订阅
func ShadowsocksSub(serviceArr []*entity.V2ProxyService, user *entity.V2User) (result string) {
	result = result + fmt.Sprintf(
		"%s://%s@%s:%s#%s\n",
		"ss",
		base64.StdEncoding.EncodeToString(
			gconv.Bytes("aes-128-gcm:"+user.Uuid),
		),
		"127.0.0.1",
		"80",
		"套餐到期："+user.ExpiredAt.Format("Y-m-d H:i"),
	)

	result = result + fmt.Sprintf(
		"%s://%s@%s:%s#%s\n",
		"ss",
		base64.StdEncoding.EncodeToString(
			gconv.Bytes("aes-128-gcm:"+user.Uuid),
		),
		"127.0.0.1",
		"80",
		"剩余流量："+fmt.Sprintf("%.2f GB", utils.BytesToGB(user.TransferEnable-user.U-user.D)),
	)

	for _, service := range serviceArr {
		service.Host = strings.ReplaceAll(service.Host, "$uuid$", user.Uuid)

		serviceJson := make(map[string]interface{})
		json.Unmarshal([]byte(service.ServiceJson), &serviceJson)
		switch strings.Split(service.Agreement, "/")[1] {
		case "ss2022":
			ssPasswd := user.Uuid
			if gconv.String(serviceJson["cypher_method"]) == "2022-blake3-aes-128-gcm" {
				ssPasswd = gconv.String(serviceJson["server_key"]) + ":" + base64.StdEncoding.EncodeToString(
					gconv.Bytes(user.Uuid[0:16]),
				)
			}

			if gconv.String(serviceJson["cypher_method"]) == "2022-blake3-aes-256-gcm" {
				ssPasswd = gconv.String(serviceJson["server_key"]) + ":" + base64.StdEncoding.EncodeToString(
					gconv.Bytes(user.Uuid[0:32]),
				)
			}

			str := base64.StdEncoding.EncodeToString(
				gconv.Bytes(gconv.String(serviceJson["cypher_method"]) + ":" + ssPasswd),
			)
			str = strings.ReplaceAll(str, "+", "-")
			str = strings.ReplaceAll(str, "/", "_")
			str = strings.ReplaceAll(str, "=", "")

			// ss://base64(加密方式:密码)@地址:端口#别名
			// ss://OjY4ZDJjNTFmLTUzMTEtNDc2MS1hYTNhLTllNDg1MmYzMGYyNQ==@127.0.0.1:9996#ss2022
			result = result + fmt.Sprintf(
				"%s://%s@%s:%s#%s\n",
				"ss",
				str,
				service.Host,
				service.Port,
				service.Name,
			)
		}
	}

	return
}

// Clash订阅
func ClashSub(serviceArr []*entity.V2ProxyService, user *entity.V2User) (result string) {

	content, err := os.ReadFile("./manifest/config/clash.yaml")
	if err != nil {
		result = "读取文件错误./manifest/config/clash.yaml，错误：" + err.Error()
		return
	}

	result = string(content)

	nodeNameArr := make([]string, 0)

	nodeInfoArr := make([]string, 0)

	for _, service := range serviceArr {
		service.Host = strings.ReplaceAll(service.Host, "$uuid$", user.Uuid)
		serviceJson := make(map[string]interface{})
		json.Unmarshal([]byte(service.ServiceJson), &serviceJson)

		nodeNameArr = append(nodeNameArr, "      - "+service.Name)

		d := map[string]interface{}{
			"name":   service.Name,
			"port":   service.Port,
			"server": service.Host,
		}

		switch strings.Split(service.Agreement, "/")[1] {
		case "vmess":
			d["type"] = "vmess"
			d["uuid"] = user.Uuid
			d["alterId"] = 0
			d["cipher"] = gconv.String(serviceJson["scy"]) //加密方式
			if gconv.String(serviceJson["tls"]) == "tls" { //tls
				d["tls"] = true
				d["skip-cert-verify"] = false
			}
			d["servername"] = gconv.String(serviceJson["sni"])
			d["network"] = gconv.String(serviceJson["net"]) //传输协议
			switch gconv.String(serviceJson["net"]) {       //传输协议
			case "ws":
				d["ws-path"] = ""
				d["ws-headers"] = map[string]interface{}{
					"Host": gconv.String(serviceJson["host"]),
				}
			case "h2":
				d["h2-opts"] = map[string]interface{}{
					"host": []string{gconv.String(serviceJson["host"])},
					"path": gconv.String(serviceJson["path"]),
				}
			case "grpc":
				d["grpc-opts"] = map[string]interface{}{
					"grpc-service-name": gconv.String(serviceJson["path"]),
				}

			}

			if gconv.String(serviceJson["type"]) == "http" { //伪装类型
				d["network"] = "http"
				d["http-opts"] = map[string]interface{}{
					"method":  "GET",
					"path":    []string{gconv.String(serviceJson["path"])},
					"headers": map[string]interface{}{"Host": []string{gconv.String(serviceJson["host"])}},
				}
			}

		case "vless":
			d["type"] = "vless"
			d["uuid"] = user.Uuid
			d["alterId"] = 0
			d["cipher"] = gconv.String(serviceJson["scy"]) //加密方式
			if gconv.String(serviceJson["tls"]) == "tls" { //tls
				d["tls"] = true
				d["skip-cert-verify"] = false
			}
			d["servername"] = gconv.String(serviceJson["sni"])
			d["network"] = gconv.String(serviceJson["net"]) //传输协议
			switch gconv.String(serviceJson["net"]) {       //传输协议
			case "ws":
				d["ws-path"] = ""
				d["ws-headers"] = map[string]interface{}{
					"Host": gconv.String(serviceJson["host"]),
				}
			case "h2":
				d["h2-opts"] = map[string]interface{}{
					"host": []string{gconv.String(serviceJson["host"])},
					"path": gconv.String(serviceJson["path"]),
				}
			case "grpc":
				d["grpc-opts"] = map[string]interface{}{
					"grpc-service-name": gconv.String(serviceJson["path"]),
				}

			}

			if gconv.String(serviceJson["type"]) == "http" { //伪装类型
				d["network"] = "http"
				d["http-opts"] = map[string]interface{}{
					"method":  "GET",
					"path":    []string{gconv.String(serviceJson["path"])},
					"headers": map[string]interface{}{"Host": []string{gconv.String(serviceJson["host"])}},
				}
			}
		case "ss2022":
			ssPasswd := user.Uuid
			if gconv.String(serviceJson["cypher_method"]) == "2022-blake3-aes-128-gcm" {
				ssPasswd = gconv.String(serviceJson["server_key"]) + ":" + base64.StdEncoding.EncodeToString(
					gconv.Bytes(user.Uuid[0:16]),
				)
			}

			if gconv.String(serviceJson["cypher_method"]) == "2022-blake3-aes-256-gcm" {
				ssPasswd = gconv.String(serviceJson["server_key"]) + ":" + base64.StdEncoding.EncodeToString(
					gconv.Bytes(user.Uuid[0:32]),
				)
			}
			d["type"] = "ss"
			d["cipher"] = gconv.String(serviceJson["cypher_method"]) //加密方式
			d["password"] = ssPasswd                                 //密码

		case "trojan":
			d["type"] = "trojan"
			d["password"] = user.Uuid
			d["sni"] = gconv.String(serviceJson["sni"])
			d["alpn"] = []string{gconv.String(serviceJson["alpn"])}
			d["skip-cert-verify"] = false
		}

		j, _ := gjson.EncodeString(d)
		fmt.Println(string(j))
		nodeInfoArr = append(nodeInfoArr, "  - "+j)
	}

	result = strings.ReplaceAll(result, "{{node_info}}", strings.Join(nodeInfoArr, "\n"))
	result = strings.ReplaceAll(result, "{{node_name}}", strings.Join(nodeNameArr, "\n"))

	return
}
