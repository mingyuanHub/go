package main

import (
	"fmt"
	"net/url"
	"path"
)

func main() {
	fileUrl := "https://hybird.rayjump.com/rv-zip-2023/0217/endcard-dsp-1302-9d3a4d363e072a0658fb934642cf02a8.zip?md5filename=9d3a4d363e072a0658fb934642cf02a8&foldername=endcard-dsp-1302&mof=1&mof_uid=1868869&n_imp=1&mof_pkg=com.dressup.doll.vlinder.anime.avatar.maker&n_region=vg&alecfc=1&plee=1&bait_click=1&mof_textmod=1&bp_test=2&wglbp=1&mof_use_get=1&dlst=1&mof_use_get=1&plmug=1&admf=5"

	u, e := url.Parse(fileUrl)
	if e != nil {
		fmt.Println(1, e.Error())
	}



	fmt.Println(1.1, u.RawQuery)
	fmt.Println(1.2, u.Scheme + u.Host)
	fmt.Println(1.3, u.Path)

	fmt.Println(2, u.Path)
	fmt.Println(3, path.Ext(u.Path))


	fmt.Println(4, path.Ext(fileUrl))
	fmt.Println(5, path.Base(fileUrl))
	fmt.Println(6, path.Dir(fileUrl))
	fmt.Println(path.Split(fileUrl))
}
