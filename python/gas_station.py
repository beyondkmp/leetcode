class Solution(object):
    def canCompleteCircuit(self, gas, cost):
        """
        :type gas: List[int]
        :type cost: List[int]
        :rtype: int
        """
        if sum(gas) < sum(cost):
            return -1

        length = len(gas)
        sorted_gas_index = sorted(range(len(gas)), key=lambda k: gas[k])

        for i in sorted_gas_index:
            j = 0
            remain = 0
            while j < length:
                index = (i + j) % length
                if remain + gas[index] < cost[index]:
                    break
                remain += gas[index] - cost[index]
                j += 1
            if j == length:
                return i
        return -1
