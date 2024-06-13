# Binary Tree Implementation in Go

This repository contains a Go implementation of a binary tree with functions to insert values and traverse the tree in different orders (in-order, pre-order, and post-order).

## Description

A binary tree is a tree data structure in which each node has at most two children, referred to as the left child and the right child. This implementation allows you to insert values into the binary tree and perform three types of tree traversal:

1. **In-Order Traversal**: Visits the left branch, then the current node, and finally the right branch.
2. **Pre-Order Traversal**: Visits the current node before its child nodes.
3. **Post-Order Traversal**: Visits the current node after its child nodes.

## Code Structure

The main components of the code are:

- `Node`: Represents a node in the binary tree.
- `Insert`: Inserts a new node with the given value into the binary tree.
- `InOrderTraversal`: Traverses the binary tree in in-order and prints the values.
- `PreOrderTraversal`: Traverses the binary tree in pre-order and prints the values.
- `PostOrderTraversal`: Traverses the binary tree in post-order and prints the values.

## How to Use

1. **Clone the repository:**

   ```bash
   git clone https://github.com/sakiphan/Binary-Search-Tree.git
   cd Binary-Search-Tree
   ```

2. **Run the code:**

   ```bash
   go run main.go
   ```

   This will insert the following values into the binary tree: 10, 5, 15, 3, 7, 13, 17, and then perform and print the results of the three different tree traversals.

## Example Output

When you run the code, you will see output similar to the following:

```
In-Order Traversal:
3
5
7
10
13
15
17

Pre-Order Traversal:
10
5
3
7
15
13
17

Post-Order Traversal:
3
7
5
13
17
15
10
```

## Requirements

- Go 1.16 or later

