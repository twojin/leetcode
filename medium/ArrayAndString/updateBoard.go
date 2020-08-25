package ArrayAndString

/*
让我们一起来玩扫雷游戏！

给定一个代表游戏板的二维字符矩阵。 'M' 代表一个未挖出的地雷，'E' 代表一个未挖出的空方块，'B' 代表没有相邻（上，下，左，右，和所有4个对角线）地雷的已挖出的空白方块，数字（'1' 到 '8'）表示有多少地雷与这块已挖出的方块相邻，'X' 则表示一个已挖出的地雷。

现在给出在所有未挖出的方块中（'M'或者'E'）的下一个点击位置（行和列索引），根据以下规则，返回相应位置被点击后对应的面板：

如果一个地雷（'M'）被挖出，游戏就结束了- 把它改为 'X'。
如果一个没有相邻地雷的空方块（'E'）被挖出，修改它为（'B'），并且所有和其相邻的未挖出方块都应该被递归地揭露。
如果一个至少与一个地雷相邻的空方块（'E'）被挖出，修改它为数字（'1'到'8'），表示相邻地雷的数量。
如果在此次点击中，若无更多方块可被揭露，则返回面板。


示例 1：

输入:

[['E', 'E', 'E', 'E', 'E'],
 ['E', 'E', 'M', 'E', 'E'],
 ['E', 'E', 'E', 'E', 'E'],
 ['E', 'E', 'E', 'E', 'E']]

Click : [3,0]

输出:

[['B', '1', 'E', '1', 'B'],
 ['B', '1', 'M', '1', 'B'],
 ['B', '1', '1', '1', 'B'],
 ['B', 'B', 'B', 'B', 'B']]


来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/minesweeper
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func updateBoard(board [][]byte, click []int) [][]byte {
	if len(board) < 1 || len(click) < 2 {
		return board
	}

	x, y := click[0], click[1]
	bx, by := len(board), len(board[0])

	if x < 0 || x > bx || y < 0 || y > by {
		return board
	}

	// 如果click的方块是雷 修改点击的方块为X
	// click的方块不是雷 遍历周围方块
	if board[x][y] == 'M' {
		board[x][y] = 'X'
	} else {
		dfs(board, x, y)
	}

	return board

}

// 辅助计算周围方块的坐标
var dirX = []int{-1, -1, -1, 0, 0, 1, 1, 1}
var dirY = []int{-1, 0, 1, -1, 1, -1, 0, 1}

func dfs(board [][]byte, x, y int) {
	cnt := 0
	// 统计周围雷的数量
	for i := 0; i < 8; i++ {
		tx, ty := x+dirX[i], y+dirY[i]
		if tx < 0 || tx >= len(board) || ty < 0 || ty >= len(board[0]) {
			continue
		}
		if board[tx][ty] == 'M' {
			cnt++
		}
	}

	// 周围雷数大于0 修改点击方块为雷的数量
	if cnt > 0 {
		board[x][y] = byte(cnt + '0')
	} else {
		board[x][y] = 'B'
		// 周围没有雷 递归遍历周围
		for i := 0; i < 8; i++ {
			tx, ty := x+dirX[i], y+dirX[i]
			// 不是M 被挖过 跳过
			if tx < 0 || tx >= len(board) || ty < 0 || ty >= len(board[0]) || board[tx][ty] != 'M' {
				continue
			}
			dfs(board, tx, ty)
		}
	}
}
