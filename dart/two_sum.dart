void main(List<String> args) {
  print(twoSum([3,2,4], 6));
}

List<int> twoSum(List<int> nums, int target) {
  List<int> result = [];
  for (var i = 0; i < nums.length; i++) {
    int a = target - nums[i];
    for (var j = i + 1; j < nums.length; j++) {
      if (nums[j] == a) {
        result.addAll([i, j]);
        break;
      }
    }
  }
  return result;
}
