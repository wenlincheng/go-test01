package main

import "fmt"

/*
	通过让 IPAddr 类型实现 fmt.Stringer 来打印点号分隔的地址。
	例如，IPAddr{1, 2, 3, 4} 应当打印为 "1.2.3.4"。
*/
// 内置数据类型具有默认的 String 方法
type IPAddr [4]byte

// TODO: 给 IPAddr 添加一个 "String() string" 方法
func (ip IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", ip[0], ip[1], ip[2], ip[3])
}

// 给 IPAddr 添加一个 "String() string" 方法 打印 ip 不会被匹配
//func (ip *IPAddr) String() string {
//	return fmt.Sprintf("%v.%v.%v.%v", ip[0], ip[1], ip[2], ip[3])
//}

type IP2 IPAddr

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}

	hosts2 := map[string]IP2{
		"local":   {127, 0, 0, 1},
		"alibaba": {8, 8, 8, 8},
	}
	for name, ip := range hosts2 {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
