// package exam

// func main() {
// 	//exam arrows = [6, 6, 6, 4, 4, 4, 2, 2, 2, 0, 0, 0, 1, 6, 5, 5, 3, 6, 0]
// 	//exam int = 3
// 	solution()
// }

// func solution(arrows []int) int {
//     var answer int64 = 0
//     var move = []int [(-1, 0), (-1, 1), (0, 1), (1, 1),
//             (1, 0), (1, -1), (0, -1), (-1, -1)]
//     var now = []iint (0, 0)

//     //visited is visite check in node
//     //visited_dir is Check node visit path ((A, B) means A -> B path)
//     visited = collections.defaultdict(int)
//     visited_dir = collections.defaultdict(int)

//     //enqueue node coordinates along arrows
//     queue = collections.deque([now])
//     for i in arrows:
//         //add 2 spaces in that direction to handle hourglass-shaped exceptions
//         for _ in range(2):
//             next = (now[0] + move[i][0], now[1] + move[i][1])
//             queue.append(next)

//             now = next

//     now = queue.popleft()
//     visited[now] = 1

//     while queue:
//         next = queue.popleft()

//         //if the node has already been visited (visited[x]==1)
//         if visited[next] == 1:
//             //if this is the first time you enter answer++
//             # when first entered, the room is first created
//             if visited_dir[(now, next)] == 0:
//                 answer += 1

// 		//if it is the first node visited, check the visit
//         else:
//             visited[next] = 1

//         //visit and check the path that entered the node
//         visited_dir[(now, next)] = 1
//         visited_dir[(next, now)] = 1
//         now = next

//     return answer
// }
