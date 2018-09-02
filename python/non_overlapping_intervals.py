# Definition for an interval.
# class Interval:
#     def __init__(self, s=0, e=0):
#         self.start = s
#         self.end = e


class Solution:
    def eraseOverlapIntervals(self, intervals):
        """
        :type intervals: List[Interval]
        :rtype: int
        """
        if len(intervals) == 0:
            return 0

        sorted_intervals = sorted(intervals, key=lambda x: x.end)
        compare = sorted_intervals[0]
        nums = 1
        for x in sorted_intervals[1:]:
            if x.start >= compare.end:
                nums += 1
                compare = x
        return len(intervals) - nums
