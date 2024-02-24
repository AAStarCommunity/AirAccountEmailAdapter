package infra

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRFC(t *testing.T) {
	src1 := "=?UTF-8?B?44CQ5Y2H57qn5YWs5ZGK44CR5aSp6J6N5L+h572R?= =?UTF-8?B?57uc5a6J5YWo5oqA5pyv5pyJ6ZmQ5YWs5Y+4VA==?= =?UTF-8?B?b3BXQUYoV0FGKeaUu+WHu+ajgOa1i+inhOWImeW6k+WNh+e6p+WFrOWRig==?="
	chn1 := "【升级公告】天融信网络安全技术有限公司TopWAF(WAF)攻击检测规则库升级公告"
	msg1, err1 := decodeRFC2047String(src1)
	assert.NoError(t, err1)
	assert.Equal(t, chn1, msg1)

	src2 := "=?gb18030?B?YWxsLTIuMCAoMjAyMDAxMDmjqS50YXI=?="
	chn2 := "all-2.0 (20200109）.tar"
	msg2, err2 := decodeRFC2047String(src2)
	assert.NoError(t, err2)
	assert.Equal(t, chn2, msg2)

	src3 := "=?gb2312?B?udjT2sXtv62197avtcTNqNaq?="
	chn3 := "关于彭凯调动的通知"
	msg3, err3 := decodeRFC2047String(src3)
	assert.NoError(t, err3)
	assert.Equal(t, chn3, msg3)

	src4 := "=?UTF-8?B?5Lit5L+h6ZO26KGM5L+h55So5Y2h5Lit5b+D56Wd5oKo55Sf5pel5b+r5LmQ?= =?UTF-8?B?77yB?="
	msg4, _ := decodeRFC2047String(src4)
	println(msg4)
}
