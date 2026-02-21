class Solution:
    def insert(self, intervals: list[list[int]], newInterval: list[int]) -> list[list[int]]:
        res = []
        for interval in intervals:
            if interval[1] < newInterval[0]:
                # interval 完全在左邊
                res.append(interval)

            elif interval[0] > newInterval[1]:
                # interval 完全在右邊
                res.append(newInterval)
                newInterval = interval

            else:
                # interval 與 newInterval 重合
                if interval[0] < newInterval[0]:
                    newInterval[0] = interval[0]

                if interval[1] > newInterval[1]:
                    newInterval[1] = interval[1]

        res.append(newInterval)

        return res