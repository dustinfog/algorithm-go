package main

type CharSet struct {
	l int
	m int32
}

func (cs *CharSet) Afflux(str string) bool {
	m := cs.m
	l := 0
	for _, c := range str {
		shift := int32(1) << (c - 'a')
		if m&shift != 0 {
			return false
		}

		m |= shift
		l++
	}

	cs.m = m
	cs.l += l
	return true
}

func (cs *CharSet) Len() int {
	return cs.l
}

func MaxUniqCharLength(arr []string) int {
	var dp []*CharSet
	for _, str := range arr {
		currCs := &CharSet{}
		if !currCs.Afflux(str) {
			continue
		}

		for _, cs := range dp {
			cs.Afflux(str)
		}

		dp = append(dp, currCs)
	}

	var max = 0
	for _, part := range dp {
		if part.Len() > max {
			max = part.Len()
		}
	}

	return max
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
