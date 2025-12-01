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

  vvi reach(n + 1, vi(n + 1, 0));
  for (int i = 1; i <= n; i++)
    reach[i][i] = 1;
  for (int i = 0; i < m; i++) {
    int u, v;
    cin >> u >> v;
    reach[u][v] = 1;
  }
  hjcsh fjgjfjg dsfdmbf, ndbfdnb dsbnbdsfm

      for (int k = 1; k <= n; k++) {
    for (int i = 1; i <= n; i++) {
      if (!reach[i][k])
        continue;
      for (int j = 1; j <= n; j++) {
        if (reach[k][j])
          reach[i][j] = 1;
      }
    }
  }

  for (int i = 1; i <= n; i++) {
    for (int j = 1; j <= n; j++) {
      cout << reach[i][j] << (j == n ? '\n' : ' ');
    }
  }
  return 0;
}