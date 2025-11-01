# CH3 - Stacks
Stacks are data structures defined by the LIFO (Last In, First Out), which means that all operations will involve the top of the pile, as adding or removing data from the top, or peeking what's there.

This model is used mostly on pile-like structures of data, like browser history, call stack, undo-redo functionalities.

## How to implement
The way I chose to implement it was based on my learning on linked lists. It is composed of a stack, which is structured on nodes. 

When initializing it, the stack top is pointing at nothing, and as nodes are added, they point to the bottom of the pile. Operations with the structure are only possible on the last element.

## Complexity of operations
The stack has some unique capacities and limitations, so there isn't any comparison.
<table align="center" border="2px solid black">
    <tr>
        <td></td>
        <td>Stack</td>
    </tr>
    <tr>
        <td>Acessing top</td>
        <td>O(1)</td>
    </tr>
    <tr>
        <td>Insert from top</td>
        <td>O(1)</td>
    </tr>
    <tr>
        <td>Remove from top</td>
        <td>O(1)</td>
    </tr>
</table>