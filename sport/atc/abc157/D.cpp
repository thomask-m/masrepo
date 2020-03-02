#include <iostream>
#include <map>
#include <vector>

// TODO: still need to finish this
int main()
{
  int n, m, k;
  std::cin >> n;
  std::cin >> m;
  std::cin >> k;
  int res[n+1];
  for (int i = 0; i < n+1; i++) {
    res[i] = 0;
  }

  // graph setup: using adj list
  std::map<int, std::vector<int>> friends;
  std::map<int, std::vector<int>> blocks;

  // recording friendships
  int a, b;
  for (int _ = 0; _ < m; _++) {
    std::cin >> a;
    std::cin >> b;
    if (!friends.at(a)) {
      friends.insert(
    }
  }

  // recording blockships
  int c, d;
  for (int _ = 0; _ < k; _++) {
    std::cin >> c;
    std::cin >> d;
  }

  for (int i = 1; i < n+1; i++) {
    // gather possible friends candidates
    bool cands[n+1];
    for (int j = 0; j < n+1; j++) {
      cands[j] = true;
    }
    
  }
}
