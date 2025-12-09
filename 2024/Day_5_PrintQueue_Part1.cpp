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

    // Debug: Print parsed data
    // cout << "Pairs:" << endl;
    // for (const auto &p : pairs) {
    //     cout << "(" << p.first << ", " << p.second << ")" << endl;
    // }

    // cout << "\nLists:" << endl;
    // for (const auto &list : lists) {
    //     for (size_t i = 0; i < list.size(); ++i) {
    //         cout << list[i] << (i < list.size() - 1 ? ", " : "");
    //     }
    //     cout << endl;
    // }

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
                    auto &vec = mp[key]; // Get reference to the vector
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
    long long val = 0;
    for (auto& v : lists) {
        if (correct(v)) {
            val += v[v.size() / 2];
        }
    }
    cout << val << endl;
    return 0;
}
