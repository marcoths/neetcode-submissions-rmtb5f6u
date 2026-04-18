class Solution:
    def maxAreaOfIsland(self, grid: List[List[int]]) -> int:
        rows, cols = len(grid), len(grid[0])
        seen = set()
        maxArea = 0
        def dfs(row, col):
            directions = (
                (-1,0),
                (0,1),
                (0,-1),
                (1,0)
            )
            if ( row < 0 or col < 0 or row >= rows or col >= cols
            or (row, col) in seen or grid[row][col] == 0):
                return 0
            
            seen.add((row,col))

            area = 1
            
            for dr, dc in directions:
                nr, nc = row + dr, col + dc
                area += dfs(nr, nc)
            
            return area
        
        for row in range(rows):
            for col in range(cols):
                if grid[row][col] == 1:
                    area = dfs(row, col)
                    maxArea = max(area, maxArea)
        return maxArea

