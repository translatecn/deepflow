package dropletpb

import (
	"net"
	"reflect"
	"testing"
)

var ipRanges = []struct{ start, end string }{
	{"0.0.0.0", "0.0.0.0"},
	{"0.0.0.0", "0.0.0.1"},
	{"255.255.255.251", "255.255.255.255"},
	{"255.255.255.255", "255.255.255.255"},
	{"0.0.0.0", "255.255.255.255"},
}

var basicResultGroup = [][]string{
	{"0.0.0.0/32"},
	{"0.0.0.0/31"},
	{"255.255.255.251/32", "255.255.255.252/30"},
	{"255.255.255.255/32"},
	{"0.0.0.0/0"},
}

var (
	start1       = net.ParseIP("10.30.100.0").To4()
	end1         = net.ParseIP("10.30.100.15").To4()
	basicResult1 = []string{"10.30.100.0/28"}

	start2       = net.ParseIP("10.25.0.1").To4()
	end2         = net.ParseIP("10.25.0.10").To4()
	basicResult2 = []string{"10.25.0.1/32", "10.25.0.2/31", "10.25.0.4/30", "10.25.0.8/31", "10.25.0.10/32"}

	start3       = net.ParseIP("0.0.0.0").To4()
	end3         = net.ParseIP("10.25.0.1").To4()
	basicResult3 = []string{"0.0.0.0/5", "8.0.0.0/7", "10.0.0.0/12", "10.16.0.0/13", "10.24.0.0/16", "10.25.0.0/31"}

	start4       = net.ParseIP("192.167.255.253").To4()
	end4         = net.ParseIP("192.169.0.4").To4()
	basicResult4 = []string{"192.167.255.253/32", "192.167.255.254/31", "192.168.0.0/16", "192.169.0.0/30", "192.169.0.4/32"}
)

func generateConvertResult(ips []net.IPNet) []string {
	result := make([]string, 0)
	for _, ip := range ips {
		result = append(result, ip.String())
	}
	return result
}

func checkParsedIPRangesResult(t *testing.T, basicResult, targetResult []string) bool {
	if !reflect.DeepEqual(basicResult, targetResult) {
		t.Log("Result:", targetResult, "\n")
		t.Log("Expect:", basicResult, "\n")
		return false
	} else {
		return true
	}
}

// 特殊IP
func TestSpecialIPsConvert(t *testing.T) {
	for index, ipRange := range ipRanges {
		start := net.ParseIP(ipRange.start).To4()
		end := net.ParseIP(ipRange.end).To4()
		ips := ipRangeConvert2CIDR(start, end)
		result := generateConvertResult(ips)
		if !checkParsedIPRangesResult(t, basicResultGroup[index], result) {
			t.Error("TestSpecialIPsConvert Check Failed!")
		}
	}
}

// 可转换成一个网段的ip
func TestOneNetIPsConvert(t *testing.T) {
	ips := ipRangeConvert2CIDR(start1, end1)
	result := generateConvertResult(ips)
	if !checkParsedIPRangesResult(t, basicResult1, result) {
		t.Error("TestOneNetIPsConvert Check Failed!")
	}
}

// 可转换成多个网段的ip
func TestMoreNetIPsConvert(t *testing.T) {
	ips := ipRangeConvert2CIDR(start2, end2)
	result := generateConvertResult(ips)
	if !checkParsedIPRangesResult(t, basicResult2, result) {
		t.Error("TestMoreNetIPsConvert Check Failed!")
	}
	ips = ipRangeConvert2CIDR(start3, end3)
	result = generateConvertResult(ips)
	if !checkParsedIPRangesResult(t, basicResult3, result) {
		t.Error("TestMoreNetIPsConvert Check Failed!")
	}
	ips = ipRangeConvert2CIDR(start4, end4)
	result = generateConvertResult(ips)
	if !checkParsedIPRangesResult(t, basicResult4, result) {
		t.Error("TestMoreNetIPsConvert Check Failed!")
	}
}
