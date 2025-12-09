#include <bits/stdc++.h>
using namespace std;

#define MOD     1000000007
#define int     long long int
#define inf		1e15
#define pb      push_back
#define MAX     LLONG_MAX
#define MIN     LLONG_MIN
#define vi      vector<int>
#define mii     map<int, int>
#define pii     pair<int, int>
#define no      cout << "No\n"
#define yes     cout << "Yes\n"
#define sz(x)   ((int)(x).size())
#define all(x)  (x).begin(), (x).end()
#define ret(n)  cout <<setprecision(20)<<n<<endl
#define vin(a)  for(int i = 0; i < (a).size(); ++i) cin >> a[i];
#define vout(a) for(int i = 0; i < (a).size(); ++i) cout << a[i] << ' ';

const int N = 50;
const int M = 50;
bool vis[N][M];
vector<string> v(N);
map<char, vector<pii>> pos;
// 0 -> {}, {}, {}
// A -> {}, {}, {}
void solve() {
    vin(v);
    for (int i = 0; i < N; i++) {
        for (int j = 0; j < M; j++) {
            if (v[i][j] == '.') continue;
            pos[v[i][j]].pb({i, j});
        }
    }

    auto valid = [&](int x, int y) {
        return x >= 0 && x < N && y >= 0 && y < M;
    };

    for (auto& [ch, vec] : pos) {
        for (int i = 0; i < vec.size(); i++) {
            int x1 = vec[i].first, y1 = vec[i].second;
            for (int j = i + 1; j < vec.size(); j++) {
                int x2 = vec[j].first, y2 = vec[j].second;
                int x12 = x1 + (x1 - x2), y12 = y1 + (y1 - y2);
                int x22 = x2 + (x2 - x1), y22 = y2 + (y2 - y1);
                if (valid(x12, y12)) 
                    vis[x12][y12] = 1;
                if (valid(x22, y22)) 
                    vis[x22][y22] = 1;
            }
        }
    } 

    int ans = 0;
    for (int i = 0; i < N; i++)
        for (int j = 0; j < M; j++)
            ans += vis[i][j];
        
    cout << ans << endl;
}

signed main() {
    int t  =1;
    // cin >> t;
    while (t--) {
        solve();
    }
}
