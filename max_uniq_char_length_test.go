package main

import "testing"

func TestCharSet(t *testing.T) {
	a := new(CharSet)

	ok, b := a.Afflux("doit")
	if !ok {
		t.Fatal(1)
	}

	if b.Len() != 4 {
		t.Fatal(2)
	}

	ok, _ = b.Afflux("world")
	if ok {
		t.Fatal(3)
	}

	ok, c := b.Afflux("wrl")
	if !ok {
		t.Fatal(4)
	}

	if c.Len() != 7 {
		t.Fatal(5)
	}
}

func TestMaxUniqCharLength(t *testing.T) {
	l := MaxUniqCharLength([]string{"un", "iq", "ue"})
	if l != 4 {
		t.Fatal(1)
	}

	l = MaxUniqCharLength([]string{"cha", "r", "act", "ers"})
	if l != 6 {
		t.Fatal(2)
	}
	l = MaxUniqCharLength([]string{"abcdefghijklmnopqrstuvwxyz"})
	if l != 26 {
		t.Fatal(3)
	}

	l = MaxUniqCharLength([]string{"a", "abc", "d", "de", "def"})
	if l != 6 {
		t.Fatal(4)
	}
}

//给定一个字符串数组 arr，字符串 s 是将 arr 某一子序列字符串连接所得的字符串，如果 s 中的每一个字符都只出现过一次，那么它就是一个可行解。
//
//请返回所有可行解 s 中最长长度。
//
//
//
//示例 1：
//
//输入：arr = ["un","iq","ue"]
//输出：4
//解释：所有可能的串联组合是 "","un","iq","ue","uniq" 和 "ique"，最大长度为 4。
//示例 2：
//
//输入：arr = ["cha","r","act","ers"]
//输出：6
//解释：可能的解答有 "chaers" 和 "acters"。
//示例 3：
//
//输入：arr = ["abcdefghijklmnopqrstuvwxyz"]
//输出：26
//
//
//提示：
//
//1 <= arr.length <= 16
//1 <= arr[i].length <= 26
//arr[i] 中只含有小写英文字母
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/maximum-length-of-a-concatenated-string-with-unique-characters
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
