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

void solve() {
    std::vector<std::string> strings;
    std::string input;
    while (std::cin >> input && input != "done") {
        strings.push_back(input); 
    }
    long long ans = 0;
    bool enable = true;
    string st1, st2;
    for (auto& str: strings) {
        int n = str.size();
        for (int i = 0; i < n; i++) {
            if (str.substr(i, 4) == "do()") enable = 1;
            if (str.substr(i, 7) == "don\'t()") enable = 0;
            string s{};
            for (int idx = i; idx < n && idx < i + 3; idx++) {
                s += str[idx];
            }
            if (s == "mul" && i + 3 < n && str[i + 3] == '(') {
                string val1{};
                int idx = i + 4;
                while(idx < n && str[idx] >= '0' && str[idx] <= '9')
                    val1 += str[idx], idx++;
                if (idx < n && str[idx] == ',') {
                    string val2{};
                    idx++;
                    while(idx < n && str[idx] >= '0' && str[idx] <= '9')
                        val2 += str[idx], idx++;
                    if (enable && !val1.empty() && !val2.empty() && idx < n && str[idx] == ')') {
                        int num1 = stoi(val1);
                        int num2 = stoi(val2);
                        ans += num1 * 1LL * num2;
                    }
                }
            }
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
