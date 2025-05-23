void main(List<String> args) {
  print(isMatch('aaa', "ab*a*c*a"));
}

bool isMatch(String s, String p) {
  int j = 0;
  int i;
  List<String> result = [];

  if (s.length > p.length) {
    return false;
  }
  for (i = 0; i < s.length; i++) {
    while (j < p.length) {
      if (p[j] == '.') {
        j++;
        result.add(s[i]);
        break;
      } else if (p[j] == '*' && (p[j - 1] == s[i] || p[j - 1] == '.')) {
        result.add(s[i]);
        j++;
        break;
      } else if (p[j] != s[i] && j + 1 < p.length && p[j + 1] == '*') {
        j += 2;
        continue;
      } else if (p[j] == s[i]) {
        j++;
        result.add(s[i]);
        break;
      } else {
        return false;
      }
    }
  }
  return result.length == s.length
      ? j == p.length
          ? true
          : false
      : false;
}
