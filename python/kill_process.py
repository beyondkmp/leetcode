class Solution:
    def killProcess(self, pid, ppid, kill):
        """
        :type pid: List[int]
        :type ppid: List[int]
        :type kill: int
        :rtype: List[int]
        """
        dic = {}
        visited = {}
        for p in pid:
            dic.setdefault(p, [])
            visited[p] = 0
        for i, p in enumerate(ppid):
            dic.setdefault(p, [])
            dic[p].append(pid[i])
        return self.killAll(visited, dic, kill)

    def killAll(self, visited, dic, kill):
        result = [kill]
        visited[kill] = 1
        for p in dic[kill]:
            if visited[p] == 0:
                visited[p] = 1
                result.extend(self.killAll(visited, dic, p))
        return result
