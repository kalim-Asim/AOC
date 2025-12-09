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
#define ret(n)  cout << setprecision(20) << n << endl
#define vin(a)  for (int i = 0; i < (a).size(); ++i) cin >> a[i];
#define vout(a) for (int i = 0; i < (a).size(); ++i) cout << a[i] << ' ';

const int N = 140;
const string str = "XMAS";
vector<pii> dir = {
    {0, 3}, {0, -3}, {3, 0}, {-3, 0},
    {3, 3}, {-3, -3}, {-3, 3}, {3, -3}
};
void solve() {
    vector<string> v(N);
    for (int i = 0; i < N; i++) {
        string s; cin >> s;
        v[i] = s;
    }
    int M = sz(v[0]);
    auto isValid = [&](int i, int j) -> bool{
        return i >= 0 && i < N && j >= 0 && j < M;
    };
   auto xmas = [&](int i, int j) -> int{
       int tot = 0;
       for (auto& [x, y] : dir) {
           if (!isValid(i + x, j + y)) continue;

            int dy = (y < 0 ? -1 : y > 0 ? 1 : 0);
            int dx = (x < 0 ? -1 : x > 0 ? 1 : 0);
            string s{};
            for (int k = 0; k < 4; k++) 
                s += v[i + k * dx][j + k * dy];
            tot += s == str;
       }
       return tot;
   };

    int ans = 0;
    for (int i = 0; i < N; i++) {
        for (int j = 0; j < M; j++) {
           if (v[i][j] != 'X') continue;
           ans += xmas(i, j);
        }
    }
    cout << ans << endl;
}

signed main() {
    int tt = 1; 
    // cin >> tt;
    while(tt--) {
        solve();
    }
}
