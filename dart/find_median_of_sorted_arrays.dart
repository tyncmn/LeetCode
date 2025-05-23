void main(List<String> args) {
  print(findMedianSortedArrays([1, 2], [3, 4]));
}

double findMedianSortedArrays(List<int> nums1, List<int> nums2) {
  nums1.addAll([...nums2]);
  nums1.sort();
  final int m = nums1.length - 1;
  print(m);
  if (nums1.isEmpty) {
    return 0;
  } else {
    if (m % 2 == 0) {
      return nums1[m ~/ 2].toDouble();
    } else {
      return (nums1[(m / 2).floor()] + nums1[(m / 2).ceil()]) / 2;
    }
  }
}
