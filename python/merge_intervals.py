# Definition for an interval.
# class Interval:
#     def __init__(self, s=0, e=0):
#         self.start = s
#         self.end = e

class Solution:
    def merge(self, intervals):
        """
        :type intervals: List[Interval]
        :rtype: List[Interval]
        """
        result = []
        for interval in sorted(intervals, key=lambda x:x.start):
            if result and interval.start <= result[-1].end:
                result[-1].end = max(interval.end, result[-1].end)
            else:
                result.append(interval)
        return result
