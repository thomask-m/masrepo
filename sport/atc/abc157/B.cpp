#include <iostream>
#include <vector>

std::string has_win(int bingo[3][3], std::vector<int> selections)
{
  // row
  int count = 0;
  for (int i = 0; i < 3; i++) {
    for (int j = 0; j < 3; j++) {
      for (int si = 0; si < selections.size(); si++) {
        if (selections[si] == bingo[i][j]) {
          count++;
          break;
        }
      }
    }
    if (count == 3) return "Yes";
    count = 0;
  }

  // column
  for (int i = 0; i < 3; i++) {
    for (int j = 0; j < 3; j++) {
      for (int si = 0; si < selections.size(); si++) {
        if (selections[si] == bingo[j][i]) {
          count++;
          break;
        }
      }
    }
    if (count == 3) return "Yes";
    count = 0;
  }

  // r->l diagonal
  for (int i = 0; i < 3; i++) {
    for (int si = 0; si < selections.size(); si++) {
      if (selections[si] == bingo[i][i]) {
        count++;
        break;
      }
    }
  }
  if (count == 3) return "Yes";
  count = 0;

  // l->rdiagonal
  for (int i = 0; i < 3; i++) {
    for (int si = 0; si < selections.size(); si++) {
      if (selections[si] == bingo[i][2 - i]) {
        count++;
        break;
      }
    }
  }
  if (count == 3) return "Yes";
  return "No";
}

int main()
{
  int bingo[3][3];
  int n;

  for (int i = 0; i < 3; i++) {
    for (int j = 0; j < 3; j++) {
      std::cin >> bingo[i][j];
    }
  }
  std::cin >> n;
  std::vector<int> b(n);
  int input;
  for (int i = 0; i < n; i++) {
    std::cin >> input;
    b.push_back(input);
  }

  std::cout << has_win(bingo, b) << std::endl;
}

