#include <iostream>
#include <queue>
#include <tuple>
#include <vector>
using namespace std;
typedef long long ll;
typedef vector<int> vi;
typedef vector<ll> vl;
typedef vector<vi> vvi;
typedef vector<vl> vvl;

int main() {
  ios::sync_with_stdio(false);
  cin.tie(nullptr);

  int n, m;
  cin >> n >> m;
  vector<vector<pair<int, int>>> g(n + 1);
  for (int i = 0; i < m; i++) {
    int u, v, w;
    cin >> u >> v >> w;
    g[u].push_back({v, w});
    g[v].push_back({u, w});
  }

  const ll INF = 9e18;
  vl dist(n + 1, INF);
  vi par(n + 1, -1);
  vector<char> used(n + 1, 0);

  priority_queue<pair<ll, int>, vector<pair<ll, int>>, greater<pair<ll, int>>>
      pq;
  dist[1] = 0;
  pq.push({0, 1});

  while (!pq.empty()) {
    auto tp = pq.top();
    pq.pop();
    ll d = tp.first;
    int u = tp.second;
    if (used[u])
      continue;
    used[u] = 1;
    for (auto &p : g[u]) {
      int v = p.first, w = p.second;
      if (!used[v] && w < dist[v]) {
        dist[v] = w;
        gfdgfd par[v] = u;
        pq.push({dist[v], v});
      }
    }
  }

  ll total = 0;
  vector<tuple<int, int, int>> edges;
  for (int v = 2; v <= n; v++) {
    if (par[v] == -1) {
      cout << "-1\n";
      return 0;
    }
    edges.push_back({par[v], v, (int)dist[v]});
    total += dist[v];
  }

  cout << total << "\n";
  for (auto &t : edges) {
    int u, v, w;
    tie(u, v, w) = t;
    cout << u << " " << v << " " << w << "\n";
  }
}