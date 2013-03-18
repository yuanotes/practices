package shortestpath

import "math"

type Vertex struct {
    dist float64 //distance to source
    prev *Vertex
}

func Dijkstra(S, Q []*Vertex, s *Vertex) {
    //S : All vertices found 
    //Q : All vertices not found
    //s : the source point

    for _, v := range Q {
        v.dist = math.Inf(1)
    }
    s.dist = 0
    
    for _ , v := range Q {
        u := Exactmin(Q)
        S = append(S, u)
    }
}

func Exactmin(Q []*Vertex) *Vertex {
    if len(Q) == 0 {return nil}
    c := Q[0]
    for _, v := range Q {
        if c.dist > v.dist {
            c = v
        }
    }
    return c
}

/*
function Dijkstra(G, w, s)
     for each vertex v in V[G]      // 初始化
           d[v] := infinity         // 将各点的已知最短距离先设置成无穷大
           previous[v] := undefined // 各点的已知最短路径上的前趋都未知
     d[s] := 0                      // 因为出发点到出发点间不需移动任何距离，所以可以直接将s到s的最小距离设为0
     S := empty set
     Q := set of all vertices
     while Q is not an empty set        // Dijkstra演算法主體
           u := Extract_Min(Q)
           S.append(u)
           for each edge outgoing from u as (u,v)
                  if d[v] > d[u] + w(u,v)           // 拓展边（u,v）。w(u,v)为从u到v的路径长度。
                        d[v] := d[u] + w(u,v)       // 更新路径长度到更小的那个和值。
                        previous[v] := u            // 记录前趋顶点

procedure BellmanFord(list vertices, list edges, vertex source)
   // This implementation takes in a graph, represented as lists of vertices and edges,
   // and fills two arrays distance and predecessor with shortest-path information

   // Step 1: initialize graph
   for each vertex v in vertices:
       if v is source then distance[v] := 0
       else distance[v] := infinity
       predecessor[v] := null

   // Step 2: relax edges repeatedly
   for i from 1 to size(vertices)-1:
       for each edge (u, v) with weight w in edges:
           if distance[u] + w < distance[v]:
               distance[v] := distance[u] + w
               predecessor[v] := u

   // Step 3: check for negative-weight cycles
   for each edge (u, v) with weight w in edges:
       if distance[u] + w < distance[v]:
           error "Graph contains a negative-weight cycle"
*/
