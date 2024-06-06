package part1

/**
给定一个数组 coordinates ，其中 coordinates[i] = [x, y] ， [x, y] 表示横坐标为 x、纵坐标为 y 的点。请你来判断，这些点是否在该坐标系中属于同一条直线上。



示例 1：



输入：coordinates = [[1,2],[2,3],[3,4],[4,5],[5,6],[6,7]]
输出：true
示例 2：



输入：coordinates = [[1,1],[2,2],[3,4],[4,5],[5,6],[7,7]]
输出：false

*/

/**
我们知道，在给定的点集中，以任意一点 P0P_0P
0
​
  为基准，如果其他所有点的 ΔyΔx\dfrac{\Delta y}{\Delta x}
Δx
Δy
​
  是不变的，那么点集内所有的点在同一条直线上。但是这种做法会涉及到除数为 000 的问题，即垂直于 xxx 轴的直线需要单独判断。而且在计算浮点除法运算时还会涉及到精度问题，虽然在力扣中通过应该是没问题的，但是如果把测试集稍微设计一下就可能会通过不了。所以我们最好另寻他法。

我们可以把点集中除了 P0P_0P
0
​
  之外的点 PiP_iP
i
​
  都看成以 P0P_0P
0
​
  为起点、PiP_iP
i
​
  为终点的向量，记为 vi\boldsymbol v_iv
i
​
 ，并选择 v1\boldsymbol v_1v
1
​
  作为基准。如果其他向量都与 v1\boldsymbol v_1v
1
​
  共线，那么点集内所有的点共线。

这里需要用到线性代数的基础知识：如果二维向量 α\boldsymbol \alphaα 与 β\boldsymbol \betaβ 共线，那么它们线性相关，且有：
∣α, β∣=0，|\boldsymbol \alpha,\space\boldsymbol \beta|=0，∣α, β∣=0，
即它们拼成的二阶矩阵的行列式为 000。
*/

func checkStraightLine(coordinates [][]int) bool {
	x, y := coordinates[1][0]-coordinates[0][0], coordinates[1][1]-coordinates[0][1]
	for i := 2; i < len(coordinates); i++ {
		xi := coordinates[i][0] - coordinates[0][0]
		yi := coordinates[i][1] - coordinates[0][1]
		if (x*yi - y*xi) != 0 {
			return false
		}
	}
	return true
}
