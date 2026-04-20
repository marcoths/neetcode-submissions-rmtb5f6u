# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right

class Solution:   
    def isSubtree(self, root: Optional[TreeNode], subRoot: Optional[TreeNode]) -> bool:
        def same(p: Optional[TreeNode], q: Optional[TreeNode]) -> bool:
            stack = [(p, q)]

            while stack:
                a, b = stack.pop()

                if not a and not b:
                    continue
                
                if (not a or not b or a.val != b.val):
                    return False
                stack.append((a.left, b.left))
                stack.append((a.right, b.right))
            return True
            
        if not root:
            return False
        stack = [root]
        while stack:
            node = stack.pop()
            if node.val == subRoot.val and same(node, subRoot):
                return True
            if node.left:
                stack.append(node.left)
            if node.right:
                stack.append(node.right)
        return False
            