package part1

/*
*
思路1
1. 使用go的map 先遍历一遍字符串s 将字符作为key 值为出现次数
2. 编译字符串t，判断key存在的话直接删除
3. 最后判断map的长度是否为0
*/
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	tmp := make(map[int32]int)
	for _, v := range s {
		tmp[v] = tmp[v] + 1
	}
	for _, v := range t {
		if tmp[v] == 1 {
			delete(tmp, v)
			continue
		}
		tmp[v] = tmp[v] - 1
	}
	return len(tmp) == 0
}
