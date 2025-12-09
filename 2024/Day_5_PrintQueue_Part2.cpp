#include <bits/stdc++.h>
using namespace std;

int main() {
    // Input data as a string
    string input_data = R"()";

    // Split the input into two parts (pairs and comma-separated lists)
    istringstream stream(input_data);
    vector<pair<int, int>> pairs;
    vector<vector<int>> lists;

    string line;
    bool isReadingPairs = true;

    while (getline(stream, line)) {
        if (line.empty()) {
            isReadingPairs = false; // Switch to reading the lists
            continue;
        }

        if (isReadingPairs) {
            // Parse pairs separated by '|'
            size_t delim_pos = line.find('|');
            int first = stoi(line.substr(0, delim_pos));
            int second = stoi(line.substr(delim_pos + 1));
            pairs.emplace_back(first, second);
        } else {
            // Parse comma-separated lists
            vector<int> currentList;
            stringstream ss(line);
            string num;
            while (getline(ss, num, ',')) {
                currentList.push_back(stoi(num));
            }
            lists.push_back(currentList);
        }
    }

    map<int, vector<int>> mp;
    for (auto& [a, b] : pairs) {
        mp[a].push_back(b);
    }

    auto correct = [&](vector<int>& v) {
        int n = v.size();
        for (int i = 0; i < n; i++) {
            int key = v[i];
            for (int j = i + 1; j < n; j++) {
                int value = v[j];
                if (mp.find(key) != mp.end()) {
                    auto &vec = mp[key]; 
                    if (find(vec.begin(), vec.end(), value) != vec.end()) {
                        continue;
                    } else {
                        return false;
                    }
                } else {
                    return false;
                }
            }
        }
        return true;
    };

    auto ord = [&](vector<int>& v) {
        int n = v.size();
        vector<pair<int,int>> cnt(n);

        for (int i = 0; i < n; i++) {
            cnt[i].second = v[i];
            int key = v[i];
            for (int j = 0; j < n; j++) {
                int val = v[j];
                if (mp.find(key) == mp.end()) continue;

                auto &vec = mp[key];
                if (find(vec.begin(), vec.end(), val) != vec.end())
                    cnt[i].first++;
            }
        }

        sort(cnt.begin(), cnt.end(), [&](const pair<int,int>& a, const pair<int,int>& b){
            if (a.first != b.first) return a.first > b.first;
            else a.second >= b.second;
        });
        vector<int> vec(n);
        for (int i = 0; i < n; i++) {
            vec[i] = cnt[i].second;
        } 
        return vec;
    };

    long long val = 0;
    for (auto& v : lists) {
        if (correct(v)) continue;
        vector<int> vec = ord(v);
        val += vec[vec.size() / 2]; 
    }

    cout << val << endl;
    return 0;
}
