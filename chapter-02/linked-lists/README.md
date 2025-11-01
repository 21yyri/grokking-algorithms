# CH3 - Arrays & Linked Lists
Arrays are a data structure that have a fixed size and that is because of it's formation: it occupies a line of memory spaces. That means that any change on the array such as adding an item on a unreserved space would make the computer rearrange the array and look for a space which it could ocuppy again.

## The solution
This process costs time and space for the computer, so the solution was to build linked lists, which as the name suggests: are data connected by elos.

These are called nodes, and are composed by data and a reference to the next node, which are stored in different memory cells. This is more eficcient in cases where you don't wanna be working with fixed sized structures. 

## How to implement
You need a struct or class representing a node, which contains a data attribute that refers to what you wanna store, and next attribute, which points to the next node on the chain.

## Complexity of operations
It varies on the operation, but there's a coparision on the complexities found on the arrays and linked lists.

<table align="center" border="2px solid black">
    <tr>
        <td></td>
        <td>Array</td>
        <td>Linked List</td>
    </tr>
    <tr>
        <td>Acessing elements</td>
        <td>O(1)</td>
        <td>O(n)</td>
    </tr>
    <tr>
        <td>Insert/remove from beggining</td>
        <td>O(n)</td>
        <td>O(1)</td>
    </tr>
    <tr>
        <td>Insert/remove from end</td>
        <td>O(1)</td>
        <td>O(n)</td>
    </tr>
    <tr>
        <td>Insert/remove from middle</td>
        <td>O(1)</td>
        <td>O(n)</td>
    </tr>
</table>