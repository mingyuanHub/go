package main

import (
	"fmt"
	"regexp"
)

func main() {
	html := "<html><a id=\"aw0\" target=\"_top\" href=\"https://www.googleadservices.com/pagead/aclk?sa=L&amp;ai=CyOxK9Rf_Y-viHNCKid4PkqezuAQAgbCUpc4JwI23ARABIMLs3CBgnbnQgZAFiAEBoAH71NLXA6kCzZ-x9b2lpj6oAwGqBL8BT9Cf0DOkS6JvxRSZvmlFZvRouM-iPrihUMsINyG3PZUhOFw6Vq9i66nTRd-vIWUp5EED_BindZOllnATFMA9JA1yq3KB5ND2nrGdy2sSHRYY73Z_7YDLIcckEiNwMBo7ITkgbNrA7EHCjlrXJirxEAmS-Xxl15VO4auuZa0AzNEb3WY0GMxQ_It74_LxiKnYZeFfnSKFa3Pqqg0DJflPUrcQAU29xHAbRAQuO54lwXBvHLFBcdWWChCZPNvMeJrABKfx3_XmAYgF0Y3N1wWQBgGgBjnABguAB-2qrSiYBwGoB6a-G6gH1ckbqAehAagH_p6xAqgH89EbqAeW2BuoB6qbsQKoB_-esQKoB9-fsQKoB47cG6gHyZyxAqgH5pqxAtgHAaAI_IipBLAIArgIAdIIDwiAYRABGF8yAooCOgKAQLgMAaoNAkNOghQYGhZtb2JpbGVhcHA6OjEtNDc2OTQzMTQ2sBUB0BUBmBYB-BYBgBcB&amp;ae=1&amp;num=1&amp;cid=CAQShgEA1BOcpqQDDoJmEYxIV3MQpS8PZGqcBm5aVeJkT_fg_zYGzkbcER04D3CtB8N8ICNH4Qi8tnsB9RJVx9zPeASkj_xy0tTATHMhdHAu3-p2wQm4-eyMRDMUZcErcZtv1ENS4COPGJh-cVuprsvTC1rltxJnHnuavCUxqQUpUUKG2Ax5yyf4DxgB&amp;sig=AOD64_1rGsl46sIUDg8m0k8KpP14cM9zag&amp;fbs_aeid=[gw_fbsaeid]&amp;nb=17&amp;adurl=https://developers.google.com/admob/%%3Fgclid%%3DEAIaIQobChMIq7HV2rK6_QIVUEXCBR2S0wxHEAEYASAAEgKVofD_BwE\" data-asoch-targets=\"ad0\"><div data-ifc=\"[[[&quot;10,10,1,10&quot;,null,9,2]]]\" style=\"height: 41px;width: 320px;\"><img src=\"https://tpc.googlesyndication-cn.com/simgad/2923042507901230016\" border=\"0\" width=\"320\" height=\"41\" alt=\"\" class=\"img_ad\"></div></a></html>"


	html = "<html><a id=\"1\" href=\"https://www.googleadservices.com/pagead/aclk?sa=L&amp;\"> 123 </a><a id=\"2\"> </a></html>"

	reSidebar := regexp.MustCompile(`<a (.*?)</a>`)

	sidebar := reSidebar.FindAllString(html, -1)

	fmt.Println(sidebar)

	reLink := regexp.MustCompile(`href="(.*?)"`)
	// 找到所有链接
	links := reLink.FindAllString(sidebar[0], -1)

	fmt.Println(links)

}
