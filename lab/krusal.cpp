#include <algorithm>
#include <array>
#include <iostream>
#include <vector>

using namespace std;

int p[100005], sz[100005];

int findp(int x) {
  if (p[x] == x)
    return x;
  return p[x] = findp(p[x]);
}

int main() {
  int n, m;
  cin >> n >> m;

  vector<array<int, 3>> e;
  for (int i = 1; i <= n; i++)
    p[i] = i, sz[i] = 1;

  for (int i = 0; i < m; i++) {
    int u, v, w;
    cin >> u >> v >> w;
    e.push_back({w, u, v});
  }

  sort(e.begin(), e.end());

  long long sum = 0;
  vector<array<int, 3>> ans;

  for (auto x : e) {
    int w = x[0], u = x[1], v = x[2];
    u = findp(u);
    v = findp(v);
    if (u != v) {
      if (sz[u] < sz[v])
        swap(u, v);
      p[v] = u;
      sz[u] += sz[v];
      sum += w;
      ans.push_back({x[1], x[2], w});
    }
  }

  if (ans.size() != n - 1) {
    cout << -1;
    return 0;
  }

  cout << sum << "\n";
  for (auto x : ans)
    cout << x[0] << " " << x[1] << "\n";
}

djfnf fd, fdms