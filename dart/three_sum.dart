void main(List<String> args) {
  print(threeSum([-1, 0, 1, 2, -1, -4]));
}

List<List<int>> threeSum(List<int> nums) {
  List<List<int>> result = [];
  if (nums.length < 3) {
    return result;
  }
  nums.sort();
  for (var i = 0; i < nums.length - 2; i++) {
    int a = nums[i];
    if (i > 0 && a == nums[i - 1]) {
      continue;
    }
    int left = i + 1;
    int right = nums.length - 1;
    while (left < right) {
      int sum = a + nums[left] + nums[right];
      if (sum == 0) {
        result.add([a, nums[left], nums[right]]);
        while (left < right && nums[left] == nums[left + 1]) {
          left++;
        }
        while (left < right && nums[right] == nums[right - 1]) {
          right--;
        }
        right--;
        left++;
      } else if (sum < 0) {
        left++;
      } else {
        right--;
      }
    }
  }
  return result;
}
