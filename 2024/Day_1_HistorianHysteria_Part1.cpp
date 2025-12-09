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

const int N = 1000;
void solve() {
    vi a(N), b(N);
    vin(a);
    vin(b);
    sort(all(a));
    sort(all(b));
    int ans = 0;
    for (int i = 0; i < N; i++)
        ans += abs(a[i] - b[i]);
    cout << ans << endl;
}

signed main() {
    int tt = 1; 
    // cin >> tt;
    while(tt--) {
        solve();
    }
}
