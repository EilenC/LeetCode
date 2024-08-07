package part1

import (
	"sort"
	"strings"
	"testing"
)

func Test_groupAnagrams(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: `输入: strs = ["eat", "tea", "tan", "ate", "nat", "bat"]`,
			args: args{
				strs: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			},
			want: [][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}},
		},
		{
			name: `输入: strs = [""]`,
			args: args{
				strs: []string{""},
			},
			want: [][]string{{""}},
		},
		{
			name: `输入: strs = ["a"]`,
			args: args{
				strs: []string{"a"},
			},
			want: [][]string{{"a"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := groupAnagrams(tt.args.strs)
			if got == nil {
				t.Errorf("groupAnagrams() = %v, want %v", got, tt.want)
				return
			}
			if !areSlicesEqual(got, tt.want) {
				t.Errorf("groupAnagrams() = %v, want %v", got, tt.want)
				return
			}
		})
	}
}

// 辅助函数：排序内部切片并将其连接成一个字符串
func sortAndJoin(slice []string) string {
	sort.Strings(slice)
	return strings.Join(slice, ",")
}
func areSlicesEqual(slice1, slice2 [][]string) bool {
	if len(slice1) != len(slice2) {
		return false
	}

	// 创建一个映射来存储每个内部切片的排序后的字符串表示
	map1 := make(map[string]int)
	map2 := make(map[string]int)

	for _, innerSlice := range slice1 {
		key := sortAndJoin(innerSlice)
		map1[key]++
	}

	for _, innerSlice := range slice2 {
		key := sortAndJoin(innerSlice)
		map2[key]++
	}

	// 比较两个映射
	for key, count := range map1 {
		if map2[key] != count {
			return false
		}
	}

	return true
}
