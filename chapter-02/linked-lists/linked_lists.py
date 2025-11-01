class Node:
    def __init__(self, data, next) -> None:
        self.data: int = data
        self.next: Node = next
    

    def __str__(self):
        return str(self.data)


class LinkedList:
    def __init__(self) -> None:
        self.head = None
        self.tail = None


    def add_to_front(self, data: int) -> None:
        new_node = Node(data, None)
        if self.head == self.tail == None:
            self.head = new_node
            self.tail = new_node
        else:
            new_node.next = self.head
            self.head = new_node


    def add_to_back(self, data: int) -> None:
        new_node = Node(data, None)
        if self.head == self.tail == None:
            self.head = new_node
            self.tail = new_node
        else:
            self.tail.next = new_node
            self.tail = new_node


    def remove(self, index: int) -> None:
        if 0 > index or index >= self.length:
            raise IndexError()
        
        if index == 0:
            if self.head.next:
                self.head = self.head.next
            else:
                self.head = self.tail = None
        elif index == self.length - 1:
            current_node = self.head
            for _ in range(index - 1):
                current_node = current_node.next
            current_node.next = None
            self.tail = current_node
        else:
            current_node = self.head
            for i in range(index):
                if i == index - 1:
                    current_node.next = current_node.next.next
                    break
                current_node = current_node.next


    def insert_at(self, index: int, data: int) -> None:
        if 0 > index or index >= self.length:
            raise IndexError()
        
        new_node = Node(data, None)
        if index == 0:
            new_node.next = self.head
            self.head = new_node
        elif index == self.length - 1:
            self.tail.next = new_node
            self.tail = new_node
        else:
            for i, node in enumerate(self.__iter__()):
                if i == index - 1:
                    new_node.next = node.next
                    node.next = new_node
        

    @property
    def length(self) -> int:
        current_node = self.head
        if current_node == None:
            return 0
        
        counter = 1
        while current_node.next != None:
            counter += 1
            current_node = current_node.next

        return counter
    
    
    def __iter__(self) -> iter:
        current_node: Node = self.head
        while current_node:
            yield current_node
            current_node = current_node.next


    def __str__(self) -> str:
        return str([node.data for node in self.__iter__()])


if __name__ == "__main__":
    ll = LinkedList()

    ll.add_to_front(5)
    ll.add_to_front(4)
    ll.add_to_front(3)
    ll.add_to_front(2)
    ll.add_to_front(1)

    ll.remove(4 )

    print(ll.__str__())
    