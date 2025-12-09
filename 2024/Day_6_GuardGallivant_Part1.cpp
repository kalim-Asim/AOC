#include <bits/stdc++.h>
using namespace std;

#define MOD     1000000007
#define inf     1e15
#define pb      push_back
#define MAX     LLONG_MAX
#define MIN     LLONG_MIN
#define vi      vector<int>
#define mii     map<int, int>
#define pii     pair<int, int>
#define no      cout << "NO\n"
#define yes     cout << "YES\n"
#define sz(x)   ((int)(x).size())
#define all(x)  (x).begin(), (x).end()
#define vin(a)  for (int i = 0; i < (a).size(); ++i) cin >> a[i];
#define vout(a) for (int i = 0; i < (a).size(); ++i) cout << a[i] << ' ';

const int N = 130;
map<char, int> mp = {
    {'^', 0}, {'>', 1}, {'v', 2}, {'<', 3}
};
map<char, int> mp2 = {
    {0, '^'}, {1, '>'}, {2, 'v'}, {3, '<'}
};
map<char, pii> dir = {
    {'^', {-1, 0}}, {'>', {0, 1}}, {'v', {1, 0}}, {'<', {0, -1}}
};
void solve() {
    vector<string> v(N);
    for (int i = 0; i < N; i++)
        cin >> v[i];
    int M = sz(v[0]);
    auto isValid = [&](int i, int j) {
        return i >= 0 && i < N && j >= 0 && j < M;
    };
    vector<vector<bool>> vis(N, vector<bool>(M, false));
    int x, y;
    char arrow;
    for (int i = 0; i < N; i++)
        for (int j = 0; j < M; j++)
            if (mp.find(v[i][j]) != mp.end()) 
                x = i, y = j, arrow = v[i][j];
    vis[x][y] = 1;
    while (true) {
        int dx = dir[arrow].first, dy = dir[arrow].second;
        while (isValid(x + dx, y + dy) && v[x + dx][y + dy] != '#') {
            x += dx, y += dy;
            vis[x][y] = 1;
        }
        if (!isValid(x + dx, y + dy)) break;
        arrow = mp2[(mp[arrow] + 1) % 4];
    }

    int tot = 0;
    for (int i = 0; i < N; i++)
        for (int j = 0; j < M; j++) 
            tot += vis[i][j];
    cout << tot << endl;
}

signed main() {
    int tt = 1; 
    // cin >> tt;
    while(tt--) {
        solve();
    }
}
