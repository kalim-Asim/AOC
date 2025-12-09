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
const string str = "MAS";
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
       string s1{}, s2{};
       if (isValid(i - 1, j - 1) && isValid(i + 1, j + 1)) {
           s1 += v[i - 1][j - 1];
           s1 += v[i][j];
           s1 += v[i + 1][j + 1];
       }
       if (isValid(i + 1, j - 1) && isValid(i - 1, j + 1)) {
           s2 += v[i + 1][j - 1];
           s2 += v[i][j];
           s2 += v[i - 1][j + 1];
       }
       string s3(s1), s4(s2);
       reverse(all(s3)); reverse(all(s4));
       if (s1 == str or s3 == str) {
           if (s2 == str or s4 == str) tot = 1;
       }
       return tot;
   };

    int ans = 0;
    for (int i = 0; i < N; i++) {
        for (int j = 0; j < M; j++) {
           if (v[i][j] != 'A') continue;
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
