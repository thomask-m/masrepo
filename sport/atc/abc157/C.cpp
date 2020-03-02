#include <iostream>

int main()
{
  int n, m;
  std::cin >> n;
  std::cin >> m;
  int conds[n+1];
  int si, ci;
  for (int i = 0; i < n+1; i++) {
    conds[i] = -1;
  }
  bool has_conflict = false;
  for (int i = 0; i < m; i++) {
    std::cin >> si;
    std::cin >> ci;
    if (conds[si] == -1) {
      conds[si] = ci;
    } else if (conds[si] != ci) {
      has_conflict = true;
    }
  }
  if (has_conflict) {
    std::cout << -1 << std::endl;
    return 0;
  }
  if (m == 0 && n == 1) {
    std::cout << 0 << std::endl;
    return 0;
  }
  for (int i = 1; i < n+1; i++) {
    if (n > 1 && i == 1 && conds[i] == 0) {
      std::cout << -1 << std::endl;
      return 0;
    }
    if (n > 1 && i > 1 && conds[i] == -1) {
      conds[i] = 0;
    }
  }
  if (conds[1] == -1) {
    conds[1] = 1;
  }
  for (int i = 1; i < n+1; i++) {
    std::cout << conds[i];
  }
  std::cout << std::endl;
  return 0;
}
