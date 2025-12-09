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

bool solve(int val, int idx, int curr, int n, vi& nums) {
    if (idx == n) return curr == val;

    // mul
    if (solve(val, idx + 1, curr + nums[idx], n, nums)) {
        return true;
    }

    // Try multiplying
    if (solve(val, idx + 1, curr * nums[idx], n, nums)) {
        return true;
    }
    // third operator -> ||
    if (curr) {
        string s = to_string(curr) + to_string(nums[idx]);
        if (solve(val, idx + 1, stoll(s), n, nums))
            return true;
    }
    
    return false;
}
void solve() {
    string line;
    // unordered_map<int, vector<int>> data;
    vector<pair<int, vector<int>>> data;
    while (getline(cin, line)) {
        stringstream ss(line);
        int key;
        char colon;
        string values;
        
        // Parse the key and the colon
        ss >> key >> colon;
        
        // Get the rest of the line as the values
        getline(ss, values);
        stringstream valuesStream(values);
        int value;
        vector<int> valueList;
        
        // Parse the space-separated integers
        while (valuesStream >> value) {
            valueList.push_back(value);
        }

        // data[key] = valueList;
        data.pb({key, valueList});
    }

    long long ans = 0;
    auto f = [&](int val, vi& nums) {
        int n = sz(nums);
        if (solve(val, 0, 0, n, nums)) 
        return true;    
        else 
        return false;
    };
    for (auto& [key, vec] : data) {
        if (f(key, vec)) ans += key;
    }

    cout << ans << endl;
}

signed main() {
    int t = 1;
    // cin >> t;
    while (t--) {
        solve();
    }
}
