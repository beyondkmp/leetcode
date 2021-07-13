function groupAnagrams(strs: string[]): string[][] {
  const dicts: Map<string, string[]> = new Map();
  strs.forEach(str => {
    const sortedStr = Array.from(str).sort().join('');
    if (!dicts.has(sortedStr)) {
      dicts.set(sortedStr, []);
    }

    dicts.set(sortedStr, [...dicts.get(sortedStr) || [], str]);
  })

  const result: string[][] = [];
  for (const v of dicts.values()) {
    result.push(v);
  }
  return result;
};

console.log(groupAnagrams([""]));
console.log(groupAnagrams(["eat", "tea", "tan", "ate", "nat", "bat"]));
console.log(groupAnagrams(["a"]));

