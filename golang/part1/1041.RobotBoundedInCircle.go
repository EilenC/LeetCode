package part1

//1041. 困于环中的机器人
//在无限的平面上，机器人最初位于 (0, 0) 处，面朝北方。注意:
//
//北方向 是y轴的正方向。
//南方向 是y轴的负方向。
//东方向 是x轴的正方向。
//西方向 是x轴的负方向。
//机器人可以接受下列三条指令之一：
//
//"G"：直走 1 个单位
//"L"：左转 90 度
//"R"：右转 90 度
//机器人按顺序执行指令 instructions，并一直重复它们。
//
//只有在平面中存在环使得机器人永远无法离开时，返回 true。否则，返回 false。

func isRobotBounded(instructions string) bool {
	var (
		row  int8
		x, y int
	)
	for i := 0; i < len(instructions); i++ {
		switch instructions[i] {
		case 'G':
			switch row {
			case 0:
				y++
			case 1:
				x--
			case 2:
				y--
			case 3:
				x++
			}
		case 'L':
			row = (row + 1) % 4
		default:
			row = (row + 3) % 4
		}
	}
	return x == 0 && y == 0 || row != 0
}
