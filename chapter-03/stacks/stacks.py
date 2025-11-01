class Node:
    def __init__(self, data) -> None:
        self.data: int = data
        self.next: Node = None


class Stack:
    def __init__(self) -> None:
        self.top: Node = None
    

    def push(self, data: int) -> None:
        new_node = Node(data)
        if self.top == None:
            self.top = new_node
        else:
            new_node.next = self.top
            self.top = new_node
    

    def pop(self) -> None:
        if self.top == None:
            return
        
        self.top = self.top.next
        

    def peek(self) -> int:
        if self.top == None:
            return -1
        
        return self.top.data
    

if __name__ == "__main__":
    stack = Stack()

    for i in range(1, 11):
        stack.push(i)

    print(stack.peek(), end="\n\n") # 10

    for i in range(10):
        print(stack.peek())
        stack.pop()
