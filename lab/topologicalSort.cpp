#include <algorithm>
#include <iostream>
#include <vector>

using namespace std;

vector<int> g[100005];
int vis[100005];
vector<int> ans;

void dfs(int u) {
  vis[u] = 1;
  for (int v : g[u])
    if (!vis[v])
      dfs(v);
  ans.push_back(u);
}

int main() {
  int n, m;
  cin >> n >> m;
  for (int i = 0; i < m; i++) {
    int u, v;
    cin >> u >> v;
    g[u].push_back(v);
  }
  for (int i = 1; i <= n; i++)
    if (!vis[i])
      dfs(i);
  reverse(ans.begin(), ans.end());
  for (int x : ans)
    cout << x << " ";
}