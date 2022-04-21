# Graph

Graph is consist of vertices and edges.

To store relation between vertices, we can represent it as Adjacency Matrix and Adjacency List

### Adjacency Matrix
- Edge Lookup: Faster to see lookup (O(1)) with array lookup
- Add Vertex: O(1)
- Remove Vertex: O(E)
- Add Edge: O(1)
- Remove Edge: O(V)

### Adjacency List
- Edge Lookup: Lookup (O(v)) since we need to traverse
- Add Vertex : O(V^2)
- Remove Vertex: O(V^2)
- Add Edge: O(1)
- Remove Edge: O(1)
